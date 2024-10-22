package server

import (
	"encoding/json"
	"errors"
	"examples/models"
	"fmt"
	"log"
)

func (s *Server) HandleSalam(connection *Connection, req models.Request) error {
	log.Println("handler salam")
	content := req.Content
	encoder := json.NewEncoder(connection.Conn)
	if content == "" {
		res := models.ErrorResponse{
			Status: "nok",
			Error:  errors.New("missing content").Error(),
		}
		encoder.Encode(res)
	}
	res := models.Response{
		Status:  "ok",
		Message: fmt.Sprintf("message received: %s", content),
	}
	encoder.Encode(res)
	return nil
}
