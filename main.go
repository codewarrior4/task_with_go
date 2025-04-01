package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"task/config"
	"task/database"
	"task/utils"
	"task/routes"
)

func main() {
	// Initialize the database

	
	config.ConnectDB()

	// Migrate the database
	database.Migrate()

	// Run my seeders
	database.Seed()

	// Create a new Fiber instance
	appConfig := fiber.Config{
		AppName:           "Task Manager",
		CaseSensitive:     true,
		StrictRouting:     true,
		EnablePrintRoutes: true,
	}
	app := fiber.New(appConfig)

	// Middleware
	app.Use(utils.Logger())

	// Setup routes
	routes.SetupRoutes(app)

	// Start the server
	if err := app.Listen(":3000"); err != nil {
		log.Fatal(err)
	}
}
