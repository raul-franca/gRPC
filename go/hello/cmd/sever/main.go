package main

import (
	"context"
	"fmt"
	"github.com/raul-franca/gRPC/go-gRPC-hello/pb"
	"google.golang.org/grpc"
	"net"
)

type Sever struct {
	pb.UnimplementedHelloServer
}

func (s Sever) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Message: "Hello, " + in.Name}, nil
}

func main() {
	fmt.Println("Running gRPC server")

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	pb.RegisterHelloServer(s, &Sever{})
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
