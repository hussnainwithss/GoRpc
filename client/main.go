package main

import (
	"context"
	"flag"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	greet "gorpc/greet"
)

const (
	greetService   = "localhost:3000"
	sumService     = "localhost:3001"
	tableService   = "localhost:3002"
	defaultService = "greet"
)

var (
	service = flag.String("service", defaultService, "Service to call, options are greet, sum, table")
)

func main() {

	flag.Parse()

	switch *service {
	case "greet":
		greetClient()

	}

}

func greetClient() {
	conn, err := grpc.Dial(greetService, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := greet.NewGreetClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := c.SayGreetings(ctx, &greet.Request{})
	if err != nil {
		log.Fatalf("Could not receive response from service: %v", err)
	}

	log.Printf("Response: %s", response.Greeting)
}
