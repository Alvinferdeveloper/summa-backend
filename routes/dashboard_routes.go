package routes

import (
	"github.com/Alvinferdeveloper/summa-backend/controllers"
	"github.com/Alvinferdeveloper/summa-backend/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupDashboardRoutes(router *gin.RouterGroup) {
	dashboard := router.Group("/employer/dashboard")
	dashboard.Use(middlewares.AuthMiddleware("employer"))
	{
		dashboard.GET("/stats", controllers.GetDashboardStats)
		dashboard.GET("/pipeline", controllers.GetPipeline)
		dashboard.GET("/candidate-insights/skills", controllers.GetCandidateSkillInsights)
		dashboard.GET("/candidate-insights/disabilities", controllers.GetDisabilityInsights)
	}
}
