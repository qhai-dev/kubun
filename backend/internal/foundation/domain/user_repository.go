package domain

import "context"

type UserRepository interface {
	Save(ctx context.Context, name, phone, hashPassword string) (int64, error)
}
