package routes

import (
	"github.com/gofiber/fiber/v2"
	"jobbotic-backend/handlers"
)

func SetupAuthRoutes(router fiber.Router) {
	//auth group route
	auth := router.Group("/auth")

	auth.Post("/register", handlers.Register)
	auth.Post("/login", handlers.Login)
}
