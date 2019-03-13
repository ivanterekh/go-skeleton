package db

import (
	"database/sql"
	"fmt"

	// Imported for using postgres driver.
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"go.uber.org/zap"

	"github.com/ivanterekh/go-skeleton/env"
)

var (
	db  *sql.DB
	log *zap.Logger
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

// Init creates db instance and pings it.
func Init(logger *zap.Logger) {
	log = logger

	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s/%s?sslmode=%s",
		cfg.user,
		cfg.password,
		cfg.address,
		cfg.name,
		cfg.sslMode,
	)

	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("could not connect to database",
			zap.Error(err),
		)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("could not ping database",
			zap.Error(err),
		)
	}

	log.Info("initialized database",
		zap.String("address", cfg.address),
		zap.String("dbName", cfg.name),
		zap.String("sslMode", cfg.sslMode),
	)
}

// SelectOne performs "SELECT 1;" sql to check
// if db works and connection is ok.
func SelectOne() error {
	rows, err := db.Query("SELECT 1")
	if err != nil {
		return err
	}

	if !rows.Next() {
		return errors.New("select 1 did not return any rows")
	}

	var one int
	if err := rows.Scan(&one); err != nil {
		return err
	}
	if one != 1 {
		return fmt.Errorf("select 1 did return: %v, expected: 1", one)
	}

	return nil
}
