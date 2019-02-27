package main

import "github.com/gin-gonic/gin"

func runServer(port string) error {
	router := setupRouter()
	return router.Run(port)
}

func setupRouter() *gin.Engine {
	router := gin.New()
	router.GET("/", hello)
	return router
}
