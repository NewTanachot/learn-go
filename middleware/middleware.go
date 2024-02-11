package middleware

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func TestMiddleware(context *fiber.Ctx) error {

	fmt.Println("this is middleware")

	return context.Next()
}
