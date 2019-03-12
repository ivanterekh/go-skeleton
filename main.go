package main

import (
	"log"

	"github.com/ivanterekh/go-skeleton/env"
	"github.com/ivanterekh/go-skeleton/server"
	"go.uber.org/zap"
)

func main() {
	logger, err := initLogger()
	if err != nil {
		log.Fatalf("could not init logger: %v", err)
	}

	address := env.GetString("ADDRESS", ":8080")

	logger.Info("starting server", zap.String("address", address))
	server.Run(address, logger)
}

func initLogger() (*zap.Logger, error) {
	if env.IsDev() {
		return zap.NewDevelopment()
	}

	return zap.NewProduction()
}
