package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

// Recovery is middleware for panic recovery. If panic occurs
// it attaches error to request processing chain.
func Recovery(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.Error(errors.Errorf("panic recovered: %v", err))
			c.String(http.StatusInternalServerError, "server error")
		}
	}()
	c.Next()
}
