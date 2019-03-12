package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-contrib/zap"
	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

const shutdownTimeout = 1 * time.Second

// Run listens given address.
func Run(listenAddr string, logger *zap.Logger) {
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
	<-quit
	logger.Info("shutdown server")

	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		logger.Error("shutdown error", zap.Error(err))
	}
}

func setupRouter(logger *zap.Logger) *gin.Engine {
	router := gin.New()
	router.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	router.GET("/", hello)
	return router
}
