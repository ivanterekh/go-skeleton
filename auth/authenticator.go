package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"

	"github.com/ivanterekh/go-skeleton/env"
	"github.com/ivanterekh/go-skeleton/model"
	"github.com/ivanterekh/go-skeleton/repository/users"
)

// Authenticator is a service for logging in and
// authenticating users.
type Authenticator struct {
	exp    time.Duration
	method *jwt.SigningMethodHMAC
	secret string
	users  users.Repository
}

// Exp returns expiry time.
func (a *Authenticator) Exp() time.Duration {
	return a.exp
}

// NewAuthenticator returns a new authenticator instance.
func NewAuthenticator() *Authenticator {
	return &Authenticator{
		exp:    time.Hour * 48,
		method: jwt.SigningMethodHS256,
		secret: env.GetString("AUTH_SECRET", "secret"),
		users:  users.NewMock(),
	}
}

// GenToken generates a new token if user with
// provided credentials exists.
func (a *Authenticator) GenToken(email, password string) (string, error) {
	user, err := a.users.ByCreds(email, password)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(a.method, jwt.MapClaims{
		"exp": time.Now().Add(a.exp),
		"sub": user.ID,
	})

	return token.SignedString([]byte(a.secret))
}

// Authenticate returns user if token is valid.
func (a *Authenticator) Authenticate(tokenStr string) (*model.User, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(a.secret), nil
	})
	if err != nil {
		return nil, errors.Wrap(err, "could not parse token")
	}

	if !token.Valid {
		return nil, errors.Errorf("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.Errorf("could not convert claims to jwt.MapClaims")
	}

	userID, ok := claims["sub"].(float64)
	if !ok {
		return nil, errors.Errorf("could not get user id from claims")
	}

	return a.users.ByID(int(userID))
}
