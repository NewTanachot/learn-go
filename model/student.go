package model

type Student struct {
	Id    int    `json:"id" validate:"required"`
	Name  string `json:"name" validate:"required,fullname"`
	Grade string `json:"grade" validate:"required"`
	Score int    `json:"score" validate:"required,min=0,max=100"`
}
