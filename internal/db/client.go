package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// DB is a struct containing the Sqlx connection instance.
type DB struct {
	Sqlx *sqlx.DB
}

// New creates a new database connection using the supplied PostgreSQL
// connection string.
func New(dsn string) (*DB, error) {
	conn, err := sqlx.Connect("postgres", dsn)

	if err != nil {
		return nil, err
	}

	return &DB{conn}, nil
}
