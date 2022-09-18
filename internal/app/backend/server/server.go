package server

import (
	"github.com/dollarkillerx/go-proxy-manager/internal/app/backend/middleware"
	"github.com/dollarkillerx/go-proxy-manager/internal/app/backend/storage"
	"github.com/dollarkillerx/go-proxy-manager/internal/app/backend/storage/simple"
	"github.com/dollarkillerx/go-proxy-manager/proto/backend"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"

	"log"
	"net"
)

type Server struct {
	app     *gin.Engine
	storage storage.Interface
}

func NewServer() *Server {
	app := gin.New()
	app.Use(middleware.SetBasicInformation())
	app.Use(middleware.Cors())
	app.Use(middleware.HttpRecover())
	app.Use(gin.Logger())

	newSimple, err := simple.NewSimple()
	if err != nil {
		log.Fatalln(err)
	}

	return &Server{
		app:     app,
		storage: newSimple,
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
