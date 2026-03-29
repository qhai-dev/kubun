package foundation

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

type App struct {
	rpcServer *grpc.Server
}

func NewApp(rpcServer *grpc.Server) *App {
	return  &App{rpcServer: rpcServer}
}


func (a *App) Run() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	if err := a.rpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
		panic(err)
	}
}