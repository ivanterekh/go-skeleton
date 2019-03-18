package users

import (
	"github.com/pkg/errors"
)

// ErrNoSuchUser is the error returned by
// functions that search for users.
var ErrNoSuchUser = errors.New("no such user")

// Repository is an interface for user searching.
type Repository interface {
	// GetByCreds should return user with given credentials
	// or ErrNoSuchUser if it doesn't exist.
	GetByCreds(email, password string) (*User, error)

	// GetByID should return user with given id
	// or ErrNoSuchUser if it doesn't exist.
	GetByID(int) (*User, error)
}
