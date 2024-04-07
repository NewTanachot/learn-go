package cvalidator

import (
	"fmt"
	"regexp"
	"sync"

	"github.com/go-playground/validator/v10"
)

var (
	cValidator *validator.Validate
	once       sync.Once
)

func SingletonSetUp() *validator.Validate {

	// use once keyword to do singleton implementation of validator connection pool
	once.Do(func() {
		cValidator = validator.New()
		customValidatorRegistry(cValidator)

		fmt.Println("Create Validator Success")
	})

	return cValidator
}

func customValidatorRegistry(v *validator.Validate) {
	v.RegisterValidation("fullname", validateFullname)
}

func validateFullname(fl validator.FieldLevel) bool {
	return regexp.MustCompile(`^[a-zA-Z\s]+$`).MatchString(fl.Field().String())
}
