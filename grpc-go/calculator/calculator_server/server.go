package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/emwp/go-studies/grpc-go/calculator/calculatorpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	firstNumber := req.FirstNumber
	secondNumber := req.GetSecondNumber()

	res := &calculatorpb.SumResponse{
		SumResult: firstNumber + secondNumber,
	}
	return res, nil
}

func main() {
	fmt.Println("Calculator Server Started!")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
}
