package bookdb

import (
	"database/sql"
	"github.com/NewTanachot/learn-go/database"
)

var dbContext *sql.DB

func main() {
	tempDb, connectionError := db.Connect()

	if connectionError != nil {
		panic(connectionError)
	}

	dbContext = tempDb
}
