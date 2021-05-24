package rmtp

import "encoding/binary"

func (rw *Protocol) Write(p []byte) (int, error) {
	return rw.ReadWriter.Write(p)
}

func (rw *Protocol) WriteUintBE(v uint32, n int) error {
	for i := 0; i < n; i++ {
		b := byte(v>>uint32((n-i-1)<<3)) & 0xff
		if err := rw.ReadWriter.WriteByte(b); err != nil {
			return err
		}
	}
	return nil
}

func (rw *Protocol) WriteUintLE(v uint32, n int) error {
	for i := 0; i < n; i++ {
		b := byte(v) & 0xff
		if err := rw.ReadWriter.WriteByte(b); err != nil {
			return err
		}
		v = v >> 8
	}
	return nil
}

func (conn *Protocol) writeChunk(c *Chunk) error {
	if c.TypeID == idSetChunkSize {
		conn.chunkSize = binary.BigEndian.Uint32(c.Data)
	}
	return c.writeChunk(conn, int(conn.chunkSize))
}
