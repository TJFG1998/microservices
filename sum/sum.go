package sum

import (
	"log"

	"golang.org/x/net/context"
)

type Server struct {
}

func (s *Server) Sum2Numbers(ctx context.Context, request *Request) (*Response, error) {
	var sum float32 = request.Number1 + request.Number2

	log.Printf("[SERVER][LOGOING] RECIEVED: %f + %f = %f", request.Number1, request.Number2, sum)

	return &Response{Sum: sum}, nil
}
