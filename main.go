package main

import (
	"log"

	"github.com/ivanterekh/go-skeleton/env"
	"github.com/ivanterekh/go-skeleton/server"
)

type config struct {
	Address string `env:"ADDRESS" envDefault:":8080"`
}

func main() {
	address := env.GetString("ADDRESS", ":8080")

	log.Printf("starting server on %v", address) // TODO: change logging
	server.Run(address)
}
