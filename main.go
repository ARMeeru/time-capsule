package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/ARMeeru/time-capsule/config"
	"github.com/ARMeeru/time-capsule/routes"
	"github.com/ARMeeru/time-capsule/utils"
)

func main() {
	// Load environment variables
	config.LoadEnv()

	// Connect to the database
	utils.ConnectDatabase()

	// Start the scheduler
	go utils.StartScheduler()

	// Set up the Gin router
	router := gin.Default()

	// Initialize routes
	routes.InitRoutes(router)

	// Start the server
	port := config.GetEnv("APP_PORT")
	if port == "" {
		port = ":8080"
		log.Printf("APP_PORT not specified. Using default port%s", port)
	} else {
		port = ":" + port
		log.Printf("Server is running on http://localhost%s", port)
	}
	err := router.Run(port)
	if err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}
