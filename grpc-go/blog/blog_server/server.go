package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/emwp/go-studies/grpc-go/blog/blogpb"
	"google.golang.org/grpc"
)

type server struct{}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	fmt.Println("Blog Service Started!")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{}

	s := grpc.NewServer(opts...)
	blogpb.RegisterBlogServiceServer(s, &server{})

	go func() {
		fmt.Println("Starting Server...")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Failed to Server: %v", err)
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	<-ch
	fmt.Println("\nStopping the server")
	s.Stop()
	fmt.Println("Closing the listener")
	lis.Close()
	fmt.Println("End of program")
}
