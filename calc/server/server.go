package main

import (
	"log"
	"net"

	pb "github.com/go_cs_training/calc"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":33333"
)

type calculatorServer struct{}

func (s *calculatorServer) Add(ctx context.Context, in *pb.AddRequest) (*pb.AddReply, error) {
	return &pb.AddReply{Result: in.Value1 + in.Value2}, nil
}

func (s *calculatorServer) Subtract(ctx context.Context, in *pb.SubtractRequest) (*pb.SubtractReply, error) {
	return &pb.SubtractReply{Result: in.Value1 - in.Value2}, nil
}

func (s *calculatorServer) Multiply(ctx context.Context, in *pb.MultiplyRequest) (*pb.MultiplyReply, error) {
	return &pb.MultiplyReply{Result: in.Value1 * in.Value2}, nil
}

func (s *calculatorServer) Divide(ctx context.Context, in *pb.DivideRequest) (*pb.DivideReply, error) {
	return &pb.DivideReply{Result: float32(in.Value1) / float32(in.Value2)}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterCalculatorServer(s, &calculatorServer{})
	//Register reflection service on grPc server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
}
