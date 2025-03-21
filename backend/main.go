package main

import (
	"fmt"
	"gemini-backend/api"
	"gemini-backend/db"

	"github.com/gin-gonic/gin"
)

func main() {
	// Database Connection
	if err := db.ConnectDatabase(); err != nil {
		panic(fmt.Sprintf("Database connection error: %v", err))
	}

	gin.SetMode(gin.ReleaseMode)

	// Setup Router
	router := gin.Default()

	router.SetTrustedProxies([]string{"127.0.0.1", "13.51.206.239"})

	// Register Routes
	api.RegisterRoutes(router)

	// Start the server
	router.Run(":8080")
}
