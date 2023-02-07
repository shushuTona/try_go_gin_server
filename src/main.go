package main

import "github.com/gin-gonic/gin"

func setupRouter() *gin.Engine {
	var router = gin.Default()

	router.GET("/api/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return router
}

func main() {
	var router = setupRouter()

	router.Run(":8080")
}