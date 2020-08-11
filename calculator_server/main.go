//Package main implements a server for Calculator service.

package main

import (
	"context"
	"log"
	"net"

	pb "github.com/craftizmv/calculator-app/calculator"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedCalculatorServer
}

func (s *server) AddNumbers(ctx context.Context, in *pb.RequestData) (*pb.ResponseData, error) {
	log.Printf("Received 1 %v", in.GetNum1())
	log.Printf("Received 2 %v", in.GetNum2())

	//Add the numbers
	output := in.GetNum1() + in.GetNum2()

	return &pb.ResponseData{Result: output}, nil
}

func (s *server) MultiplyNumbers(ctx context.Context, in *pb.RequestData) (*pb.ResponseData, error) {
	log.Printf("Received %v", in.GetNum1())
	log.Printf("Received %v", in.GetNum2())

	//Add the numbers
	output := in.GetNum1() * in.GetNum2()

	return &pb.ResponseData{Result: output}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterCalculatorServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
