package main

import (
	"fmt"
	"log"
	"os"
	"userOnboard/config"
	"userOnboard/router"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()

	r := gin.Default()

	router.SetupRoutes(r)

	// Get port from environment variable or use default
	port := os.Getenv("APP_PORT")
	if port == "" && len(os.Args) > 1 {
		port = os.Args[1]
	}
	if port == "" {
		port = "8080" // Default port
	}

	log.Printf("Server is running on port %s", port)
	if err := r.Run(fmt.Sprintf(":%s", port)); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
