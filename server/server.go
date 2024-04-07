package server

import (
	"github.com/NewTanachot/learn-go/auth"
	gormbook "github.com/NewTanachot/learn-go/gorm_book"
	"github.com/NewTanachot/learn-go/goroutine"
	"github.com/NewTanachot/learn-go/middleware"
	cvalidator "github.com/NewTanachot/learn-go/validator"
	"github.com/gofiber/fiber/v2"
)

func Setup() *fiber.App {

	app := fiber.New()
	cvalidator.SingletonSetUp()

	app.Get("goroutine/channel", goroutine.TestChannel)
	app.Get("goroutine/waitgroup", goroutine.TestWaitGroup)
	app.Get("goroutine/mutex", goroutine.TestMuTexLock)
	app.Get("goroutine/cond", goroutine.TestCond)
	app.Get("goroutine/pubsub/:message", goroutine.TestPubSub)

	app.Use(middleware.InterMiddleware)
	app.Use("gorm/book", middleware.AuthRequiredMiddleware)
	app.Use(middleware.OuterMiddleware)

	app.Post("register", auth.Register)
	app.Post("login", auth.Login)

	app.Post("gorm/author", gormbook.CreateAuthorGrom)

	app.Get("gorm/book/:id", gormbook.GetBookByIdGorm)
	app.Get("gorm/book/filter/:filter", gormbook.GetBookFilterGorm)
	app.Post("gorm/book", gormbook.CreateBookGorm)
	app.Put("gorm/book", gormbook.UpdateBookGorm)
	app.Delete("gorm/book/:id/:hard?", gormbook.DeleteBookGorm)

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

	app.Post("student", studentValidator)

	return app
}
