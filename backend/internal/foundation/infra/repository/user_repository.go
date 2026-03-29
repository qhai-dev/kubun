package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/qhai-dev/ozma/foundation/infra/database"
)

type UserRepository struct {
	db *database.Database
}

func NewUserRepository(db *database.Database) *UserRepository {
	return &UserRepository {
		db: db,
	}
}

func (u *UserRepository) Save(ctx context.Context, name, phone, hashPassword string) (int64, error) {
	defer u.db.Conn.Close(ctx)

	queries :=database.New(u.db.Conn)

	user, err :=queries.CreateUser(ctx, database.CreateUserParams{
		Name: name,
		Phone: pgtype.Text{String: phone, Valid: true},
		HashPassword: pgtype.Text{String: hashPassword, Valid: true},
	})
	if err != nil {
		return 0, err
	}

	return user.ID, nil
}