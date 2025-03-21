package main

import (
	"fmt"
	"gemini-backend/api"
	"gemini-backend/db"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Loading environment variables...")
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: No .env file found, using system environment variables")
	}

	fmt.Println("Connecting to database...")
	if err := db.ConnectDatabase(); err != nil {
		log.Fatalf("Database connection error: %v", err)
	}

	fmt.Println("Database connected successfully!")
	router := gin.Default()
	api.RegisterRoutes(router)
	router.Run(":8080")
}
