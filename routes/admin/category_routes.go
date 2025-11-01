package admin

import (
	"github.com/Alvinferdeveloper/summa-backend/controllers/admin"
	"github.com/Alvinferdeveloper/summa-backend/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupCategoryRoutes(router *gin.RouterGroup) {
	categories := router.Group("/admin/categories")
	categories.Use(middlewares.AuthMiddleware("admin"))
	{
		categories.GET("", admin.GetAllCategories)
		categories.POST("", admin.CreateCategory)
		categories.PUT("/:id", admin.UpdateCategory)
		categories.DELETE("/:id", admin.DeleteCategory)
	}
}
