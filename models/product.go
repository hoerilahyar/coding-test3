package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name  string `json:"name" gorm:"required"`
	Desc  string `json:"description"`
	Image string `json:"image"`
	Price string `json:"price"`
}
