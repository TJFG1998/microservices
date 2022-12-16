package main

import (
	"log"
	"net"
	"os"

	"github.com/TJFG1998/microservices/div"
	"github.com/TJFG1998/microservices/mul"
	"github.com/TJFG1998/microservices/sub"
	"github.com/TJFG1998/microservices/sum"

	"google.golang.org/grpc"
)

func sum_server(lis net.Listener, err error) {
	s := sum.Server{}

	grpcServer := grpc.NewServer()

	sum.RegisterSumServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("[GRPC][ERORR] %v", err)
	}
}

func sub_server(lis net.Listener, err error) {
	s := sub.Server{}

	grpcServer := grpc.NewServer()

	sub.RegisterSubServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("[GRPC][ERORR] %v", err)
	}
}

func mul_server(lis net.Listener, err error) {
	s := mul.Server{}

	grpcServer := grpc.NewServer()

	mul.RegisterMulServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("[GRPC][ERORR] %v", err)
	}
}

func div_server(lis net.Listener, err error) {
	s := div.Server{}

	grpcServer := grpc.NewServer()

	div.RegisterDivServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("[GRPC][ERORR] %v", err)
	}
}

func main() {
	lis, err := net.Listen("tcp", os.Args[2])

	if err != nil {
		log.Fatalf("[NETWORK][ERORR] %v", err)
		return
	}

	log.Printf("[%s-SERVICE]: Listening on %s", os.Args[1], os.Args[2])

	if os.Args[1] == "SUM" {
		sum_server(lis, err)
	} else if os.Args[1] == "SUB" {
		sub_server(lis, err)
	} else if os.Args[1] == "MUL" {
		mul_server(lis, err)
	} else if os.Args[1] == "DIV" {
		div_server(lis, err)
	} else {
		log.Printf("[ERROR]: SERIVCE %s IS NOT AVAILABLE", os.Args[1])
	}
}
