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

	models := []any{
		&model.Link{},
		&model.Click{},
		&model.User{},
		&model.DailyLinkStats{},
		&model.LinkCountryStats{},
		&model.LinkDeviceStats{},
		&model.LinkBrowserStats{},
		&model.LinkReferrerStats{},
		&model.QRCode{},
		&model.QRBrowserStats{},
		&model.QRCountryStats{},
		&model.QRDeviceStats{},
		&model.QRReferrerStats{},
		&model.DailyQRStats{},
	}

	db.AutoMigrate(models...)

	fmt.Println("Database connection established")
	DB = db
	return DB
}
