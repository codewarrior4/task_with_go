package routes

import "github.com/gofiber/fiber/v2"

// SetupRoutes initializes all application routes
func SetupRoutes(app *fiber.App) {
	SetupAuthRoutes(app)
}
