package postgres

import (
	"context"

	"github.com/jackc/pgx/v4"
)

const (
	connstr = `postgres://xsolla_user:qwerty@localhost:5433/xsolla`
)

type DB struct {
	Conn *pgx.Conn
}

func NewDB() (*DB, error) {
	conn, err := pgx.Connect(context.Background(), connstr)
	if err != nil {
		return nil, err
	}

	return &DB{Conn: conn}, nil
}
