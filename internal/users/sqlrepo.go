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

// SQLRepository is implementation of
// UserRepository for postgres database.
type SQLRepository struct {
	db *sql.DB
}

// NewSQLRepository inits a new instance
// of SQLRepository with provided db
// instance.
func NewSQLRepository(db *sql.DB) Repository {
	return &SQLRepository{
		db: db,
	}
}

// GetByID returns user with provided id
// queried from the postgres db or returns
// ErrNoSuchUser if such user is not in
// database.
func (r *SQLRepository) GetByID(id int) (*User, error) {
	row := r.db.QueryRow(selectByID, id)
	return scanUser(row)
}

// GetByCreds returns user with provided
// credentials queried from the postgres db
// or returns ErrNoSuchUser if such user is
// not in database.
func (r *SQLRepository) GetByCreds(email, password string) (*User, error) {
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
