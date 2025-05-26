package main

import (
	"fmt"
	"log"
	"os"

	"jobbotic-backend/database"
	"jobbotic-backend/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {

	//loading the env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("❌ Could not load .env file: %v", err)
	}

	//connecting to the database
	database.ConnectDatabase()
	//initliasing a new fiber app
	app := fiber.New()
	//route setup
	routes.SetupRoutes(app)
	//running app
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	//running the app
	fmt.Printf("🚀 Server listening on port %s\n", port)
	log.Fatal(app.Listen(":" + port))
}
