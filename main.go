package main

import (
	"context"
	"log"

	"github.com/ivanterekh/go-skeleton/env"
	"github.com/ivanterekh/go-skeleton/server"
	"go.uber.org/zap"
)

func main() {
	logger, err := newLogger()
	if err != nil {
		log.Fatalf("could not init logger: %v", err)
	}

	address := env.GetString("ADDRESS", ":8080")

	logger.Info("starting server", zap.String("address", address))
	server.Run(context.Background(), address, logger)
}

func newLogger() (*zap.Logger, error) {
	if env.IsDev() {
		return zap.NewDevelopment()
	}

	return zap.NewProduction()
}
