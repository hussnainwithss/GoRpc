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
	sum "gorpc/sum"
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
	case "sum":
		sumClient()

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

// TODO: Fix format issues happening because of the go routine runnings
func sumClient() {
	var rpcMethod int
	var reqNum int32
	conn, err := grpc.Dial(sumService, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	log.Println("Connected to Sum service")

	c := sum.NewSumClient(conn)

	fmt.Println("Options:")
	fmt.Println("1: Take list of numbers as input and sum them")
	fmt.Println("2: Continuous sum client send number its sum is returned back and so on")
	fmt.Printf("Please choose what kind of sum functionality you want?: ")
	fmt.Scanln(&rpcMethod)

	if rpcMethod == 1 {
		fmt.Println("---Great you choice option 1---")
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		stream, err := c.SumNumbers(ctx)
		if err != nil {
			log.Fatalf("%v.SumNumbers(_) = _, %v", c, err)
		}

		for reqNum != -9999 {
			fmt.Printf("Enter Number to Sum (-9999 to exit): ")
			fmt.Scanln(&reqNum)
			req := &sum.RequestNumber{Num: reqNum}
			if reqNum != -9999 {
				if err := stream.Send(req); err != nil {
					log.Fatalf("%v.Send(%v) = %v", stream, reqNum, err)
				}
			}

		}
		reply, err := stream.CloseAndRecv()
		if err != nil {
			log.Fatalf("%v.CloseAndRecv() got error %v, want %v", stream, err, nil)
		}
		log.Printf("Final Sum: %d", reply.GetSum())

	} else if rpcMethod == 2 {
		fmt.Println("---Great you choice option 2---")
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		stream, err := c.ContinuousSum(ctx)
		if err != nil {
			log.Fatalf("%v.ContinuousSum(_) = _, %v", c, err)
		}

		waitc := make(chan struct{})
		go func() {
			for {
				in, err := stream.Recv()
				if err == io.EOF {
					// read done.
					close(waitc)
					return
				}
				if err != nil {
					log.Fatalf("Failed to receive Sum : %v", err)
				}
				log.Printf("Got Sum %v\n", in.GetSum())
			}
		}()

		for reqNum != -9999 {
			fmt.Println("Enter Number to Sum (-9999 to exit): ")
			fmt.Scan(&reqNum)
			req := &sum.RequestNumber{Num: reqNum}
			if reqNum != -9999 {
				if err := stream.Send(req); err != nil {
					log.Fatalf("%v.Send(%v) = %v", stream, reqNum, err)
				}
			}
		}
		stream.CloseSend()
		<-waitc
	} else {
		log.Fatalln("Invalid Choice! Exiting")
	}

}
