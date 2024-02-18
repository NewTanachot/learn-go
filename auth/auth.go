package auth

import (
	"os"
	"time"

	"github.com/NewTanachot/learn-go/model"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB, user *model.User) error {

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	user.Password = string(hash)

	db.Create(user)

	return nil
}

func LoginUser(db *gorm.DB, user *model.User) (string, error) {

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
	claim["exp"] = time.Now().Add(72 * time.Hour).Unix()

	secretKey := os.Getenv("SECRETKEY")
	secretKeyBytes := []byte(secretKey)
	resultJwt, err := token.SignedString(secretKeyBytes)

	if err != nil {
		return "", err
	}

	return resultJwt, nil
}
