package test

import (
	"encoding/json"
	"examples/models"
	"examples/server"
	"fmt"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServer1(t *testing.T) {
	ln, _ := net.Listen("tcp", ":8080")

	server, _ := server.NewServer(ln)

	go server.Start()

	conn, err := net.Dial("tcp", ":8080")
	assert.NoError(t, err, "server should start successfully")
	req := models.Request{
		Route:   "salam",
		Content: "test content",
	}
	encoder := json.NewEncoder(conn)
	err = encoder.Encode(&req)
	assert.NoError(t, err)

	decoder := json.NewDecoder(conn)
	var res models.Response
	err = decoder.Decode(&res)
	assert.NoError(t, err)

	expectedMessage := fmt.Sprintf("message received: %v", req.Content)
	assert.Equal(t, expectedMessage, res.Message, "content not correct")
}
