package controllers

import "github.com/gin-gonic/gin"

type TransactionHeaderInput struct {
	UserID     string `json:"user_id" binding:"required"`
	GrandTotal string `json:"grand_total" binding:"required"`
	Discount   string `json:"discount" binding:"required"`
}

type TransactionDetailInput struct {
	ProductID     string `json:"product_id" binding:"required"`
	TransactionID string `json:"transaction_id" binding:"required"`
	Amount        string `json:"amount" binding:"required"`
	Qty           string `json:"qty" binding:"required"`
}

func ShowTransaction(c *gin.Context)      {}
func AddToCartTransaction(c *gin.Context) {}
func CheckoutTransaction(c *gin.Context)  {}
