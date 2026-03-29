package rpc

import (
	v1 "github.com/qhai-dev/ozma/api/foundation/v1"
	"google.golang.org/grpc"
)

func NewGRPCServer(so grpc.ServerOption, userServer *UserServer) *grpc.Server {
	s := grpc.NewServer(so)

	v1.RegisterUserServiceServer(s, userServer)
	return s
}
