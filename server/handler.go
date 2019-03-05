package server

import (
	"github.com/ivanterekh/go-skeleton/environment"
	"net/http"

	"github.com/gin-gonic/gin"
)

func hello(c *gin.Context) {
	c.String(http.StatusOK, "Hello, world!")
}

func env(c *gin.Context) {
	c.String(http.StatusOK, environment.Get())
}
