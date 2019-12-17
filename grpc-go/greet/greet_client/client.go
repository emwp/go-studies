package main

import (
	"fmt"
	"github.com/emwp/go-studies/grpc-go/greet/greetpb"
	"google.golang.org/grpc"
	"log"
)

func main() {
	fmt.Println("Hello I'm a client!")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	defer conn.Close()

	c := greetpb.NewGreetServiceClient(conn)
	fmt.Printf("gRPC Client Created: %f", c)
}
