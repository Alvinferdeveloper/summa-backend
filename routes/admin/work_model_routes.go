package admin

import (
	"github.com/Alvinferdeveloper/summa-backend/controllers/admin"
	"github.com/Alvinferdeveloper/summa-backend/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupWorkModelRoutes(router *gin.RouterGroup) {
	workModels := router.Group("/admin/work-models")
	workModels.Use(middlewares.AuthMiddleware("admin"))
	{
		workModels.GET("", admin.GetAllWorkModels)
		workModels.POST("", admin.CreateWorkModel)
		workModels.PUT("/:id", admin.UpdateWorkModel)
		workModels.DELETE("/:id", admin.DeleteWorkModel)
	}
}
