package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func InterMiddleware(context *fiber.Ctx) error {

	fmt.Println("this is inter middleware")

	return context.Next()
}

func OuterMiddleware(context *fiber.Ctx) error {

	next := context.Next()

	fmt.Println("this is outer middleware")

	return next
}
