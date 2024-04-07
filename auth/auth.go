package auth

import (
	"os"
	"time"

	db "github.com/NewTanachot/learn-go/database"
	"github.com/NewTanachot/learn-go/model"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Register(context *fiber.Ctx) error {
	user := new(model.User)
	err := context.BodyParser(user)

	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(err)
	}

	gormDb := db.GormSingletonConnection()
	err = createUser(gormDb, user)

	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(err)
	}

	return context.SendStatus(fiber.StatusCreated)
}

func Login(context *fiber.Ctx) error {
	user := new(model.User)
	err := context.BodyParser(user)

	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(err)
	}

	gormDb := db.GormSingletonConnection()
	jwt, err := loginUser(gormDb, user)

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

func createUser(db *gorm.DB, user *model.User) error {

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	user.Password = string(hash)

	db.Create(user)

	return nil
}

func loginUser(db *gorm.DB, user *model.User) (string, error) {

	selectUser := new(model.User)
	result := db.Where("email = ?", user.Email).First(selectUser)

	if result.Error != nil {
		return "", result.Error
	}

	err := bcrypt.CompareHashAndPassword([]byte(selectUser.Password), []byte(user.Password))

	if err != nil {
		return "", err
	}

	token, err := createJwtToken(selectUser.ID)

	if err != nil {
		return "", err
	}

	return token, nil
}

func createJwtToken(userId uint) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claim := token.Claims.(jwt.MapClaims)
	claim["user_id"] = userId
	claim["exp"] = time.Now().Add(168 * time.Hour).Unix()

	secretKey := os.Getenv("SECRETKEY")
	secretKeyBytes := []byte(secretKey)
	resultJwt, err := token.SignedString(secretKeyBytes)

	if err != nil {
		return "", err
	}

	return resultJwt, nil
}
