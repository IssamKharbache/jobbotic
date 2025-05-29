package routes

import (
	"jobbotic-backend/handlers"
	"jobbotic-backend/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupEmailRoutes(router fiber.Router) {
	auth := router.Group("/email")

	auth.Get("/get-email/:id", middleware.RequireAuth, handlers.FetchEmails)
}
