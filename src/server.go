//go:generate protoc -I ../protofiles --go_out=plugins=grpc:../protofiles ../protofiles/*.proto

package main

import (
	"log"
	"net"
	"fmt"
	
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "../protofiles"
	"google.golang.org/grpc/reflection"
)

const (
	port = "6969"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func (s *server) GetName(ctx context.Context, in *pb.NameRequest) (*pb.NameReply, error) {
	return &pb.NameReply{Message: "Your name is: " + in.Name}, nil
}

func main() {
	fmt.Println("Port is:", port)
	lis, err := net.Listen("tcp",":" + port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	pb.RegisterCharacterStatsServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
