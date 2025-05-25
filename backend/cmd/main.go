package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"jobbotic-backend/database"
)

func main() {
	database.Connect()
	defer database.Close()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Jobbotic backend is up and running!")
	})

	port := "3000"
	fmt.Printf("🚀 Server listening on port %s\n", port)
	app.Listen(":" + port)
}
