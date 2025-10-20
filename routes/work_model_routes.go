package routes

import (
	"github.com/Alvinferdeveloper/summa-backend/controllers"
	"github.com/gin-gonic/gin"
)

// SetupWorkModelRoutes sets up the routes for work models.
func SetupWorkModelRoutes(router *gin.RouterGroup) {
	workModelRoutes := router.Group("/work-models")
	{
		workModelRoutes.GET("", controllers.GetAllWorkModels)
	}
}
