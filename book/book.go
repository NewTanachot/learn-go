package book

import (
	"fmt"
	"github.com/google/uuid"
)

type Book struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Title string `json:"title"`
}

// global var scope has (STATE)
var sliceOfBook []Book = []Book{
	{
		Id:    uuid.NewString(),
		Name:  "Dota 1",
		Title: "Dota 1",
	},
	{
		Id:    uuid.NewString(),
		Name:  "Dota 2",
		Title: "Dota 2",
	},
}

func GetBookById(name *string) Book {

	fmt.Println(*name)

	for _, book := range sliceOfBook {
		if book.Name == *name {
			return book
		}
	}

	return Book{}
}

func GetBooks() []Book {

	//  local var scope has (no STATE)

	if len(sliceOfBook) == 0 {

		sliceOfBook = append(sliceOfBook, Book{
			Id:    uuid.NewString(),
			Name:  "Dota 1",
			Title: "Dota 1",
		})

		sliceOfBook = append(sliceOfBook, Book{
			Id:    uuid.NewString(),
			Name:  "Dota 2",
			Title: "Dota 2",
		})
	}

	return sliceOfBook
}
