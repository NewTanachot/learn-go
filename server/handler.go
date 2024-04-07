package server

import (
	"net/url"
	"strconv"

	book "github.com/NewTanachot/learn-go/Book"
	"github.com/NewTanachot/learn-go/model"
	"github.com/NewTanachot/learn-go/product"
	cvalidator "github.com/NewTanachot/learn-go/validator"
	"github.com/gofiber/fiber/v2"
)

// -=-=-=-=-=-=-=-=-=-=-=- [ Product ] -=-=-=-=-=-=-=-=-=-=-=-

func createProduct(context *fiber.Ctx) error {

	newProduct := new(product.Product)
	error := context.BodyParser(newProduct)

	if error != nil {
		return context.SendStatus(fiber.StatusBadRequest)
	}

	err := product.CreateProduct(newProduct)

	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(err)
	}

	return context.SendStatus(fiber.StatusCreated)
}

func getProducts(context *fiber.Ctx) error {
	result, err := product.GetProducts()

	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(err)
	}

	return context.JSON(result)
}

func getProductById(context *fiber.Ctx) error {
	id := context.Params("id")
	intId, err := strconv.Atoi(id)

	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(err)
	}

	result, err := product.GetProduct(intId)

	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(err)
	}

	return context.JSON(result)
}

func updateProduct(context *fiber.Ctx) error {
	updateProduct := new(product.Product)
	err := context.BodyParser(updateProduct)

	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(err)
	}

	result, err := product.UpdateProduct(updateProduct.Id,
		updateProduct.Name, updateProduct.Price)

	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(err)
	}

	return context.JSON(result)
}

func deleteProduct(context *fiber.Ctx) error {
	id := context.Params("id")
	intId, err := strconv.Atoi(id)

	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(err)
	}

	err = product.DeleteProduct(intId)

	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(err)
	}

	return context.SendStatus(fiber.StatusOK)
}

// -=-=-=-=-=-=-=-=-=-=-=- [ BOOK ] -=-=-=-=-=-=-=-=-=-=-=-

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
	err := context.BodyParser(newBook)

	if err != nil {
		return context.SendStatus(fiber.StatusBadRequest)
	}

	customValidator := cvalidator.SingletonSetUp()
	err = customValidator.Struct(newBook)

	if err != nil {
		return context.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"error": err.Error()})
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

// -=-=-=-=-=-=-=-=-=-=-=- [ Test Validator ] -=-=-=-=-=-=-=-=-=-=-=-

func studentValidator(context *fiber.Ctx) error {
	student := new(model.Student)

	if err := context.BodyParser(student); err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(err)
	}

	customValidator := cvalidator.SingletonSetUp()
	if err := customValidator.Struct(student); err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(err)
	}

	return context.JSON(student)
}
