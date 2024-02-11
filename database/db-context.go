package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	dbname   = "mydatabase"
	username = "myuser"
	password = "mypassword"
)

func Connect() (*sql.DB, error) {
	// Connection string
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, username, password, dbname)

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
