
package routes

import (
	"jobbotic-backend/handlers"
	"jobbotic-backend/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupAuthRoutes(router fiber.Router) {
	auth := router.Group("/auth")

	// Public routes
	auth.Post("/register", handlers.Register)
	auth.Post("/login", handlers.Login)

	// Google OAuth Signup/Login
	auth.Get("/google/login", handlers.GoogleLogin)       // Redirect to Google login
	auth.Get("/google/callback", handlers.GoogleCallback) // Callback with JWT in state

	// Gmail linking (requires frontend to generate state with JWT)
	auth.Get("/google/link", handlers.GoogleLink)                 // Returns Google OAuth URL
	auth.Get("/google/callback/link", middleware.RequireAuth, handlers.GoogleLinkCallback) // Handles linking
}
