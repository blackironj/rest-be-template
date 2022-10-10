package controller

import (
	"github.com/gofiber/fiber/v2"

	"github.com/blackironj/rest-be-template/model"
	"github.com/blackironj/rest-be-template/server/common"
	"github.com/blackironj/rest-be-template/server/service"
)

// RegisterBook is a function to register a book
// @Summary register a book by isbn
// @Description register a book by isbn
// @Tags books
// @Accept json
// @Produce json
// @Param book body model.Book true "Register book"
// @Success 200 {object} common.ResponseHTTP{}
// @Failure 400 {object} common.ResponseHTTP{}
// @Failure 500 {object} common.ResponseHTTP{}
// @Router /books [post]
func RegisterBook(c *fiber.Ctx) error {
	var newBook model.Book

	if err := c.BodyParser(&newBook); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.ResponseHTTP{
			Success: false,
			Data:    nil,
			Message: "wrong param",
		})
	}

	if err := service.RegisterBook(&newBook); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(common.ResponseHTTP{
			Success: false,
			Data:    err.Error(),
			Message: "db err",
		})
	}
	return c.Status(fiber.StatusOK).JSON(common.ResponseHTTP{
		Success: true,
	})
}

// GetBookByISBN is a function to get a book by isbn
// @Summary Get a book by isbn
// @Description Get a book by isbn
// @Tags books
// @Accept json
// @Produce json
// @Param isbn path string true "Book ISBN code"
// @Success 200 {object} common.ResponseHTTP{data=model.Book}
// @Failure 400 {object} common.ResponseHTTP{}
// @Failure 404 {object} common.ResponseHTTP{}
// @Router /books/{isbn} [get]
func GetBookByISBN(c *fiber.Ctx) error {
	isbn := c.Params("isbn")
	if isbn == "" {
		return c.Status(fiber.StatusBadRequest).JSON(common.ResponseHTTP{
			Success: false,
			Data:    nil,
			Message: "wrong isbn",
		})
	}

	book := service.GetBookByISBN(isbn)
	if book == nil {
		return c.Status(fiber.StatusNotFound).JSON(common.ResponseHTTP{
			Success: false,
			Data:    nil,
			Message: "book not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(common.ResponseHTTP{
		Success: true,
		Data:    book,
	})
}
