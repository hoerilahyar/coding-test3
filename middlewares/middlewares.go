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
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}
		c.Next()
	}
}

func Authorize(isAdmin bool) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user_id, err := token.ExtractTokenID(ctx)

		if err != nil {
			var user models.User

			err := models.DB.Model(models.User{}).Where("id = ?", user_id).First(&user).Error

			if err != nil {
				ctx.String(http.StatusBadRequest, "User not found")
				ctx.Abort()
				return
			}

			ok, err := models.AUTH.CheckUserRole(user.ID, "admin")

			if err != nil {
				ctx.String(http.StatusUnauthorized, "Unauthorized")
				ctx.Abort()
				return
			}

			_ = ok
		}
		ctx.Next()
	}
}
