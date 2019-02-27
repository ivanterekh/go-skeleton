package main

import (
	"github.com/ivanterekh/go-skeleton/server"
	"github.com/xlab/closer"
)

func main() {
	server.Start(":8080")
	closer.Hold()
}
