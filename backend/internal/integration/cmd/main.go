package main

import (
	"fmt"
	"net"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", "localhost:8003")
	if err != nil {
		fmt.Println(err)
		return
	}

	s := grpc.NewServer()

	if err := s.Serve(lis); err != nil {
		fmt.Println(err)
	}
}
