package postgres

import "github.com/jackc/pgx"

const (
	connstr = `postgres://xsolla_user:qwerty@localhost:5433/xsolla`
)

type DB struct {
	Conn *pgx.Conn
}

func NewDB() (*DB, error) {
	cfg, err := pgx.ParseConnectionString(connstr)
	if err != nil {
		return nil, err
	}
	conn, err := pgx.Connect(cfg)
	if err != nil {
		return nil, err
	}

	return &DB{Conn: conn}, nil
}
