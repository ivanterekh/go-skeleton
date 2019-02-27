package server

import "github.com/gin-gonic/gin"

// Run starts a new server on the port.
func Run(port string) error {
	router := setupRouter()
	return router.Run(port)
}

func setupRouter() *gin.Engine {
	router := gin.New()
	router.GET("/", hello)
	return router
}
