package routes

import (
	"jobbotic-backend/handlers"
	"jobbotic-backend/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupAuthRoutes(router fiber.Router) {
	auth := router.Group("/auth")

	auth.Post("/register", handlers.Register)
	auth.Post("/login", handlers.Login)
	// Google OAuth
	auth.Get("/google/login", handlers.GoogleLogin)     // Redirects to Google
	auth.Get("/google/callback", handlers.GoogleSignup) // Handles callback and signup/login

	auth.Use(middleware.RequireAuth) // Protect routes below this line
	auth.Get("/google/link", handlers.GoogleLink)
	auth.Get("/google/callback/link", handlers.GoogleLinkCallback)
}
