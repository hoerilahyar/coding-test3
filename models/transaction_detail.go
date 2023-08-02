package models

import "gorm.io/gorm"

type TransactionDetail struct {
	gorm.Model
	Qty               int `json:"qty"`
	Amount            int `json:"amount"`
	ProductID         int
	TransactionID     int
	Product           Product
	TransactionHeader TransactionHeader
}
