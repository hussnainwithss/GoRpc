package main

import (
	"fmt"
	"log"
	"net"

	table "gorpc/table"

	"google.golang.org/grpc"
)

type TableGeneratorServer struct {
	table.UnimplementedTableServer
}

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func (s *TableGeneratorServer) Times(in *table.Request, stream table.Table_TimesServer) error {
	num := in.GetNum()
	log.Printf("got input: %d", num)

	for _, multiplier := range makeRange(1, 10) {
		multiple := table.Result{Result: fmt.Sprintf(" %d x %d = %d", num, multiplier, num*int32(multiplier))}
		if err := stream.Send(&multiple); err != nil {
			return err
		}
	}
	return nil

}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 3002))
	if err != nil {
		log.Fatalf("Failed to listen to port: %v", err)
	}

	tableServer := grpc.NewServer()

	table.RegisterTableServer(tableServer, &TableGeneratorServer{})
	log.Printf("table service listeing at %v", lis.Addr())

	if err := tableServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
