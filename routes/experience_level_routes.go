package routes

import (
	"github.com/Alvinferdeveloper/summa-backend/controllers"
	"github.com/gin-gonic/gin"
)

// SetupExperienceLevelRoutes sets up the routes for experience levels.
func SetupExperienceLevelRoutes(router *gin.RouterGroup) {
	experienceLevelRoutes := router.Group("/experience-levels")
	{
		experienceLevelRoutes.GET("", controllers.GetAllExperienceLevels)
	}
}
