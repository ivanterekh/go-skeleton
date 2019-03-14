package server

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	"github.com/ivanterekh/go-skeleton/version"
)

type env struct {
	db *sql.DB
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
		return fmt.Errorf("select 1 did return: %v, expected: 1", one)
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

func loginHandler(c *gin.Context) {
	user := c.Param("user")
	password := c.Param("password")
}
