package main

import (
	// "fmt"
	"github.com/NewTanachot/learn-go/book"
	// "github.com/NewTanachot/learn-go/book-db"
	"github.com/NewTanachot/learn-go/database"
	"github.com/NewTanachot/learn-go/middleware"
	"github.com/gofiber/fiber/v2"
	"net/url"
)

func main() {

	app := fiber.New()

	app.Use(middleware.TestMiddleware)

	db.Connect()

	app.Get("books", getBooks)
	app.Get("book/:name", getBookById)
	app.Post("book", insertBook)
	app.Put("book", updateBook)
	app.Delete("book/:name", deleteBook)

	app.Use(middleware.TestMiddleware)

	app.Listen(":3000")
}

// g is lower case. So this is private function
func getBookById(context *fiber.Ctx) error {
	name := context.Params("name")
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

func insertBook(context *fiber.Ctx) error {
	newBook := new(book.Book)
	error := context.BodyParser(newBook)

	if error != nil {
		return context.SendStatus(fiber.StatusBadRequest)
	}

	bookResponse := book.InsertBook(newBook)
	result := context.JSON(bookResponse)

	return result
}

func updateBook(context *fiber.Ctx) error {
	updateBook := new(book.Book)
	error := context.BodyParser(updateBook)

	if error != nil {
		return context.SendStatus(fiber.StatusBadRequest)
	}

	bookResponse := book.UpdateBook(updateBook)
	result := context.JSON(bookResponse)

	return result
}

func deleteBook(context *fiber.Ctx) error {
	name := context.Params("name")
	var decodedName, err = url.QueryUnescape(name)

	if err != nil {
		return context.SendStatus(fiber.StatusBadRequest)
	}

	idResponse := book.DeleteBook(&decodedName)
	result := context.JSON(idResponse)

	return result
}
