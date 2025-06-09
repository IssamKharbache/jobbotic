package routes

import (
	"jobbotic-backend/handlers"
	"jobbotic-backend/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(router fiber.Router) {
	//user group route
	users := router.Group("/users", middleware.RequireAuth)

	users.Get("/user/get", handlers.GetUserData)

}
