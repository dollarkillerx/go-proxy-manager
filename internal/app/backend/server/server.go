package server

import (
	"github.com/dollarkillerx/go-proxy-manager/proto/backend"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"

	"log"
	"net"
)

type Server struct {
	app *gin.Engine
}

func NewServer() *Server {
	app := gin.New()
	app.Use(gin.Recovery())
	app.Use(gin.Logger())

	return &Server{
		app: app,
	}
}

func (s *Server) Run() error {
	log.Println("=========================")
	log.Println("GO Proxy Manager Backend: ")
	log.Println("HTTP Backend: 8746")
	log.Println("GRPC: 8502")
	log.Println("=========================")

	// grpc
	listen, err := net.Listen("tcp", ":8502")
	if err != nil {
		return err
	}
	rpc := grpc.NewServer(grpc.UnaryInterceptor(grpcAuth))

	backend.RegisterBackendServer(rpc, s)

	go func() {
		if err := rpc.Serve(listen); err != nil {
			log.Fatalln(err)
		}
	}()

	return s.app.Run(":8746")
}
