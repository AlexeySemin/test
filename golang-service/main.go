package main

import (
	"log"

	"github.com/AlexeySemin/test/golang-service/server"
)

func main() {
	port := 8081

	srv, err := server.NewServer(port)
	if err != nil {
		log.Fatal(err)
	}

	srv.Start()
}
