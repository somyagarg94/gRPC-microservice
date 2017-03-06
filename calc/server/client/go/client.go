package main

import (
	"log"

	pb "github.com/go_cs_training/calc"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address = "localhost:33333"
)

func main() {
	//set up the connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewCalculatorClient(conn)

	r, err := c.Add(context.Background(), &pb.AddRequest{Value1: 67, Value2: 133})
	if err != nil {
		log.Fatalf("could not add: %v", err)
	}
	log.Printf("Adding: %d", r.Result)
}
