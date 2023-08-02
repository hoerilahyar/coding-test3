package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hoerilahyar/coding-test3/models"
	"github.com/hoerilahyar/coding-test3/utils/token"
)

func CurrentUser(c *gin.Context) {

	user_id, err := token.ExtractTokenID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := models.GetUserByID(user_id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": u})
}

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {

	var input LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := models.User{}

	u.Username = input.Username
	u.Password = input.Password

	token, err := models.LoginCheck(u.Username, u.Password)
	// fmt.Println(err)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username or password is incorrect."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})

}

type RegisterInput struct {
	Fullname string `json:"fullname" binding:"required"`
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Register(c *gin.Context) {
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

	c.JSON(http.StatusOK, gin.H{"message": "registration success"})
}

func CreateRole(c *gin.Context)             {}
func CreatePermission(c *gin.Context)       {}
func ShowAllRole(c *gin.Context)            {}
func ShowAllPermission(c *gin.Context)      {}
func AssignRole(c *gin.Context)             {}
func AssignPermission(c *gin.Context)       {}
func RevokeRole(c *gin.Context)             {}
func RevokePermission(c *gin.Context)       {}
func AssignRoleToUser(c *gin.Context)       {}
func AssignPermissionToUser(c *gin.Context) {}
func RevokeRoleToUser(c *gin.Context)       {}
func RevokePermissionToUser(c *gin.Context) {}
