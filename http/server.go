package http

import "github.com/gin-gonic/gin"

func RunServer(port string) error {
	router := gin.New()

	router.GET("/", hello)

	return router.Run(port)
}
