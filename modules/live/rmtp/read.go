package rmtp

import "io"

func (rw *Protocol) Read(p []byte) (int, error) {
	n, err := io.ReadAtLeast(rw.ReadWriter, p, len(p))
	return n, err
}

func (rw *Protocol) ReadUintBE(n int) (uint32, error) {
	ret := uint32(0)
	for i := 0; i < n; i++ {
		b, err := rw.ReadWriter.ReadByte()
		if err != nil {
			return 0, err
		}
		ret = ret<<8 + uint32(b)
	}
	return ret, nil
}

func (rw *Protocol) ReadUintLE(n int) (uint32, error) {
	ret := uint32(0)
	for i := 0; i < n; i++ {
		b, err := rw.ReadWriter.ReadByte()
		if err != nil {
			return 0, err
		}
		ret += uint32(b) << uint32(i*8)
	}
	return ret, nil
}

func (conn *Protocol) ReadChunk(c *Chunk) error {
	for {
		h, _ := conn.ReadUintBE(1)
		// if err != nil {
		// 	log.Println("read from conn error: ", err)
		// 	return err
		// }
		format := h >> 6
		csid := h & 0x3f
		cs, ok := conn.chunks[csid]
		if !ok {
			cs = Chunk{}
			conn.chunks[csid] = cs
		}
		cs.tmpFromat = format
		cs.CSID = csid
		err := cs.readChunk(conn, conn.remoteChunkSize, conn.pool)
		if err != nil {
			return err
		}
		conn.chunks[csid] = cs
		if cs.full() {
			*c = cs
			break
		}
	}

	conn.handleControlMsg(c)

	conn.ack(c.Length)

	return nil
}
