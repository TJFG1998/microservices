package sub

import (
	"log"

	"golang.org/x/net/context"
)

type Server struct {
}

func (s *Server) Sub2Numbers(ctx context.Context, request *Request) (*Response, error) {
	var sub float32 = request.Number1 - request.Number2

	log.Printf("[SERVER][LOGOING] RECIEVED: %f - %f = %f", request.Number1, request.Number2, sub)

	return &Response{Sub: sub}, nil
}
