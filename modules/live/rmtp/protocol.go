package rmtp

import (
	"bufio"
	"encoding/binary"
	"net"
	"time"

	"zodream.cn/godream/modules/live/pio"
	"zodream.cn/godream/modules/live/pool"
)

const (
	_ = iota
	idSetChunkSize
	idAbortMessage
	idAck
	idUserControlMessages
	idWindowAckSize
	idSetPeerBandwidth
)

type Protocol struct {
	net.Conn
	ReadWriter          *bufio.ReadWriter
	chunkSize           uint32
	remoteChunkSize     uint32
	windowAckSize       uint32
	remoteWindowAckSize uint32
	received            uint32
	ackReceived         uint32
	pool                *pool.Pool
	chunks              map[uint32]Chunk
}

func NewProtocol(c net.Conn, bufferSize int) *Protocol {
	return &Protocol{
		Conn:                c,
		ReadWriter:          bufio.NewReadWriter(bufio.NewReaderSize(c, bufferSize), bufio.NewWriterSize(c, bufferSize)),
		chunkSize:           128,
		remoteChunkSize:     128,
		windowAckSize:       2500000,
		remoteWindowAckSize: 2500000,
		pool:                pool.NewPool(),
		chunks:              make(map[uint32]Chunk),
	}
}

func (rw *Protocol) Flush() error {
	if rw.ReadWriter.Writer.Buffered() == 0 {
		return nil
	}
	return rw.ReadWriter.Flush()
}

func (conn *Protocol) Close() error {
	return conn.Conn.Close()
}

func (conn *Protocol) RemoteAddr() net.Addr {
	return conn.Conn.RemoteAddr()
}

func (conn *Protocol) LocalAddr() net.Addr {
	return conn.Conn.LocalAddr()
}

func (conn *Protocol) SetDeadline(t time.Time) error {
	return conn.Conn.SetDeadline(t)
}

func (conn *Protocol) NewAck(size uint32) Chunk {
	return initControlMsg(idAck, 4, size)
}

func (conn *Protocol) NewSetChunkSize(size uint32) Chunk {
	return initControlMsg(idSetChunkSize, 4, size)
}

func (conn *Protocol) NewWindowAckSize(size uint32) Chunk {
	return initControlMsg(idWindowAckSize, 4, size)
}

func (conn *Protocol) NewSetPeerBandwidth(size uint32) Chunk {
	ret := initControlMsg(idSetPeerBandwidth, 5, size)
	ret.Data[4] = 2
	return ret
}

func (conn *Protocol) handleControlMsg(c *Chunk) {
	if c.TypeID == idSetChunkSize {
		conn.remoteChunkSize = binary.BigEndian.Uint32(c.Data)
	} else if c.TypeID == idWindowAckSize {
		conn.remoteWindowAckSize = binary.BigEndian.Uint32(c.Data)
	}
}

func (conn *Protocol) ack(size uint32) {
	conn.received += uint32(size)
	conn.ackReceived += uint32(size)
	if conn.received >= 0xf0000000 {
		conn.received = 0
	}
	if conn.ackReceived >= conn.remoteWindowAckSize {
		cs := conn.NewAck(conn.ackReceived)
		cs.writeChunk(conn, int(conn.chunkSize))
		conn.ackReceived = 0
	}
}

func initControlMsg(id, size, value uint32) Chunk {
	ret := Chunk{
		Format:   0,
		CSID:     2,
		TypeID:   id,
		StreamID: 0,
		Length:   size,
		Data:     make([]byte, size),
	}
	pio.PutU32BE(ret.Data[:size], value)
	return ret
}

const (
	streamBegin      uint32 = 0
	streamEOF        uint32 = 1
	streamDry        uint32 = 2
	setBufferLen     uint32 = 3
	streamIsRecorded uint32 = 4
	pingRequest      uint32 = 6
	pingResponse     uint32 = 7
)

/*
   +------------------------------+-------------------------
   |     Event Type ( 2- bytes )  | Event Data
   +------------------------------+-------------------------
   Pay load for the ‘User Control Message’.
*/
func (conn *Protocol) userControlMsg(eventType, buflen uint32) Chunk {
	var ret Chunk
	buflen += 2
	ret = Chunk{
		Format:   0,
		CSID:     2,
		TypeID:   4,
		StreamID: 1,
		Length:   buflen,
		Data:     make([]byte, buflen),
	}
	ret.Data[0] = byte(eventType >> 8 & 0xff)
	ret.Data[1] = byte(eventType & 0xff)
	return ret
}

func (conn *Protocol) SetBegin() {
	ret := conn.userControlMsg(streamBegin, 4)
	for i := 0; i < 4; i++ {
		ret.Data[2+i] = byte(1 >> uint32((3-i)*8) & 0xff)
	}
	conn.writeChunk(&ret)
}

func (conn *Protocol) SetRecorded() {
	ret := conn.userControlMsg(streamIsRecorded, 4)
	for i := 0; i < 4; i++ {
		ret.Data[2+i] = byte(1 >> uint32((3-i)*8) & 0xff)
	}
	conn.writeChunk(&ret)
}
