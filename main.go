package main

import (
	"log"
)

func main() {
	if err := runServer(":8080"); err != nil {
		log.Fatal(err)
	}
}
