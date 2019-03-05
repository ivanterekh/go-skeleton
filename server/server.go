package server

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xlab/closer"
)

// Start runs a server in a new goroutine.
func Start(listenAddr string) {
	server := http.Server{
		Addr:    listenAddr,
		Handler: setupRouter(),
	}

	closer.Bind(func() {
		log.Println("Stopping the server...") // TODO: change logging
		if err := server.Shutdown(context.Background()); err != nil {
			log.Printf("could not gracefully stop server: %v\n", err) // TODO: change logging
		}
	})

	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Println(err) // TODO: change logging
		}
	}()
}

func setupRouter() *gin.Engine {
	router := gin.New()
	router.GET("/", hello)
	router.GET("/env", env)
	return router
}
