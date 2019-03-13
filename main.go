package main

import (
	"context"
	"log"

	"github.com/ivanterekh/go-skeleton/env"
	"github.com/ivanterekh/go-skeleton/server"
	"github.com/ivanterekh/go-skeleton/version"
	"github.com/pkg/errors"

	"go.uber.org/zap"
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

	address := env.GetString("ADDRESS", ":8080")

	logger.Info("starting server", zap.String("address", address))
	server.Run(context.Background(), address, logger)
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
