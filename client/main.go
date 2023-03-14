package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	greet "gorpc/greet"
	table "gorpc/table"
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
	case "table":
		tableClient()

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

func tableClient() {
	conn, err := grpc.Dial(tableService, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	log.Println("Connected to table service")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	c := table.NewTableClient(conn)

	var input int32
	fmt.Printf("Enter number to generate table: ")
	fmt.Scanln(&input)

	req := &table.Request{Num: int32(input)}

	stream, err := c.Times(ctx, req)
	if err != nil {
		log.Fatalf("Could not receive response from service: %v", err)
	}
	for {
		times, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.Times(_) = _, %v", c, err)
		}
		log.Println(times)
	}

}
