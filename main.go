package main

import (
	"Upload-Dwonload/handlers"
	"log"
	"net/http"
)

func main() {
	// Initialize your database connection here
	// Example: db.Init()

	// Set up middleware if needed
	// Example: router.Use(myMiddleware)

	// Initialize your handlers
	handlers.Init()

	// Run the server
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
