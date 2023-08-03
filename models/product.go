package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name  string `json:"name" gorm:"required"`
	Desc  string `json:"description"`
	Image string `json:"image"`
	Price int    `json:"price"`
}

func (p *Product) SaveProduct() (*Product, error) {

	var err error
	err = DB.Create(&p).Error
	if err != nil {
		return &Product{}, err
	}
	return p, nil
}
