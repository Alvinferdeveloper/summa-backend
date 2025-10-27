package routes

import (
	"github.com/Alvinferdeveloper/summa-backend/controllers"
	"github.com/Alvinferdeveloper/summa-backend/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupExperienceRoutes(router *gin.RouterGroup) {
	newEmployer := router.Group("/new-employers")
	{
		newEmployer.POST("", middlewares.AuthMiddleware("job_seeker"), controllers.CreateNewEmployer)
	}

	profile := router.Group("/profile")
	{
		profile.POST("/experience", middlewares.AuthMiddleware("job_seeker"), controllers.CreateExperience)
	}
}
