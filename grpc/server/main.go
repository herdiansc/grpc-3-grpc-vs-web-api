package main

import (
	"context"
	"google.golang.org/grpc"
	pb "grpc-3-grpc-vs-web-api/proto"
	"log"
	"net"
)

type server struct {
	pb.UnimplementedHelloWorldServer
}

func (s *server) SayHello(ctx context.Context, req *pb.HelloWorldRequest) (*pb.HelloWorldResponse, error) {
	return &pb.HelloWorldResponse{Content:"Hello " + req.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed listen: %v", err.Error())
	}
	log.Printf("grpc server listening on %v", lis.Addr())
	s := grpc.NewServer()
	pb.RegisterHelloWorldServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed serve: %v", err.Error())
	}
}
