package rpc

import (
	"context"

	v1 "github.com/qhai-dev/ozma/api/foundation/v1"
	"github.com/qhai-dev/ozma/foundation/application"
)

type UserServer struct {
	v1.UnimplementedUserServiceServer

	createHandler *application.CreateUserHandler
}

func NewUserServer(createHandler *application.CreateUserHandler) *UserServer {
	return &UserServer{
		createHandler: createHandler,
	}
}

func (s *UserServer) CreateUser(ctx context.Context, req *v1.CreateUserRequest) (*v1.CreateUserResponse, error) {
	 id , err :=s.createHandler.Handler(ctx, &application.CreateUserCommand{
		Name: req.GetName(),
		Phone: req.GetPhone(),
		HashPassword: "123456789",
	})

	if err != nil {
		return nil, err
	}

	return &v1.CreateUserResponse{
		Id: id,
	}, nil
}

