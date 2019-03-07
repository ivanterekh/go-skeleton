package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
)

// Run listens given address.
func Run(listenAddr string) {
	server := http.Server{
		Addr:    listenAddr,
		Handler: setupRouter(),
	}

	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatal(err) //TODO: add logging
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...") //TODO: add logging

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown: ", err) //TODO: add logging
	}
}

func setupRouter() *gin.Engine {
	router := gin.New()
	router.GET("/", hello)
	return router
}
