package models

import "gorm.io/gorm"

type TransactionHeader struct {
	gorm.Model
	Discount   int  `json:"discount"`
	GrandTotal int  `json:"grand_total"`
	Status     bool `json:"status"`
	UserID     int
	User       User
}
