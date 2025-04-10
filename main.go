package main

import (
	// "log"
	"net/http"
	"task/config"
	"task/database"
	"task/routes"
	"task/utils"
	"task/jobs"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Load env and connect to DB
	config.LoadEnv()
	config.ConnectDB()

	// Start the scheduler in the background
	jobs.StartVerificationCleanup()

	// Migrate and seed
	database.Migrate()
	database.Seed()

	// Setup Fiber app
	appConfig := fiber.Config{
		AppName:           "Task Manager",
		CaseSensitive:     true,
		StrictRouting:     true,
		EnablePrintRoutes: true,
	}
	app := fiber.New(appConfig)

	// Middleware
	app.Use(utils.Logger())

	// Routes
	routes.SetupRoutes(app)

	// Start the server in a goroutine so it doesn't block the main thread
	go func() {
		// if err := app.Listen(":4023"); err != nil {
		// 	log.Fatal(err)
		// }
		port := config.GetEnv("DB_PORT", "")
		 
		http.ListenAndServe("0.0.0.0:"+port, nil)

	}()

	// 👇 Block forever so your background job keeps running
	select {}
}
