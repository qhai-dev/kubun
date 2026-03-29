package database

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func NewDatabase() (*Database, error) {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, "postgres://postgres:qwer1234@14.103.230.114:5432/foundation")
	if err != nil {
		return nil, err
	}

	return  &Database{
		Conn: conn,
	}, nil
}

type Database struct {
	Conn *pgx.Conn
}