package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Logging writes request results.
func Logging(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		end := time.Now()
		latency := end.Sub(start)

		requestInfo := []zap.Field{
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.String("query", c.Request.URL.RawQuery),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.Duration("latency", latency),
		}

		if errors := c.Errors.Errors(); len(errors) > 0 {
			requestInfo = append(
				[]zap.Field{zap.Strings("errors", errors)},
				requestInfo...,
			)
			logger.Error("request error", requestInfo...)
		} else {
			logger.Info("request processed", requestInfo...)
		}
	}
}
