package main

import (
	"log"
	"os"

	"github.com/ivanterekh/go-skeleton/server"
)

func main() {
	address := os.Getenv("ADDRESS")
	if address == "" {
		address = ":8080"
	}

	log.Printf("starting server on %v", address) // TODO: change logging
	server.Run(address)
}
