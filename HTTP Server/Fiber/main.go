package main

import (
	"log"

	"yourapp/config"
	"yourapp/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Initialize the database
	config.InitDB()

	// Initialize Fiber app
	app := fiber.New()

	// Setup routes
	routes.SetupRoutes(app)

	// Start the server
	log.Println("Starting server on :8080")
	log.Fatal(app.Listen(":8080"))
}
