package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hoerilahyar/coding-test3/models"
	"github.com/hoerilahyar/coding-test3/utils/token"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := token.TokenValid(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		c.Next()
	}
}

func AdminAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user_id, err := token.ExtractTokenID(ctx)

		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "token aborted"})
			ctx.Abort()
			return
		}

		var user models.User

		err = models.DB.Model(models.User{}).Where("id = ?", user_id).First(&user).Error

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
			ctx.Abort()
			return
		}

		ok, err := models.AUTH.CheckUserRole(user.ID, "admin")

		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			ctx.Abort()
			return
		} else if ok == false {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unatuhorized"})
			ctx.Abort()
			return
		}

		_ = ok
		ctx.Next()
	}
}
