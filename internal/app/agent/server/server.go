package server

import (
	"github.com/dollarkillerx/go-proxy-manager/proto/agent"
	"google.golang.org/grpc"

	"log"
	"net"
)

type Server struct {
	cache *cache
}

func NewServer() *Server {
	return &Server{
		cache: newCache(),
	}
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
	log.Println("GRPC: 8501")
	log.Println("=========================")

	// grpc
	listen, err := net.Listen("tcp", ":8501")
	if err != nil {
		return err
	}
	rpc := grpc.NewServer(grpc.UnaryInterceptor(grpcAuth))

	agent.RegisterAgentServer(rpc, s)

	go func() {
		if err := rpc.Serve(listen); err != nil {
			log.Fatalln(err)
		}
	}()

	return s.listenAndTLS()
}
