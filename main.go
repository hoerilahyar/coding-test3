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

	// RBAC routes
	authorize := r.Group("/api/auth")
	authorize.Use(middlewares.JwtAuthMiddleware())
	authorize.Use(middlewares.AdminAuthMiddleware())
	authorize.GET("/role", controllers.ShowAllRole)
	authorize.GET("/permission", controllers.ShowAllPermission)
	authorize.POST("/create-role", controllers.CreateRole)
	authorize.POST("/create-permission", controllers.CreatePermission)
	authorize.PUT("/assign-permissions-to-role", controllers.AssignPermissionsToRole)
	authorize.DELETE("/revoke-permission-from-role", controllers.RevokeRolePermission)
	authorize.PUT("/assign-role-to-user", controllers.AssignRoleToUser)
	authorize.DELETE("/revoke-role-to-user", controllers.RevokeRoleToUser)

	// product routes
	product := r.Group("/api/product")
	product.Use(middlewares.JwtAuthMiddleware())
	product.GET("/", controllers.ShowAllProduct)
	product.GET("/:id", controllers.ShowProduct)

	// protect only admin can create update delete product
	product.Use(middlewares.AdminAuthMiddleware())
	product.POST("/", controllers.CreateProduct)
	product.PUT("/:id", controllers.UpdateProduct)
	product.DELETE("/:id", controllers.DeleteProduct)

	// transaction routes
	transaction := r.Group("/api/transaction")
	transaction.Use(middlewares.JwtAuthMiddleware())
	transaction.GET("/", controllers.ShowAllTransaction)
	transaction.GET("/:id", controllers.ShowTransaction)
	transaction.POST("/add-to-cart", controllers.AddToCartTransaction)
	transaction.POST("/checkout/:id", controllers.CheckoutTransaction)

	r.Run(":8080")

}
