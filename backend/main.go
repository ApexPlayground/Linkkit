package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ApexPlayground/Linkkit/config"
	"github.com/ApexPlayground/Linkkit/controller"
	"github.com/ApexPlayground/Linkkit/routes"
	"github.com/ApexPlayground/Linkkit/service"
	"github.com/gin-gonic/gin"
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

	geoipPath := os.Getenv("GEOIP_DB_PATH")
	clickSvc := service.NewClickService(db, geoipPath)
	redirectSvc := service.NewRedirectService(db, clickSvc)

	controller.InitRedirectController(redirectSvc)

	router := gin.Default()

	// Mount all route groups
	routes.UserSetupRouter(router)
	routes.ShortenerSetupRouter(router)
	// routes.ClickSetupRouter(router)

	fmt.Println("Starting server on :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to run server:", err)
	}
}
