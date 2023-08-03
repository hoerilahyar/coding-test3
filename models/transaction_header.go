package models

import "gorm.io/gorm"

type TransactionHeader struct {
	gorm.Model
	Discount          int  `json:"discount"`
	GrandTotal        int  `json:"grand_total"`
	Status            bool `json:"status"`
	UserID            uint
	User              User
	TransactionDetail []TransactionDetail `gorm:"foreignKey:TransactionID"`
}
