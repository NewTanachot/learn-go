package main

import (
	// "fmt"
	"github.com/NewTanachot/learn-go/book"
	"github.com/gofiber/fiber/v2"
	"net/url"
)

func main() {

	app := fiber.New()

	app.Get("books", getBooks)
	app.Get("book/:id", getBookById)

	app.Listen(":3000")
}

func getBookById(context *fiber.Ctx) error {
	name := context.Params("id")
	var decodedName, err = url.QueryUnescape(name)

	if err != nil {
		return err
	}

	book := book.GetBookById(&decodedName)
	var result = context.JSON(book)

	return result
}

func getBooks(context *fiber.Ctx) error {
	books := book.GetBooks()
	var result = context.JSON(books)

	return result
}
