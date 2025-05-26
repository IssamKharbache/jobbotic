package routes

import (
	"github.com/gofiber/fiber/v2"
	"jobbotic-backend/handlers"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	auth := api.Group("/auth")

	auth.Post("/register", handlers.Register)
}
