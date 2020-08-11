package main

import (
	"context"
	"log"
	"time"

	pb "github.com/craftizmv/calculator-app/calculator"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	//Setup the connection with the server

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewCalculatorClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.AddNumbers(ctx, &pb.RequestData{Num1: 2, Num2: 3})
	if err != nil {
		log.Fatalf("could not add the numbers: %v", err)
	}
	log.Printf("Addition Result is: %d", r.GetResult())

	r, err = c.MultiplyNumbers(ctx, &pb.RequestData{Num1: 2, Num2: 3})
	if err != nil {
		log.Fatalf("could not multiply the numbers: %v", err)
	}
	log.Printf("Multiplication Result is: %d", r.GetResult())

}
