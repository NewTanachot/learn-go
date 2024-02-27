package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func Connect() (*sql.DB, error) {
	dbData, err := getDbMetaData()

	if err != nil {
		return nil, err
	}

	// Connection string
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbData.Host, dbData.Port, dbData.UserName, dbData.Password, dbData.DbName)

	dbContext, connectionError := sql.Open("postgres", connectionString)

	if connectionError != nil {
		return nil, connectionError
	}

	println("Connect: success")

	pingError := dbContext.Ping()

	if pingError != nil {
		return nil, pingError
	}
	println("Ping: success")

	return dbContext, nil
}

func Close(db *sql.DB) {
	db.Close()
	println("db is closed")
}
