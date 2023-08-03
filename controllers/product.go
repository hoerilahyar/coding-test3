package controllers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hoerilahyar/coding-test3/models"
	"gorm.io/gorm"
)

type ProductCreate struct {
	Name  string `json:"name" binding:"required"`
	Price int    `json:"price" binding:"required"`
	Image string `json:"image" binding:"required"`
	Desc  string `json:"desc"`
}
type ProductUpdate struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
	Image string `json:"image"`
	Desc  string `json:"desc"`
}

func ShowAllProduct(c *gin.Context) {
	var products []models.Product

	if err := models.DB.Find(&products).Error; err != nil || errors.Is(err, gorm.ErrRecordNotFound) {

		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Here is data!", "data": products})
}

func ShowProduct(c *gin.Context) {
	id := c.Params.ByName("id")
	var product models.Product

	if err := models.DB.Model(models.Product{}).Where("id = ?", id).First(&product).Error; err != nil || errors.Is(err, gorm.ErrRecordNotFound) {

		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Here is data!", "data": product})
}

func CreateProduct(c *gin.Context) {

	var input ProductCreate

	if err := c.ShouldBindJSON(&input); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	p := models.Product{}

	p.Name = input.Name
	p.Price = input.Price
	p.Image = input.Image
	p.Desc = input.Desc

	_, err := p.SaveProduct()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "product added"})
}

func UpdateProduct(c *gin.Context) {
	id := c.Params.ByName("id")
	var product models.Product
	var input ProductUpdate

	if err := c.ShouldBindJSON(&input); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.DB.Model(models.Product{}).Where("id = ?", id).First(&product).Error; err != nil || errors.Is(err, gorm.ErrRecordNotFound) {

		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	if input.Name != "" {
		product.Name = input.Name
	}

	if input.Price != 0 {
		product.Price = input.Price
	}

	if input.Desc != "" {
		product.Desc = input.Desc
	}

	if input.Image != "" {
		product.Image = input.Image
	}

	models.DB.Save(&product)

	c.JSON(http.StatusCreated, gin.H{"message": "updated", "data": product})

}

func DeleteProduct(c *gin.Context) {
	id := c.Params.ByName("id")
	var product models.Product

	if err := models.DB.Model(models.Product{}).Where("id = ?", id).Delete(&product).Error; err != nil || errors.Is(err, gorm.ErrRecordNotFound) {

		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "deleted", "data": product})

}
