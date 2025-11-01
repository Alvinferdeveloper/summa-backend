package admin

import (
	"github.com/Alvinferdeveloper/summa-backend/controllers/admin"
	"github.com/Alvinferdeveloper/summa-backend/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupExperienceLevelRoutes(router *gin.RouterGroup) {
	experienceLevels := router.Group("/admin/experience-levels")
	experienceLevels.Use(middlewares.AuthMiddleware("admin"))
	{
		experienceLevels.GET("", admin.GetAllExperienceLevels)
		experienceLevels.POST("", admin.CreateExperienceLevel)
		experienceLevels.PUT("/:id", admin.UpdateExperienceLevel)
		experienceLevels.DELETE("/:id", admin.DeleteExperienceLevel)
	}
}
