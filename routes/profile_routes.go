package routes

import (
	"github.com/Alvinferdeveloper/summa-backend/controllers"
	"github.com/Alvinferdeveloper/summa-backend/middlewares"
	"github.com/gin-gonic/gin"
)

// SetupProfileRoutes registers the profile-related routes.
func SetupProfileRoutes(router *gin.RouterGroup) {
	router.GET("/disability-types", controllers.GetDisabilityTypes)

	profile := router.Group("/profile")
	{
		profile.PUT("/", middlewares.AuthMiddleware(), controllers.CompleteOnboarding)
	}
}
