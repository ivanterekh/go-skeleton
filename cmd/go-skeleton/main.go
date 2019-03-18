package main

import (
	"context"
	"log"

	"github.com/pkg/errors"
	"go.uber.org/zap"

	"github.com/ivanterekh/go-skeleton/internal/db"
	"github.com/ivanterekh/go-skeleton/internal/env"
	"github.com/ivanterekh/go-skeleton/internal/server"
	"github.com/ivanterekh/go-skeleton/internal/version"
)

func main() {
	logger, err := newLogger()
	if err != nil {
		log.Fatal(errors.Wrap(err, "could not init logger"))
	}

	logger.Info("running app",
		zap.String("version", version.Version),
		zap.String("commit", version.Commit),
		zap.String("buildTime", version.BuildTime),
	)

	newDb, err := db.New()
	if err != nil {
		log.Fatal("could not create db instance", zap.Error(err))
	}

	address := env.GetString("ADDRESS", ":8080")
	logger.Info("starting server", zap.String("address", address))
	server.Run(context.Background(), address, logger, newDb)
}

func newLogger() (*zap.Logger, error) {
	var cfg zap.Config
	if env.IsDev() {
		cfg = zap.NewDevelopmentConfig()
	} else {
		cfg = zap.NewProductionConfig()
	}

	cfg.DisableStacktrace = true
	cfg.DisableCaller = true
	return cfg.Build()
}
