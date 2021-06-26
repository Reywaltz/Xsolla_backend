package postgres

import (
	"context"
	"errors"
	"os"

	"github.com/jackc/pgx/v4"
)

type DB struct {
	Conn *pgx.Conn
}

func NewDB() (*DB, error) {
	connStr := os.Getenv("CONN_DB")
	if connStr == "" {
		return nil, errors.New("Connection string is not set")
	}
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		return nil, err
	}

	return &DB{Conn: conn}, nil
}
