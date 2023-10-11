package routes

import (
	"sportscorner/pkg/api/handler"
	"sportscorner/pkg/api/middleware"

	"github.com/gin-gonic/gin"
)

func AdminRoutes(engine *gin.RouterGroup, adminHandler *handler.AdminHandler, productHandler *handler.ProductHandler, userHandler *handler.UserHandler, categoryHandler *handler.CategoryHandler,orderHandler *handler.OrderHandler) {
	engine.POST("/adminlogin", adminHandler.LoginHandler)

	engine.Use(middleware.AdminAuthMiddleware)
	{
		usermanagement := engine.Group("/users")
		{
			usermanagement.POST("/block", adminHandler.BlockUser)
			usermanagement.POST("/unblock", adminHandler.UnBlockUser)
			usermanagement.GET("/getusers", adminHandler.GetUsers)
		}

		categorymanagement := engine.Group("/category")
		{
			categorymanagement.POST("/add", categoryHandler.AddCategory)
			categorymanagement.PUT("/update", categoryHandler.UpdateCategory)
			categorymanagement.DELETE("/delete", categoryHandler.DeleteCategory)
		}

		productmanagement := engine.Group("/product")
		{
			productmanagement.POST("/add", productHandler.AddProduct)
			productmanagement.PUT("/update", productHandler.UpdateProduct)
			productmanagement.DELETE("/delete", productHandler.DeleteProduct)
		}
	}
}
