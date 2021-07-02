package postgres

import (
	"context"
	"errors"
	"os"

	"github.com/jackc/pgx/v4"
)

func NewDB() (*pgx.Conn, error) {
	connStr := os.Getenv("CONN_DB")
	if connStr == "" {
		return nil, errors.New("Connection string is not set")
	}
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
