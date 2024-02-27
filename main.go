package main

import (
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/NewTanachot/learn-go/auth"
	"github.com/NewTanachot/learn-go/book"
	db "github.com/NewTanachot/learn-go/database"
	"github.com/NewTanachot/learn-go/dto"
	"github.com/NewTanachot/learn-go/goroutine"
	"github.com/NewTanachot/learn-go/middleware"
	"github.com/NewTanachot/learn-go/model"
	"github.com/NewTanachot/learn-go/product"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var dbContext *gorm.DB

func main() {

	if err := godotenv.Load(); err != nil {
		panic("load .env error")
	}

	dbContext = db.GormConnect()
	db.Migrate(dbContext)

	app := fiber.New()

	app.Get("goroutine/channel", goroutine.TestChannel)
	app.Get("goroutine/waitgroup", goroutine.TestWaitGroup)
	app.Get("goroutine/mutex", goroutine.TestMuTexLock)
	app.Get("goroutine/cond", goroutine.TestCond)
	app.Get("goroutine/pubsub/:message", goroutine.TestPubSub)

	// slice := make([]int, 5)

	// fmt.Println(slice)
	// fmt.Println(len(slice))
	// fmt.Println(cap(slice))

	app.Use(middleware.InterMiddleware)
	app.Use("gorm/book", middleware.AuthRequiredMiddleware)
	app.Use(middleware.OuterMiddleware)

	app.Post("register", register)
	app.Post("login", login)

	app.Post("gorm/author", createAuthorGrom)

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

func register(context *fiber.Ctx) error {
	user := new(model.User)
	err := context.BodyParser(user)

	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(err)
	}

	err = auth.CreateUser(dbContext, user)

	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(err)
	}

	return context.SendStatus(fiber.StatusCreated)
}

func login(context *fiber.Ctx) error {
	user := new(model.User)
	err := context.BodyParser(user)

	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(err)
	}

	jwt, err := auth.LoginUser(dbContext, user)

	if err != nil {
		return context.Status(fiber.StatusUnauthorized).JSON(err)
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    jwt,
		Expires:  time.Now().Add(time.Hour * 72),
		HTTPOnly: true,
	}

	context.Cookie(&cookie)

	response := fiber.Map{
		"status": "success",
		"jwt":    jwt,
	}

	// response := map[string]interface{}{
	// 	"status": "success",
	// 	"jwt":    jwt,
	// }

	return context.JSON(response)
}

// ---

func createAuthorGrom(context *fiber.Ctx) error {
	author := new(model.Author)
	error := context.BodyParser(author)

	if error != nil {
		return context.SendStatus(fiber.StatusBadRequest)
	}

	dbContext.Create(author)

	return context.SendStatus(fiber.StatusCreated)
}

func createBookGorm(context *fiber.Ctx) error {
	gormBookRequest := new(dto.BookRequestDto)
	err := context.BodyParser(gormBookRequest)

	fmt.Println(gormBookRequest)

	if err != nil {
		return context.SendStatus(fiber.StatusBadRequest)
	}

	user := new(model.User)
	queryResult := dbContext.First(&user, gormBookRequest.User.Id)

	if queryResult.Error != nil {
		return context.Status(fiber.StatusNotFound).JSON(queryResult.Error)
	}

	fmt.Println(*user)

	author := new(model.Author)
	queryResult = dbContext.First(&author, gormBookRequest.Author.Id)

	if queryResult.Error != nil {
		return context.Status(fiber.StatusNotFound).JSON(queryResult.Error)
	}

	fmt.Println(*author)

	book := model.GormBook{
		Name:        gormBookRequest.Book.Name,
		Description: gormBookRequest.Book.Description,
		Price:       gormBookRequest.Book.Price,
		User:        *user,
		Author:      []model.Author{*author},
	}

	createErr := dbContext.Create(&book)

	if createErr.Error != nil {
		return context.Status(fiber.StatusBadRequest).JSON(createErr.Error)
	}

	return context.SendStatus(fiber.StatusCreated)
}

func getBookFilterGorm(context *fiber.Ctx) error {

	filter := context.Params("filter")

	booksResponse := new([]model.GormBook)
	queryResult := dbContext.
		Preload("Author").
		Preload("User").
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
	queryResult := dbContext.
		Preload("Author").
		Preload("User").
		First(&bookResponse, intId)

	if queryResult.Error != nil {
		return context.Status(fiber.StatusBadRequest).JSON(queryResult.Error)
	}

	// context.Response().Header.Del("Content-Type")
	// return context.Send([]byte(bookResponse.Name))

	return context.JSON(bookResponse)
}

func updateBookGorm(context *fiber.Ctx) error {

	gormBookRequest := new(model.GormBook)
	err := context.BodyParser(gormBookRequest)

	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(err)
	}

	// update all entity
	// result := dbContext.Save(gormBookRequest)

	// update specifix column in entity
	result := dbContext.Model(gormBookRequest).Updates(gormBookRequest)

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
