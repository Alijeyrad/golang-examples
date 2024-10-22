package main

import (
	"examples/server"
	"log"
	"net"
)

func main() {
	ln, _ := net.Listen("tcp", ":8080")
	server, _ := server.NewServer(ln)

	log.Fatal(server.Start())

}
