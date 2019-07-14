package main

import (
	"github.com/AlexeySemin/test/golang-service/server"
)

func main() {
	port := 8081
	server.NewServer(port).Start()
}
