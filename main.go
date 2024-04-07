package main

import (
	"github.com/NewTanachot/learn-go/server"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		panic("load .env error")
	}

	app := server.Setup()
	app.Listen(":3000")
}
