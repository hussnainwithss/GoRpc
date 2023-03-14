package main

import (
	"fmt"
	"io"
	"log"
	"net"

	sum "gorpc/sum"

	"google.golang.org/grpc"
)

type SumationServer struct {
	sum.UnimplementedSumServer
}

func (s *SumationServer) SumNumbers(stream sum.Sum_SumNumbersServer) error {
	var finalSum int32
	for {
		inputNum, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&sum.FinalSum{
				Sum: finalSum,
			})
		}
		if err != nil {
			return err
		}
		finalSum += inputNum.GetNum()
	}
}

func (s *SumationServer) ContinuousSum(stream sum.Sum_ContinuousSumServer) error {
	var finalSum int32
	for {
		inputNum, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		finalSum += inputNum.GetNum()
		if err := stream.Send(&sum.FinalSum{
			Sum: finalSum,
		}); err != nil {
			return err
		}
	}
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 3001))
	if err != nil {
		log.Fatalf("Failed to listen to port: %v", err)
	}

	sumServer := grpc.NewServer()

	sum.RegisterSumServer(sumServer, &SumationServer{})
	log.Printf("Sumation service listeing at %v", lis.Addr())

	if err := sumServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
