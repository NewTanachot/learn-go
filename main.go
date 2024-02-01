package main

import (
	"fmt"
	"github.com/google/uuid"
)

func main() {

	a := uuid.New()

	fmt.Println("hello world")
	fmt.Println(a)
}
