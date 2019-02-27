package main

import (
	"log"

	"github.com/ivanterekh/go-skeleton/server"
)

func main() {
	if err := server.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
