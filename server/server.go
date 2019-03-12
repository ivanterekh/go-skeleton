package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/ivanterekh/go-skeleton/server/middleware"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

const shutdownTimeout = 5 * time.Second

// Run listens given address.
func Run(ctx context.Context, listenAddr string, logger *zap.Logger) {
	server := http.Server{
		Addr:    listenAddr,
		Handler: setupRouter(logger),
	}

	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			logger.Fatal("server stopped", zap.Error(err))
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	select {
	case <-quit:
	case <-ctx.Done():
	}
	logger.Info("shutdown server")

	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		logger.Error("shutdown error", zap.Error(err))
	}
}

func setupRouter(logger *zap.Logger) *gin.Engine {
	router := gin.New()

	router.Use(middleware.Logging(logger))
	router.Use(middleware.Recovery)

	router.GET("/", hello)
	router.GET("/error", err)
	router.GET("/panic", panicHandler)

	return router
}
