package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/ApexPlayground/Linkkit/config"
	"github.com/ApexPlayground/Linkkit/controller"
	"github.com/ApexPlayground/Linkkit/routes"
	"github.com/ApexPlayground/Linkkit/service"
	"github.com/gin-contrib/cors"
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

	// Initialize Click Service (shared by both links and QR codes)
	geoipPath := os.Getenv("GEOIP_DB_PATH")
	clickSvc := service.NewClickService(db, geoipPath, 5)
	fmt.Println("Click Service initialized with 5 workers")

	// Initialize QR Service (NEW!)
	qrSvc := service.NewQRService(db, clickSvc)
	controller.InitQRController(qrSvc)
	fmt.Println("QR Service initialized")

	// Initialize Redirect Service (existing)
	redirectSvc := service.NewRedirectService(db, clickSvc)
	controller.InitLinkController(redirectSvc)
	fmt.Println("Redirect Service initialized")

	// Setup Gin router
	router := gin.Default()

	// CORS configuration
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	// Mount all route groups
	routes.UserSetupRouter(router)
	routes.AppSetupRouter(router)

	// Graceful shutdown setup
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Start server in goroutine
	go func() {
		fmt.Println("Starting server on :8080")
		if err := router.Run(":8080"); err != nil {
			log.Fatal("Failed to run server:", err)
		}
	}()

	// Wait for shutdown signal
	sig := <-quit
	fmt.Println("Received signal:", sig)

	// Graceful shutdown for ClickService
	clickSvc.Close()
	fmt.Println("ClickService stopped")
	fmt.Println("Server exiting")
}
