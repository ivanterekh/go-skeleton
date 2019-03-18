package users

import (
	"database/sql"

	"github.com/pkg/errors"
)

const (
	selectByID = `
SELECT id, name, role, email, password
FROM "USER"
WHERE id = $1`

	selectByCreds = `
SELECT id, name, role, email, password
FROM "USER"
WHERE email = $1 AND password = $2`
)

// PostgresRepository is implementation of
// UserRepository for postgres database.
type PostgresRepository struct {
	db *sql.DB
}

// NewPostgresRepository inits a new instance
// of PostgresRepository with provided db
// instance.
func NewPostgresRepository(db *sql.DB) Repository {
	return &PostgresRepository{
		db: db,
	}
}

// GetByID returns user with provided id
// queried from the postgres db or returns
// ErrNoSuchUser if such user is not in
// database.
func (r *PostgresRepository) GetByID(id int) (*User, error) {
	row := r.db.QueryRow(selectByID, id)
	return scanUser(row)
}

// GetByCreds returns user with provided
// credentials queried from the postgres db
// or returns ErrNoSuchUser if such user is
// not in database.
func (r *PostgresRepository) GetByCreds(email, password string) (*User, error) {
	row := r.db.QueryRow(selectByCreds, email, password)
	return scanUser(row)
}

func scanUser(row *sql.Row) (*User, error) {
	user := &User{}
	err := row.Scan(
		&user.ID,
		&user.Name,
		&user.Role,
		&user.Email,
		&user.Password,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNoSuchUser
		}
		return nil, errors.Wrap(err, "could not select user from postgres db")
	}

	return user, nil
}
