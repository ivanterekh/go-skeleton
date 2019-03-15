package server

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	"github.com/ivanterekh/go-skeleton/auth"
	// TODO: import without custom package name after handlers reorganization
	globalEnv "github.com/ivanterekh/go-skeleton/env"
	"github.com/ivanterekh/go-skeleton/model"
	"github.com/ivanterekh/go-skeleton/repository/users"
	"github.com/ivanterekh/go-skeleton/version"
)

type env struct {
	db   *sql.DB
	auth *auth.Authenticator
}

func (*env) helloHandler(c *gin.Context) {
	c.String(http.StatusOK, "Hello, world!")
}

func (*env) errorHandler(c *gin.Context) {
	c.Error(errors.New("some error"))
	c.String(http.StatusInternalServerError, "some error")
}

func (*env) panicHandler(c *gin.Context) {
	panic(errors.New("some error"))
}

func (e *env) dbCheckHandler(c *gin.Context) {
	if err := selectOne(e.db); err != nil {
		c.Error(errors.Wrap(err, "could not check db with select 1"))
		c.String(http.StatusInternalServerError, "error")
		return
	}

	c.String(http.StatusOK, "Ok")
}

func selectOne(db *sql.DB) error {
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
		return errors.Errorf("select 1 did return: %v, expected: 1", one)
	}

	return nil
}

func (e *env) healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{
		"version":   version.Version,
		"commit":    version.Commit,
		"buildTime": version.BuildTime,
	})
}

func (e *env) loginHandler(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	token, err := e.auth.GenToken(email, password)
	if err != nil {
		if err == users.ErrNoSuchUser {
			c.String(http.StatusUnauthorized, "wrong credentials")
			return
		}
		c.Error(errors.Wrap(err, "could not generate token"))
		return
	}

	setJWT(c, token, int(e.auth.Exp()/time.Second))
	c.Redirect(http.StatusFound, "/example/private")
}

func (e *env) logoutHandler(c *gin.Context) {
	deleteJWT(c)
	c.Redirect(http.StatusFound, "/")
}

func deleteJWT(c *gin.Context) {
	c.SetCookie(
		"jwt",
		"",
		-1,
		"/",
		globalEnv.GetString("DOMAIN", ""),
		true,
		false)
}

func setJWT(c *gin.Context, token string, maxAge int) {
	c.SetCookie(
		"jwt",
		token,
		maxAge,
		"/",
		globalEnv.GetString("DOMAIN", ""),
		true,
		false)
}

func (e *env) privateHandler(c *gin.Context) {
	userValue, ok := c.Get("user")
	if !ok {
		c.Error(errors.New("could not get user from context"))
		return
	}

	user, ok := userValue.(*model.User)
	if !ok {
		c.Error(errors.New("user value in context has invalid type"))
		return
	}

	c.String(http.StatusOK, "Hello, %s!", user.Name)
}
