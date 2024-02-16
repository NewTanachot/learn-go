package model

import "gorm.io/gorm"

type GormBook struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Description string `json:"description"`
	Price       uint   `json:"price"`
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
