package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	"github.com/ivanterekh/go-skeleton/version"
)

func helloHandler(c *gin.Context) {
	c.String(http.StatusOK, "Hello, world!")
}

func errorHandler(c *gin.Context) {
	c.Error(errors.New("some error"))
	c.String(http.StatusInternalServerError, "some error")
}

func panicHandler(c *gin.Context) {
	panic(errors.New("some error"))
}

func healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{
		"version":   version.Version,
		"commit":    version.Commit,
		"buildTime": version.BuildTime,
	})
}
