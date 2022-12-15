package main

import (
	"log"
	"net"

	"github.com//TJFG1998/microservices/chat"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")

	if err != nil {
		log.Fatalf("[NETWORK][ERORR] %v", err)
	}

	s := chat.Server{}

	grpcServer := grpc.NewServer()

	chat.RegisterChatServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err {
		log.Fatalf("[GRPC][ERORR] %v", err)
	}

}
