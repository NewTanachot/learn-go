package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/NewTanachot/learn-go/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func GormConnect() *gorm.DB {
	dbData, err := getDbMetaData()

	if err != nil {
		panic(err)
	}

	// Connection string
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbData.Host, dbData.Port, dbData.UserName, dbData.Password, dbData.DbName)

	gormLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Enable color
		},
	)

	gormOption := gorm.Config{
		Logger: gormLogger,
	}

	db, err := gorm.Open(postgres.Open(connectionString), &gormOption)

	if err != nil {
		panic("Fail to connect GORM DbContext")
	}

	println("Connect to GORM Success")

	return db
}

func Migrate(db *gorm.DB) {

	err := db.AutoMigrate(&model.GormBook{}, &model.User{})

	if err != nil {
		panic(err)
	}

	println("Migrate Success!")
}
