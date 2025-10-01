package routes

import (
	"github.com/Alvinferdeveloper/summa-backend/controllers"
	"github.com/Alvinferdeveloper/summa-backend/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupJobPostRoutes(router *gin.RouterGroup) {
	jobPost := router.Group("/jobs")
	{
		jobPost.POST("", middlewares.AuthMiddleware("employer"), controllers.CreateJobPost)
		jobPost.GET("", controllers.GetJobPosts)
	}
}
