package main

import (
	"github.com/ivanterekh/go-skeleton/server"
	"github.com/xlab/closer"
	"go.uber.org/zap"
)

func main() {
	server.Start(":8080")
	closer.Hold()
}

func NewLogger() (*zap.SugaredLogger, error) {
	log, err := zap.NewDevelopment()
	if err != nil {
		return nil, err
	}
	return log.Sugar(), nil
}
