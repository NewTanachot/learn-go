package db

import (
	"os"
	"strconv"
)

type DbMetaData struct {
	Host     string
	Port     int16
	DbName   string
	UserName string
	Password string
}

func getDbMetaData() (*DbMetaData, error) {
	port, err := strconv.Atoi(os.Getenv("PORT"))

	if err != nil {
		return nil, err
	}

	result := DbMetaData{
		Host:     os.Getenv("HOST"),
		Port:     int16(port),
		DbName:   os.Getenv("DBNAME"),
		UserName: os.Getenv("USERNAME"),
		Password: os.Getenv("PASSWORD"),
	}

	return &result, nil
}

// const (
// 	host     = "localhost"
// 	port     = 5432
// 	dbname   = "mydatabase"
// 	username = "myuser"
// 	password = "mypassword"
// )
