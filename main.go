package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"task/config"
	"task/database"
	// "task/routes"
)

func main() {
	// Initialize the database
	config.ConnectDB()

	// Migrate the database
	database.Migrate()

	// Run my seeders
	database.Seed()

	// Create a new Fiber instance
	app := fiber.New()

	// Setup routes
	// routes.SetupRoutes(app)

	// Start the server
	if err := app.Listen(":3000"); err != nil {
		log.Fatal(err)
	}
}
