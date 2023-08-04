package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hoerilahyar/coding-test3/models"
	"gorm.io/gorm"
)

func ShowAllUsers(c *gin.Context) {
	var user []models.User

	if err := models.DB.Find(&user).Error; err != nil || errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Here is data!", "data": user})
}

func ShowUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var user models.User

	if err := models.DB.Model(models.User{}).Where("id = ?", id).First(&user).Error; err != nil || errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Here is data!", "data": user})
}

func CreateUser(c *gin.Context) {
	var input RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := models.User{}

	u.Fullname = input.Fullname
	u.Username = input.Username
	u.Email = input.Email
	u.Password = input.Password

	_, err := u.SaveUser()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "create user success"})
}

func UpdateUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var user models.User
	var input RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.DB.Model(models.User{}).Where("id = ?", id).First(&user).Error; err != nil || errors.Is(err, gorm.ErrRecordNotFound) {

		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	if input.Fullname != "" {
		user.Fullname = input.Fullname
	}

	if input.Username != "" {
		user.Username = input.Username
	}

	if input.Email != "" {
		user.Email = input.Email
	}

	if input.Password != "" {
		user.Password = input.Password
	}

	models.DB.Save(&user)

	c.JSON(http.StatusCreated, gin.H{"message": "updated", "data": user})
}

func DeleteUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var user models.User

	if err := models.DB.Model(models.User{}).Where("id = ?", id).Delete(&user).Error; err != nil || errors.Is(err, gorm.ErrRecordNotFound) {

		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "deleted", "data": user})
}
