package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func hello(c *gin.Context) {
	c.String(http.StatusOK, "Hello, world!")
}
