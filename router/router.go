package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// Initialize the router using Gin's default settings
	router := gin.Default()
	// Defining a route
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run(":8080")
}
