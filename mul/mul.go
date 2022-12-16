package mul

import (
	"log"

	"golang.org/x/net/context"
)

type Server struct {
}

func (s *Server) Mul2Numbers(ctx context.Context, request *Request) (*Response, error) {
	var mul float32 = request.Number1 * request.Number2

	log.Printf("[SERVER][LOGOING] RECIEVED: %f - %f = %f", request.Number1, request.Number2, mul)

	return &Response{Mul: mul}, nil
}
