package server

import (
	"log"
)

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Run() error {
	go func() {
		if err := s.listen(); err != nil {
			log.Fatalln(err)
		}
	}()

	return s.listenAndTLS()
}
