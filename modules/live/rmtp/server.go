package rmtp

import (
	"errors"
	"net/http"
	"sync"

	log "github.com/sirupsen/logrus"
)

type Server struct {
	Items *sync.Map
}

func NewServer() *Server {
	return &Server{
		Items: &sync.Map{},
	}
}

func (s *Server) Serve(w http.ResponseWriter, r *http.Request, channel string) error {
	h, ok := w.(http.Hijacker)
	if !ok {
		return errors.New("Hijacker error")
	}
	netConn, _, err := h.Hijack()
	if err != nil {
		return err
	}
	conn := NewProtocol(netConn, 4*1024)
	s.Items.Store(channel, conn)
	go s.handleConn(conn)
	return nil
}

func (s *Server) handleConn(conn *Protocol) error {
	if err := conn.HandshakeServer(); err != nil {
		conn.Close()
		log.Error("handleConn HandshakeServer err: ", err)
		return err
	}

	return nil
}
