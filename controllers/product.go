package controllers

import "github.com/gin-gonic/gin"

type ProductInput struct {
	Name  string `json:"name" binding:"required"`
	Price string `json:"price" binding:"required"`
	Image string `json:"image" binding:"required"`
	Desc  string `json:"desc"`
}

func ShowAllProduct(c *gin.Context) {}
func ShowProduct(c *gin.Context)    {}
func CreateProduct(c *gin.Context)  {}
func UpdateProduct(c *gin.Context)  {}
func DeleteProduct(c *gin.Context)  {}
