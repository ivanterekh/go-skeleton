package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ivanterekh/go-skeleton/auth"
)

// Auth checks if user agent provided valid jwt.
func Auth(a *auth.Authenticator) gin.HandlerFunc {
	return func(c *gin.Context) {
		jwt, err := c.Cookie("jwt")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, "unauthorized")
			return
		}

		user, err := a.Authenticate(jwt)
		if err != nil {
			c.Error(err)
			return
		}

		c.Set("user", user)
		c.Next()
	}
}
