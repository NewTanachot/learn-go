package model

import "gorm.io/gorm"

type GormBook struct {
	gorm.Model
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Price       uint     `json:"price"`
	UserId      uint     // if name UserID it will be default (no need to set foreignKey)
	User        User     `json:"user" gorm:"foreignKey:UserId"`
	Author      []Author `json:"author" gorm:"many2many:author_books"`
}

// gorm.Model
// type Model struct {
//     ID        uint `gorm:"primarykey"`
//     CreatedAt time.Time
//     UpdatedAt time.Time
//     DeletedAt DeletedAt `gorm:"index"`
// }

// if entity have DeletedAt field it will be auto soft delete
// if it not it will be hard delete
