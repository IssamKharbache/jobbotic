package main

import (
	"fmt"
	"log"
	"os"

	"jobbotic-backend/database"
	"jobbotic-backend/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173", // React dev server
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: true,
	}))
	api := app.Group("/api")
	//route setup
	routes.SetupAuthRoutes(api)
	routes.SetupUserRoutes(api)
	routes.SetupEmailRoutes(api)
	//running app
	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}
	//running the app
	fmt.Printf("🚀 Server listening on port %s\n", port)
	log.Fatal(app.Listen(":" + port))
}
