package gormbook

import (
	"fmt"
	"strconv"

	db "github.com/NewTanachot/learn-go/database"
	"github.com/NewTanachot/learn-go/dto"
	"github.com/NewTanachot/learn-go/model"
	"github.com/gofiber/fiber/v2"
)

func CreateAuthorGrom(context *fiber.Ctx) error {
	author := new(model.Author)
	error := context.BodyParser(author)

	if error != nil {
		return context.SendStatus(fiber.StatusBadRequest)
	}

	gormDb := db.GormSingletonConnection()
	gormDb.Create(author)

	return context.SendStatus(fiber.StatusCreated)
}

func CreateBookGorm(context *fiber.Ctx) error {
	gormBookRequest := new(dto.BookRequestDto)
	err := context.BodyParser(gormBookRequest)

	fmt.Println(gormBookRequest)

	if err != nil {
		return context.SendStatus(fiber.StatusBadRequest)
	}

	user := new(model.User)
	gormDb := db.GormSingletonConnection()
	queryResult := gormDb.First(&user, gormBookRequest.User.Id)

	if queryResult.Error != nil {
		return context.Status(fiber.StatusNotFound).JSON(queryResult.Error)
	}

	fmt.Println(*user)

	author := new(model.Author)
	queryResult = gormDb.First(&author, gormBookRequest.Author.Id)

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

	createErr := gormDb.Create(&book)

	if createErr.Error != nil {
		return context.Status(fiber.StatusBadRequest).JSON(createErr.Error)
	}

	return context.SendStatus(fiber.StatusCreated)
}

func GetBookFilterGorm(context *fiber.Ctx) error {

	filter := context.Params("filter")

	booksResponse := new([]model.GormBook)
	gormDb := db.GormSingletonConnection()
	queryResult := gormDb.
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

func GetBookByIdGorm(context *fiber.Ctx) error {

	id := context.Params("id")
	intId, err := strconv.Atoi(id)

	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(err)
	}

	bookResponse := new(model.GormBook)
	gormDb := db.GormSingletonConnection()

	queryResult := gormDb.
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

func UpdateBookGorm(context *fiber.Ctx) error {

	gormBookRequest := new(model.GormBook)
	err := context.BodyParser(gormBookRequest)

	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(err)
	}

	gormDb := db.GormSingletonConnection()

	// update all entity
	// result := gormDb.Save(gormBookRequest)

	// update specifix column in entity
	result := gormDb.Model(gormBookRequest).Updates(gormBookRequest)

	if result.Error != nil {
		return context.Status(fiber.StatusBadRequest).JSON(result.Error)
	}

	return context.SendStatus(fiber.StatusOK)
}

func DeleteBookGorm(context *fiber.Ctx) error {
	id := context.Params("id")
	isHardDelete := context.Params("hard")

	intId, err := strconv.Atoi(id)

	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(err)
	}

	bookResponse := new(model.GormBook)
	gormDb := db.GormSingletonConnection()

	if isHardDelete != "" {
		gormDb = gormDb.Unscoped().Delete(bookResponse, intId)
	} else {
		gormDb = gormDb.Delete(bookResponse, intId)
	}

	if gormDb.Error != nil {
		return context.Status(fiber.StatusBadRequest).JSON(gormDb.Error)
	}

	return context.JSON(bookResponse)
}
