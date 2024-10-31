package server

import (
	"context"
	"encoding/json"
	"examples/models"
	"fmt"
	"log"
	"net"
	"time"
)

type Server struct {
	Ln          net.Listener
	Connections []net.Conn
}

type Connection struct {
	Conn       net.Conn
	Ctx        context.Context
	CancelFunc context.CancelFunc
}

func NewServer(ln net.Listener) (*Server, error) {
	return &Server{
		Ln: ln,
	}, nil
}

func (s *Server) Start() error {
	log.Println("server started...")
	for {
		// Accept a connection
		conn, err := s.Ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)

		connection := &Connection{
			Conn:       conn,
			Ctx:        ctx,
			CancelFunc: cancel,
		}

		// Handle the connection in a new goroutine
		go s.handleClient(connection)
	}
}

func (s *Server) handleClient(connection *Connection) {
	log.Printf("Client %s connected.", connection.Conn.RemoteAddr())
Loop:
	for {
		select {
		case <-connection.Ctx.Done():
			break Loop
		default:
			s.Routes(connection)
		}
	}
}

func (s *Server) Routes(connection *Connection) {
	decoder := json.NewDecoder(connection.Conn)
	var req models.Request
	decoder.Decode(&req)

	route := req.Route

	switch route {
	case "salam":
		s.HandleSalam(connection, req)
	}
}
