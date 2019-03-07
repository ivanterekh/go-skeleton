package main

import (
	"log"
	"os"

	"github.com/ivanterekh/go-skeleton/server"
	"github.com/xlab/closer"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("starting server on port %v", port) // TODO: change logging
	server.Start(":" + port)
	closer.Hold()
}
