package routes

import (
	"jobbotic-backend/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(router fiber.Router) {
	//user group route
	users := router.Group("/users")

	users.Get("/", handlers.GetUsers)

}
