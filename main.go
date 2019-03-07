package main

import (
	"github.com/xlab/closer"

	"github.com/ivanterekh/go-skeleton/server"
)

func main() {
	server.Start(":8080")
	closer.Hold()
}
