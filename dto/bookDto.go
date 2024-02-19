package dto

type BookRequestDto struct {
	Book   bookDto   `json:"book"`
	User   userDto   `json:"user"`
	Author authorDto `json:"author"`
}

type bookDto struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       uint   `json:"price"`
}

type userDto struct {
	Id uint `json:"id"`
}

type authorDto struct {
	Id uint `json:"id"`
}
