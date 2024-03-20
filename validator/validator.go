package custom_validator

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func CustomValidatorRegistry(v *validator.Validate) {
	v.RegisterValidation("fullname", validateFullname)
}

func validateFullname(fl validator.FieldLevel) bool {
	return regexp.MustCompile(`^[a-zA-Z\s]+$`).MatchString(fl.Field().String())
}
