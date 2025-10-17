package routes

import (
	"github.com/Alvinferdeveloper/summa-backend/controllers"
	"github.com/Alvinferdeveloper/summa-backend/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupUploadRoutes(router *gin.RouterGroup) {
	employer := router.Group("/employer")
	employer.Use(middlewares.AuthMiddleware("employer"))
	{
		employer.POST("/logo", controllers.UploadEmployerLogo)
	}
	profile := router.Group("/profile")
	profile.Use(middlewares.AuthMiddleware("job_seeker"))
	{
		profile.POST("/picture", controllers.UploadProfilePicture)
		profile.POST("/banner", controllers.UploadProfileBanner)
		profile.POST("/cv", controllers.UploadCV)
	}
}
