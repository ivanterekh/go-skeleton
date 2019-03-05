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
		log.Fatal("PORT must be provided as environment variable")
	}

	server.Start(":" + port)
	closer.Hold()
}
