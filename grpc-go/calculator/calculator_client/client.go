package main

import (
	"context"
	"fmt"
	"log"

	"github.com/emwp/go-studies/grpc-go/calculator/calculatorpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Calculator Client")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	defer conn.Close()

	c := calculatorpb.NewCalculatorServiceClient(conn)

	doUnary(c)
}

func doUnary(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting Sum Unary RPC...")
	req := &calculatorpb.SumRequest{
		FirstNumber:  4444,
		SecondNumber: 889,
	}

	res, err := c.Sum(context.Background(), req)

	if err != nil {
		log.Fatalf("There was an error while calling Sum RPC: %v", err)
	}

	log.Printf("Response from Sum: %v", res.SumResult)
}
