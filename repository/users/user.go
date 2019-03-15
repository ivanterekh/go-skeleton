package users

import (
	"github.com/pkg/errors"

	"github.com/ivanterekh/go-skeleton/model"
)

// ErrNoSuchUser is the error returned by
// functions that search for users.
var ErrNoSuchUser = errors.New("no such user")

// UserRepository is an interface for user searching.
type UserRepository interface {
	// GetByCreds should return user with given credentials
	// or ErrNoSuchUser if it doesn't exist.
	GetByCreds(email, password string) (*model.User, error)

	// GetByID should return user with given id
	// or ErrNoSuchUser if it doesn't exist.
	GetByID(int) (*model.User, error)
}
