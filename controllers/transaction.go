package controllers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hoerilahyar/coding-test3/models"
	"github.com/hoerilahyar/coding-test3/utils/token"
	"gorm.io/gorm"
)

type TransactionProduct struct {
	ProductID int `json:"product_id" binding:"required"`
	Qty       int `json:"qty" binding:"required"`
}

type TransactionInput struct {
	Product  []TransactionProduct `json:"product" binding:"required"`
	Discount int                  `json:"discount"`
}

func ShowAllTransaction(c *gin.Context) {
	var transactions []models.TransactionHeader

	if err := models.DB.Preload("TransactionDetail").Find(&transactions).Error; err != nil || errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Here is data!", "data": transactions})
}

func ShowTransaction(c *gin.Context) {
	id := c.Params.ByName("id")
	var transaction models.TransactionHeader

	if err := models.DB.Preload("TransactionDetail").Model(models.TransactionHeader{}).Where("id = ?", id).First(&transaction).Error; err != nil || errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	fmt.Println(transaction)
	c.JSON(http.StatusOK, gin.H{"message": "Here is data!", "data": transaction})
}

func AddToCartTransaction(c *gin.Context) {
	var input TransactionInput
	var errProduct []string
	// var product TransactionProduct

	if err := c.ShouldBindJSON(&input); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(input.Product) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "product must be fill"})
		return
	}

	user_id, err := token.ExtractTokenID(c)

	var user models.User

	err = models.DB.Model(models.User{}).Where("id = ?", user_id).First(&user).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		c.Abort()
		return
	}

	th := models.TransactionHeader{
		UserID:     user_id,
		Discount:   input.Discount,
		Status:     false,
		GrandTotal: 0,
	}

	_ = models.DB.Transaction(func(tx *gorm.DB) error {
		tx.Create(&th)
		return nil
	})

	for _, element := range input.Product {

		var product models.Product

		err = models.DB.Model(models.Product{}).Where("id = ?", element.ProductID).First(&product).Error

		amout := product.Price * element.Qty

		td := models.TransactionDetail{
			ProductID:     int(product.ID),
			Qty:           element.Qty,
			TransactionID: uint(th.ID),
			Amount:        product.Price * element.Qty,
		}

		err = models.DB.Create(&td).Error

		if err != nil {
			errProduct = append(errProduct, err.Error())
		}

		th.GrandTotal = amout
		models.DB.Save(&th)
	}

	fmt.Println(errProduct)

	c.JSON(http.StatusCreated, gin.H{"message": "Product added!", "data": th})

}

func CheckoutTransaction(c *gin.Context) {
	var transaction models.TransactionHeader
	var user models.User

	id := c.Params.ByName("id")
	user_id, err := token.ExtractTokenID(c)

	err = models.DB.Model(models.User{}).Where("id = ?", user_id).First(&user).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		c.Abort()
		return
	}

	err = models.DB.Model(models.TransactionHeader{}).Where("id = ?", id).First(&transaction).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Transaction not found"})
		c.Abort()
		return
	}

	transaction.Status = true

	models.DB.Save(&transaction)

	c.JSON(http.StatusCreated, gin.H{"message": "Transaction has been checkout!", "data": transaction})
}
