package routes

import (
	"github.com/Alvinferdeveloper/summa-backend/controllers"
	"github.com/Alvinferdeveloper/summa-backend/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupUniversitySuggestionRoutes(router *gin.RouterGroup) {
	suggestion := router.Group("/university-suggestions")
	{
		suggestion.POST("/", middlewares.AuthMiddleware("job_seeker"), controllers.CreateUniversitySuggestion)
	}
}
