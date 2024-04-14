package main

import (
	"Upload-Dwonload/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Set up routes and handlers
	handlers.AdminInit(router)
	handlers.UserInit(router)

	// Run the server
	router.Run(":8080")
}
