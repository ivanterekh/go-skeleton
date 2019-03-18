package db

import (
	"database/sql"
	"fmt"

	// Imported for using postgres driver.
	_ "github.com/lib/pq"
	"github.com/pkg/errors"

	"github.com/ivanterekh/go-skeleton/pkg/env"
)

var cfg = struct {
	user, password, address, name, sslMode string
}{
	user:     env.GetString("DB_USER", "postgres"),
	password: env.GetString("DB_PASSWORD", "password"),
	address:  env.GetString("DB_ADDRESS", "localhost"),
	name:     env.GetString("DB_NAME", "goskeleton"),
	sslMode:  env.GetString("DB_SSL_MODE", "disable"),
}

// New creates db instance and pings it.
func New() (*sql.DB, error) {
	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s/%s?sslmode=%s",
		cfg.user,
		cfg.password,
		cfg.address,
		cfg.name,
		cfg.sslMode,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, errors.Wrap(err, "could not connect to database")
	}

	if err := db.Ping(); err != nil {
		return nil, errors.Wrap(err, "could not ping database")
	}

	return db, nil
}
