package controllers

import (
	"strconv"

	"yourapp/models"
	"yourapp/services"

	"github.com/gofiber/fiber/v2"
)

type BookController struct {
	BookService services.BookService
}

func NewBookController(bookService services.BookService) *BookController {
	return &BookController{
		BookService: bookService,
	}
}

func (c *BookController) GetAllBooks(ctx *fiber.Ctx) error {
	books, err := c.BookService.GetAllBooks()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return ctx.JSON(books)
}

func (c *BookController) GetBookByID(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString("Invalid ID")
	}
	book, err := c.BookService.GetBookByID(uint(id))
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).SendString("Book not found")
	}
	return ctx.JSON(book)
}

func (c *BookController) CreateBook(ctx *fiber.Ctx) error {
	var book models.Book
	if err := ctx.BodyParser(&book); err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	err := c.BookService.CreateBook(&book)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return ctx.Status(fiber.StatusCreated).JSON(book)
}

func (c *BookController) UpdateBook(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString("Invalid ID")
	}
	var book models.Book
	if err := ctx.BodyParser(&book); err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	book.ID = uint(id)
	err = c.BookService.UpdateBook(&book)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return ctx.JSON(book)
}

func (c *BookController) DeleteBook(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString("Invalid ID")
	}
	err = c.BookService.DeleteBook(uint(id))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return ctx.SendStatus(fiber.StatusNoContent)
}
