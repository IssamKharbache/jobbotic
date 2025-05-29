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
	auth.Get("/google/login", handlers.GoogleLogin)       // Redirect to Google login (for signup/login)
	auth.Get("/google/callback", handlers.GoogleCallback) // Callback after signup/login
	auth.Get("/test-token/:id", handlers.TestAccessToken)
	// Gmail Linking (Requires logged-in user)
	auth.Get("/google/link", middleware.RequireAuth, handlers.GoogleLink) // Returns OAuth URL for linking
	auth.Get("/google/callback/link", handlers.GoogleLinkCallback)        // Handles OAuth2 code and saves tokens
}
