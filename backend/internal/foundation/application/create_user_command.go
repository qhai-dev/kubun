package application

import (
	"context"
	"fmt"

	"github.com/qhai-dev/ozma/foundation/domain"
	"github.com/qhai-dev/ozma/foundation/infra/repository"
)

type CreateUserCommand struct {
    Name string
    Phone    string
	HashPassword string
}

type CreateUserHandler struct {
	repo domain.UserRepository
}

func NewCreateUserHandler(repo *repository.UserRepository) *CreateUserHandler {
	return &CreateUserHandler{
		repo: repo,
	}
}

func (c *CreateUserHandler) Handler(ctx context.Context, cmd *CreateUserCommand) (int64, error) {

	fmt.Printf("CreateUserCommand: Name=%s, Phone=%s, HashPassword=%s\n", cmd.Name, cmd.Phone, cmd.HashPassword)

	id, err := c.repo.Save(ctx, cmd.Name, cmd.Phone, cmd.HashPassword)
	if err != nil {
		return 0, err
	}

	fmt.Printf("id %d", id)

	return id, nil
}