package server

import (
	pb "test.go.dev/grpc-hello-world/proto"

	"golang.org/x/net/context"
)

type helloService struct{}

func NewHelloService() *helloService {
	return &helloService{}
}
func (h helloService) SayHelloWorld(ctx context.Context, r *pb.HelloWorldRequest) (*pb.HelloWorldResponse, error) {
	return &pb.HelloWorldResponse{
		Message: "test12",
	}, nil
}
