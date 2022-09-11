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

	log.Println("=========================")
	log.Println("GO Proxy Manager Agent: ")
	log.Println("HTTP: 80")
	log.Println("HTTPS: 443")
	log.Println("=========================")
	return s.listenAndTLS()
}
