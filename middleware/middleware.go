package middleware

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
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

func AuthRequiredMiddleware(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	secretKey := os.Getenv("SECRETKEY")

	token, err := jwt.ParseWithClaims(cookie, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil || !token.Valid {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	println("Authorization Pass!")
	return c.Next()
}
