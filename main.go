package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hoerilahyar/coding-test3/controllers"
	"github.com/hoerilahyar/coding-test3/middlewares"
	"github.com/hoerilahyar/coding-test3/models"
)

func main() {

	models.ConnectDataBase()

	r := gin.Default()

	auth := r.Group("/api")

	auth.POST("/register", controllers.Register)
	auth.POST("/login", controllers.Login)
	// auth.POST("/login", controllers.Login)

	authorize := r.Group("/api/auth")
	authorize.Use(middlewares.JwtAuthMiddleware())
	authorize.Use(middlewares.Authorize(true))
	authorize.GET("/role", controllers.ShowAllRole)
	authorize.GET("/permission", controllers.ShowAllPermission)
	authorize.POST("/create-role", controllers.CreateRole)
	authorize.POST("/create-permission", controllers.CreatePermission)
	authorize.PUT("/assign-role", controllers.AssignRole)
	authorize.PUT("/assign-permission", controllers.AssignPermission)
	authorize.DELETE("/revoke-role", controllers.RevokeRole)
	authorize.DELETE("/revoke-permission", controllers.RevokePermission)
	authorize.PUT("/assign-role-to-user", controllers.AssignRoleToUser)
	authorize.PUT("/assign-permission-to-user", controllers.AssignPermissionToUser)
	authorize.DELETE("/revoke-role-to-user", controllers.RevokeRoleToUser)
	authorize.DELETE("/revoke-permission-to-user", controllers.RevokePermissionToUser)

	product := r.Group("/api/product")
	product.Use(middlewares.JwtAuthMiddleware())
	product.GET("/", controllers.ShowAllProduct)
	product.GET("/show", controllers.ShowProduct)
	product.POST("/create", controllers.CreateProduct)
	product.PUT("/update", controllers.UpdateProduct)
	product.DELETE("/delete", controllers.DeleteProduct)

	transaction := r.Group("/api/transaction")
	transaction.Use(middlewares.JwtAuthMiddleware())
	transaction.GET("/show", controllers.ShowTransaction)
	transaction.POST("/add-to-cart", controllers.AddToCartTransaction)
	transaction.POST("/checkout", controllers.CheckoutTransaction)

	r.Run(":8080")

}
