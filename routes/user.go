package routes

import (
	"github.com/gofiber/fiber/v2"
	"task/handlers"
	"task/middlewares"
)

func SetupUserRoutes(app *fiber.App) {
	api := app.Group("/api/v1", middlewares.AuthMiddleware()) // API Prefix

	// User routes
	api.Get("/dashboard", handlers.Dashboard)

	// task routes
	tasks := api.Group("/tasks", middlewares.AuthMiddleware())

	tasks.Get("", handlers.GetTasks)                   // Get all tasks (with optional pagination/filtering)
	tasks.Get("/:id", handlers.GetTaskByID)                 // Get one task
	// tasks.Post("/", handlers.CreateTask)               // Create one task
	// tasks.Post("/bulk", handlers.CreateTasksBulk)      // Create multiple tasks
	// tasks.Put("/:id", handlers.UpdateTask)             // Update a task
	// tasks.Delete("/:id", handlers.DeleteTask)          // Soft delete a task
	// tasks.Delete("/bulk", handlers.DeleteTasksBulk)    // Soft delete multiple tasks
	// tasks.Get("/filter", handlers.FilterTasks)         // Filter by status/due date
	// tasks.Delete("/force/:id", handlers.ForceDeleteTask) // Permanently delete a task

}	