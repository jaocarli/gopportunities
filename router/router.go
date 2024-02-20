package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// Initialize the router using Gin's default settings
	router := gin.Default()

	// Defining a route
	initializeRoutes(router)

	// Run the server
	router.Run(":8080")
}
