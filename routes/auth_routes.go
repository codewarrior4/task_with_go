package routes

import (
	"github.com/gofiber/fiber/v2"
	"task/handlers"
	"task/middlewares"
)

// SetupAuthRoutes registers authentication-related routes
func SetupAuthRoutes(app *fiber.App) {
	api := app.Group("/api/v1") // API Prefix

	auth := api.Group("/auth")

	// Apply different rate limits to different routes
	auth.Post("/register", middlewares.RateLimiter(1), handlers.Register)         // 10 requests per 1 min
	auth.Post("/login", middlewares.RateLimiter(2), handlers.Login)               // 10 requests per 2 min
	auth.Post("/forgot-password", middlewares.RateLimiter(5), handlers.ForgotPassword) // 10 requests per 5 min
	auth.Post("/reset-password", middlewares.RateLimiter(5), handlers.ResetPassword)
	auth.Post("/change-password", middlewares.RateLimiter(3), handlers.ChangePassword)
}
