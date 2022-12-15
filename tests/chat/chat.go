package chat

import (
	"log"

	"golang.org/x/net/context"
)

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, message *Message) (*Message, error) {
	log.Printf("[SERVER][LOGOING] RECIEVED: %s", &message.BODY)
	return &Message{Body: "HELLO MDF"}, nil
}
