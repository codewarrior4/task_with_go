package main

import (
	"log"

	"task/config"
	"task/database"
	"task/routes"
	"task/utils"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Initialize the database

	config.LoadEnv()
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
	if err := app.Listen(":4023"); err != nil {
		log.Fatal(err)
	}
}
