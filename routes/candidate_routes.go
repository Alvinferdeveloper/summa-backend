package routes

import (
	"github.com/Alvinferdeveloper/summa-backend/controllers"
	"github.com/Alvinferdeveloper/summa-backend/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupCandidateRoutes(router *gin.RouterGroup) {
	candidateRoutes := router.Group("/talent-pool")
	candidateRoutes.Use(middlewares.AuthMiddleware("employer"))
	{
		candidateRoutes.GET("/candidates", controllers.GetCandidates)
	}
}
