package main

import (
	"fmt"
	"log"

	"github.com/ApexPlayground/Linkkit/config"
	"github.com/ApexPlayground/Linkkit/routes"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	// Connect to DB
	db := config.Connect()
	myDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to get raw DB:", err)
	}
	defer myDB.Close()
	fmt.Println("Database connected")

	// Connect to Redis
	config.InitRedis()
	fmt.Println("Redis connected")

	// Setup Gin router
	router := routes.ShortnerSetupRouter(db)
	fmt.Println("Starting server on :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to run server:", err)
	}
}
