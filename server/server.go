package server

import (
	"context"
	"database/sql"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/ivanterekh/go-skeleton/auth"
	"github.com/ivanterekh/go-skeleton/server/middleware"
)

const shutdownTimeout = 5 * time.Second

// Run listens given address.
func Run(ctx context.Context, listenAddr string, logger *zap.Logger, db *sql.DB) {
	server := http.Server{
		Addr:    listenAddr,
		Handler: setupRouter(logger, db),
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

func setupRouter(logger *zap.Logger, db *sql.DB) *gin.Engine {
	router := gin.New()

	router.Use(middleware.Logging(logger))
	router.Use(middleware.Recovery)

	authenticator := auth.DefaultAuthenticator()

	env := env{
		db:   db,
		auth: authenticator,
	}

	router.StaticFile("/login", "./templates/login.html")
	router.POST("/login", env.loginHandler)
	router.GET("/logout", env.logoutHandler)

	router.GET("/", env.helloHandler)
	router.GET("/health", env.healthHandler)

	example := router.Group("/example")
	example.GET("/error", env.errorHandler)
	example.GET("/panic", env.panicHandler)
	example.GET("/db-check", env.dbCheckHandler)
	example.GET("/private", middleware.Auth(authenticator), env.privateHandler)

	return router
}
