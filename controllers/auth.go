package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/harranali/authority"
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

	c.JSON(http.StatusOK, gin.H{"message": "hooray! you got the token", "token": token})

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

	c.JSON(http.StatusCreated, gin.H{"message": "registration success"})
}

type RbacInput struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type AssignInput struct {
	Roles       []string `json:"roles"`
	Role        string   `json:"role"`
	Permissions []string `json:"permissions"`
	Permission  string   `json:"permission"`
	Users       []string `json:"users"`
	User        string   `json:"user"`
}

func CreateRole(c *gin.Context) {
	var input RbacInput
	var tx = models.AUTH.BeginTX()

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := tx.CreateRole(authority.Role{
		Name: input.Name,
		Slug: input.Slug,
	})

	if err != nil {
		tx.Rollback() // transaction rollback incase of error
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Role added"})

}

func CreatePermission(c *gin.Context) {
	var input RbacInput
	var tx = models.AUTH.BeginTX()

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := tx.CreatePermission(authority.Permission{
		Name: input.Name,
		Slug: input.Slug,
	})

	if err != nil {
		tx.Rollback() // transaction rollback incase of error
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Permission added"})
}

func ShowAllRole(c *gin.Context) {
	roles, err := models.AUTH.GetAllRoles()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Here is data!", "data": roles})
}

func ShowAllPermission(c *gin.Context) {
	permissions, err := models.AUTH.GetAllPermissions()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Here is data!", "data": permissions})
}

func AssignPermissionsToRole(c *gin.Context) {
	var input AssignInput
	var tx = models.AUTH.BeginTX()

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.Permissions == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Permission empty"})
		return
	}

	if input.Role == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Role empty"})
		return
	}

	err := tx.AssignPermissionsToRole(input.Role, input.Permissions)

	if err != nil {
		tx.Rollback() // transaction rollback incase of error
		c.JSON(http.StatusBadRequest, gin.H{"error": "assign failed"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Permission assigned"})
}

func RevokeRolePermission(c *gin.Context) {
	var input AssignInput
	var tx = models.AUTH.BeginTX()

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.Permission == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Permission empty"})
		return
	}

	if input.Role == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Role empty"})
		return
	}

	err := tx.RevokeRolePermission(input.Role, input.Permission)

	if err != nil {
		tx.Rollback() // transaction rollback incase of error
		c.JSON(http.StatusBadRequest, gin.H{"error": "revoke failed"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Permission revoked"})

}

func AssignRoleToUser(c *gin.Context) {
	var input AssignInput
	var tx = models.AUTH.BeginTX()

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.User == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User empty"})
		return
	}

	if input.Role == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Role empty"})
		return
	}

	err := tx.AssignRoleToUser(input.User, input.Role)

	if err != nil {
		tx.Rollback() // transaction rollback incase of error
		c.JSON(http.StatusBadRequest, gin.H{"error": "assign failed"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Role assigned to user"})

}

func RevokeRoleToUser(c *gin.Context) {
	var input AssignInput
	var tx = models.AUTH.BeginTX()

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.User == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User empty"})
		return
	}

	if input.Role == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Role empty"})
		return
	}

	err := tx.RevokeUserRole(input.User, input.Role)

	if err != nil {
		tx.Rollback() // transaction rollback incase of error
		c.JSON(http.StatusBadRequest, gin.H{"error": "revoke failed"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Role revoke from user"})
}
