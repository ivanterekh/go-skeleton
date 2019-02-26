package main

import (
	"github.com/ivanterekh/go-skeleton/http"
	"log"
	"os"
)

func main() {
	if err := http.RunServer(":8080"); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
