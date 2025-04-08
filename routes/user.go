package routes

import (
	"github.com/gofiber/fiber/v2"
	"task/handlers"
	"task/middlewares"
)

func SetupUserRoutes(app *fiber.App) {
	api := app.Group("/api/v1", middlewares.AuthMiddleware()) // API Prefix

	api.Get("/dashboard", handlers.Dashboard)
}