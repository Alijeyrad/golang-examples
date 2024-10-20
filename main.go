package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
)

func main() {
	// Listen on TCP port 8080
	ln, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ln.Close()
	fmt.Println("TCP server listening on port 8080")

	for {
		// Accept a connection
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		// Handle the connection in a new goroutine
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	for {
		// request
		decoder := json.NewDecoder(conn)
		type Request struct {
			Message string `json:"message"`
		}
		var req Request
		err := decoder.Decode(&req)
		if err != nil {
			fmt.Println("Error Decode:", err)
			return
		}
		log.Println(req)

		// response
		msg := Request{
			Message: "hello",
		}
		encoder := json.NewEncoder(conn)
		err = encoder.Encode(&msg)
		if err != nil {
			fmt.Println("Error Encode:", err)
			return
		}
	}
}
