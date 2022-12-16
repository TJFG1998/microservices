package main

import (
	"log"

	"github.com/TJFG1998/microservices/div"
	"github.com/TJFG1998/microservices/mul"
	"github.com/TJFG1998/microservices/sub"
	"github.com/TJFG1998/microservices/sum"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func test_sum() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()
	c := sum.NewSumServiceClient(conn)

	response, err := c.Sum2Numbers(
		context.Background(), &sum.Request{Number1: 12.4, Number2: 1.1})

	if err != nil {
		log.Fatalf("[ERROR][SUM-SERVICE]: %s", err)
		return
	}
	log.Printf("[RESPONSE][SUM-SERVICE]: %f", response.Sum)
}

func test_sub() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()
	c := sub.NewSubServiceClient(conn)

	response, err := c.Sub2Numbers(
		context.Background(), &sub.Request{Number1: 12.4, Number2: 1.1})

	if err != nil {
		log.Fatalf("[ERROR][SUB-SERVICE]: %s", err)
		return
	}
	log.Printf("[RESPONSE][SUB-SERVICE]: %f", response.Sub)
}

func test_mul() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9002", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()
	c := mul.NewMulServiceClient(conn)

	response, err := c.Mul2Numbers(
		context.Background(), &mul.Request{Number1: 12.4, Number2: 1.1})

	if err != nil {
		log.Fatalf("[ERROR][MUL-SERVICE]: %s", err)
		return
	}
	log.Printf("[RESPONSE][MUL-SERVICE]: %f", response.Mul)
}

func test_div() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9003", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := div.NewDivServiceClient(conn)

	response, err := c.Div2Numbers(
		context.Background(), &div.Request{Number1: 12.4, Number2: 1.1})

	if err != nil {
		log.Fatalf("[ERROR][DIV-SERVICE]: %s", err)
		return
	}
	log.Printf("[RESPONSE][DIV-SERVICE]: %f", response.Div)
}

func main() {
	test_sum()
	test_sub()
	test_div()
	test_mul()
}
