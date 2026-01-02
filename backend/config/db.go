package config

import (
	"fmt"
	"log"
	"os"

	"github.com/ApexPlayground/Linkkit/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() *gorm.DB {
	// Connection string
	dbURL := os.Getenv("DB_CONN_STR")

	// Open the connection
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	models := []interface{}{
		&model.Link{},
		&model.Click{},
	}

	db.AutoMigrate(models...)

	fmt.Println("Database connection established")
	DB = db
	return DB
}
