package student

import (
	"github.com/NewTanachot/learn-go/model"
	cvalidator "github.com/NewTanachot/learn-go/validator"
	"github.com/gofiber/fiber/v2"
)

func StudentValidator(context *fiber.Ctx) error {
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
