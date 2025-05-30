package database

import (
	"fmt"
	"jobbotic-backend/models"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ Failed to connect to database: %v", err)
	}
	err = db.AutoMigrate(&models.User{}, &models.GmailAccount{}, &models.JobApplication{})
	if err != nil {
		log.Fatalf("❌ Failed to migrate database: %v", err)
	}
	DB = db
	fmt.Println("✅ Connected to the database with GORM")
}
