package routes

import (
	"github.com/Alvinferdeveloper/summa-backend/controllers"
	"github.com/gin-gonic/gin"
)

// SetupSkillRoutes sets up the routes for skills.
func SetupSkillRoutes(router *gin.RouterGroup) {
	skillRoutes := router.Group("/skills")
	{
		skillRoutes.GET("", controllers.GetAllSkills)
	}
}
