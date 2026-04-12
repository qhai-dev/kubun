package rpc

import (
	"net"

	"google.golang.org/grpc"
)

func Run() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()

	if err := s.Serve(lis); err != nil {

	}
}
