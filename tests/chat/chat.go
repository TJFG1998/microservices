package chat

import (
	"log"

	"golang.org/x/net/context"
)

type Server struct {
}

func (s *Server) Sum2Numbers(ctx context.Context, message *Message) (*Response, error) {
	var sum float32 = message.Number1 + message.Number2

	log.Printf("[SERVER][LOGOING] RECIEVED: %f + %f = %f", message.Number1, message.Number2, sum)

	return &Response{Sum: sum}, nil
}
