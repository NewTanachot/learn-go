package main

import (
	"net/url"
	"strconv"

	"github.com/NewTanachot/learn-go/book"
	db "github.com/NewTanachot/learn-go/database"
	"github.com/NewTanachot/learn-go/middleware"
	"github.com/NewTanachot/learn-go/model"
	"github.com/NewTanachot/learn-go/product"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var dbContext *gorm.DB

func main() {

	dbContext = db.GormConnect()
	db.MigrateLearnGorm(dbContext)

	app := fiber.New()

	app.Use(middleware.InterMiddleware)
	app.Use(middleware.OuterMiddleware)

	app.Get("gorm/book/:id", getBookByIdGorm)
	app.Get("gorm/book/filter/:filter", getBookFilterGorm)
	app.Post("gorm/book", createBookGorm)
	app.Put("gorm/book", updateBookGorm)
	app.Delete("gorm/book/:id/:hard?", deleteBookGorm)

	app.Get("book", getBooks)
	app.Get("book/:name", getBookById)
	app.Post("book", insertBook)
	app.Put("book", updateBook)
	app.Delete("book/:name", deleteBook)

	app.Get("product", getProducts)
	app.Get("product/:id", getProductById)
	app.Post("product", createProduct)
	app.Put("product", updateProduct)
	app.Delete("product/:id", deleteProduct)

	app.Listen(":3000")
}

func createBookGorm(context *fiber.Ctx) error {
	gormBookRequest := new(model.GormBook)
	error := context.BodyParser(gormBookRequest)

	if error != nil {
		return context.SendStatus(fiber.StatusBadRequest)
	}

	dbContext.Create(gormBookRequest)

	return context.SendStatus(fiber.StatusCreated)
}

func getBookFilterGorm(context *fiber.Ctx) error {

	filter := context.Params("filter")

	booksResponse := new([]model.GormBook)
	queryResult := dbContext.
		Where("author = ?", filter).
		Order("id desc").
		Find(&booksResponse)

	if queryResult.Error != nil {
		return context.Status(fiber.StatusBadRequest).JSON(queryResult.Error)
	}

	return context.JSON(booksResponse)
}

func getBookByIdGorm(context *fiber.Ctx) error {

	id := context.Params("id")
	intId, err := strconv.Atoi(id)

	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(err)
	}

	bookResponse := new(model.GormBook)
	queryResult := dbContext.First(&bookResponse, intId)

	if queryResult.Error != nil {
		return context.Status(fiber.StatusBadRequest).JSON(queryResult.Error)
	}

	return context.JSON(bookResponse)
}

func updateBookGorm(context *fiber.Ctx) error {

	gormBookRequest := new(model.GormBook)
	err := context.BodyParser(gormBookRequest)

	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(err)
	}

	result := dbContext.Save(gormBookRequest)

	if result.Error != nil {
		return context.Status(fiber.StatusBadRequest).JSON(result.Error)
	}

	return context.SendStatus(fiber.StatusOK)
}

func deleteBookGorm(context *fiber.Ctx) error {
	id := context.Params("id")
	isHardDelete := context.Params("hard")

	intId, err := strconv.Atoi(id)

	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(err)
	}

	bookResponse := new(model.GormBook)
	var response *gorm.DB

	if isHardDelete != "" {
		response = dbContext.Unscoped().Delete(bookResponse, intId)
	} else {
		response = dbContext.Delete(bookResponse, intId)
	}

	if response.Error != nil {
		return context.Status(fiber.StatusBadRequest).JSON(response.Error)
	}

	return context.JSON(bookResponse)
}

// region Product

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

// endregion

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
