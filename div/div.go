package div

import (
	"log"

	"golang.org/x/net/context"
)

type Server struct {
}

func (s *Server) Div2Numbers(ctx context.Context, request *Request) (*Response, error) {
	var div float32 = request.Number1 / request.Number2

	log.Printf("[SERVER][LOGOING] RECIEVED: %f / %f = %f", request.Number1, request.Number2, div)

	return &Response{Div: div}, nil
}
