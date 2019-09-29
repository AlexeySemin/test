package main

import (
	"log"

	"github.com/AlexeySemin/test/golang-service/server"
)

// @title Swagger Example API
// @version 0.0.1
// @host localhost:8081
// @license.name MIT
func main() {
	port := 8081

	srv, err := server.NewServer(port)
	if err != nil {
		log.Fatal(err)
	}

	srv.Start()
}
