package auth

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/ivanterekh/go-skeleton/env"
	"github.com/ivanterekh/go-skeleton/model"
	"github.com/pkg/errors"
	"time"
)

var ErrNoSuchUser = errors.New("no such user")

type userRepository interface {
	byCreds(email, password string) (*model.User, error)
}

type mockUserRepository map[string]model.User

func (repo mockUserRepository) byCreds(email, password string) (*model.User, error) {
	if user, ok := repo[email]; ok {
		return &user, nil
	}
	return nil, errors.New("not found")
}

type auth struct {
	exp    time.Duration
	method jwt.SigningMethod
	users  userRepository
}

func NewAuth() *auth {
	return &auth{
		exp:    time.Hour,
		method: jwt.SigningMethodHS256,
		users: mockUserRepository{
			"user1@gmail.com": model.User{
				Email: "user1@gmail.com",
				Name:  "User Friendly",
				Role:  "user",
				ID:    42,
			},
		},
	}
}

func (a *auth) token(u model.User) (string, error) {
	token := jwt.NewWithClaims(a.method, jwt.MapClaims{
		"exp": time.Now().Add(a.exp),
		"sub": u.ID,
	})

	return token.SignedString(env.GetString("AUTH_SECRET", "secret"))
}

func (a *auth) GetUser(email, password string) (string, error) {
	user, err := a.users.byCreds(email, password)
	if err != nil {
		return "", ErrNoSuchUser
	}

	return a.token(*user)
}
