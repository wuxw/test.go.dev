package server

import (
	pb "test.go.dev/grpc-hello-world/proto"

	"golang.org/x/net/context"
)

type userService struct{}

func NewUserService() *userService {
	return &userService{}
}
func (h userService) SayUserFind(ctx context.Context, r *pb.UserFindRequest) (*pb.UserFindResponse, error) {
	return &pb.UserFindResponse{
		Message: "userfind",
	}, nil
}
