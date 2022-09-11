package main

import (
	"github.com/dollarkillerx/go-proxy-manager/internal/app/agent/server"
	"log"
)

func main() {
	ser := server.NewServer()
	if err := ser.Run(); err != nil {
		log.Fatalln(err)
	}
}
