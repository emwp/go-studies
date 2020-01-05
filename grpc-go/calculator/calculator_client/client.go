package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/emwp/go-studies/grpc-go/calculator/calculatorpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	fmt.Println("Calculator Client")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	defer conn.Close()

	c := calculatorpb.NewCalculatorServiceClient(conn)

	// doUnary(c)

	// doServerStreaming(c)

	// doClientStreaming(c)

	// doBiDiStreaming(c)

	doErrorUnary(c)
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

func doServerStreaming(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting Prime Decomposition Server Stream RPC...")
	req := &calculatorpb.PrimeNumberDecompositionRequest{
		Number: 55555666666,
	}

	stream, err := c.PrimeNumberDecomposition(context.Background(), req)

	if err != nil {
		log.Fatalf("There was an error while calling PrimeNumberDecomposition RPC: %v", err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Something went wrong: %v", err)
		}

		fmt.Println(res.GetPrimeFactor())
	}
}

func doClientStreaming(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("\nStarting ComputeAverage Client Stream RPC...")

	stream, err := c.ComputeAverage(context.Background())

	if err != nil {
		log.Fatalf("Error while opening stream: %v", err)
	}

	numbers := []int32{5, 2, 27, 33, 49}

	for _, number := range numbers {
		fmt.Printf("Sending Number: %v\n", number)
		stream.Send(&calculatorpb.ComputeAverageRequest{
			Number: number,
		})
		time.Sleep(500 * time.Millisecond)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while receiving response: %v", err)
	}

	fmt.Printf("The Average is: %v\n", res.GetAverage())
}

func doBiDiStreaming(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("\nStarting Find Maximum BiDi Stream RPC...")

	stream, err := c.FindMaximum(context.Background())

	if err != nil {
		log.Fatalf("Error while opening stream and calling FindMaximum: %v", err)
	}

	waitc := make(chan struct{})

	// send go routine
	go func() {
		numbers := []int32{4, 7, 2, 19, 4, 6, 32}
		for _, number := range numbers {
			stream.Send(&calculatorpb.FindMaximumRequest{
				Number: number,
			})
			fmt.Printf("\nSending the number: %v\n", number)
			time.Sleep(500 * time.Millisecond)
		}
		stream.CloseSend()
	}()

	// receive go routine
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Problem while reading server stream: %v", err)
				break
			}

			max := res.GetMaximum()
			fmt.Printf("New maximum value is: %v\n", max)
		}
		close(waitc)
	}()

	<-waitc
}

func doErrorUnary(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("\nStarting Square Root Unary RPC...")

	// correct call
	doErrorCall(c, 144)

	// error call
	doErrorCall(c, -3)
}

func doErrorCall(c calculatorpb.CalculatorServiceClient, n int32) {
	res, err := c.SquareRoot(context.Background(), &calculatorpb.SquareRootRequest{Number: n})

	if err != nil {
		resErr, ok := status.FromError(err)
		if ok {
			// Actual error from gRPC (User Error)
			fmt.Printf("\nError message from server: %v\n", resErr.Message())
			fmt.Println(resErr.Code())
			if resErr.Code() == codes.InvalidArgument {
				fmt.Println("You are probably trying to submit a negative number")
				return
			}
		} else {
			log.Fatalf("Big error calling squareRoot: %v", resErr)
			return
		}
	}

	fmt.Printf("\nResult of square root of %v: %v\n", n, res.GetNumberRoot())
}
