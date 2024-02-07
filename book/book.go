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

// G is upper case. So this is public function
func GetBookById(name *string) Book {

	fmt.Println(*name)

	// _ is index (if you want to use index change _ to i or something)
	for _, book := range sliceOfBook {
		if book.Name == *name {
			return book
		}
	}

	// new(Struct) is how to allocate memory address for this.Struct (use *new(Struct) to get value)
	return *new(Book)
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

func InsertBook(book *Book) []Book {
	sliceOfBook = append(sliceOfBook, *book)
	return sliceOfBook
}

func UpdateBook(updateBook *Book) *string {
	for index, book := range sliceOfBook {
		if book.Id == updateBook.Id {
			sliceOfBook[index].Name = updateBook.Name
			sliceOfBook[index].Title = updateBook.Title

			return &book.Id
		}
	}

	return nil
}

func DeleteBook(name *string) *string {
	for index, book := range sliceOfBook {
		if book.Name == *name {
			sliceOfBook = append(sliceOfBook[:index], sliceOfBook[index+1:]...)

			return &book.Id
		}
	}

	return nil
}
