package main

import (
	"github.com/dollarkillerx/go-proxy-manager/internal/app/backend/server"

	"log"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Llongfile)

	ser := server.NewServer()
	if err := ser.Run(); err != nil {
		log.Fatalln(err)
	}
}
