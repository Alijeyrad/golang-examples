package routes

import (
	"yourapp/controllers"
	"yourapp/middleware"
	"yourapp/repositories"
	"yourapp/services"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	bookRepo := repositories.NewBookRepository()
	bookService := services.NewBookService(bookRepo)
	bookController := controllers.NewBookController(bookService)

	app.Use(middleware.RateLimiter())

	app.Get("/books", bookController.GetAllBooks)
	app.Get("/books/:id", bookController.GetBookByID)
	app.Post("/books", bookController.CreateBook)
	app.Put("/books/:id", bookController.UpdateBook)
	app.Delete("/books/:id", bookController.DeleteBook)
}
