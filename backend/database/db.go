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

	// Auto migrate User struct to create "users" table if not exist
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("❌ Failed to migrate database: %v", err)
	}
	//Auto migrate GmailAccount  struct to create gmails tables if not exist
	err = db.AutoMigrate(&models.GmailAccount{})
	if err != nil {
		log.Fatalf("❌ Failed to migrate database: %v", err)
	}
	DB = db
	fmt.Println("✅ Connected to the database with GORM")
}
