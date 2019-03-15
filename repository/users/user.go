package users

import (
	"github.com/pkg/errors"

	"github.com/ivanterekh/go-skeleton/model"
)

// ErrNoSuchUser is the error returned by
// functions that search for users.
var ErrNoSuchUser = errors.Errorf("no such user")

// Repository is an interface for user searching.
type Repository interface {
	ByCreds(email, password string) (*model.User, error)
	ByID(int) (*model.User, error)
}
