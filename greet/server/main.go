package main

import (
	"context"
	"fmt"
	"log"
	"net"

	greet "gorpc/greet"

	"google.golang.org/grpc"
)

type GreetServer struct {
	greet.UnimplementedGreetServer
}

func (s *GreetServer) SayGreetings(ctx context.Context, in *greet.Request) (*greet.Greeting, error) {
	log.Printf("Recieved Greet user Request")
	return &greet.Greeting{Greeting: "hello I am the wonderful engineer who joined Tweeq"}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 3000))
	if err != nil {
		log.Fatalf("Failed to listen to port: %v", err)
	}

	greetingServer := grpc.NewServer()

	greet.RegisterGreetServer(greetingServer, &GreetServer{})
	log.Printf("Greeting service listeing at %v", lis.Addr())

	if err := greetingServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
