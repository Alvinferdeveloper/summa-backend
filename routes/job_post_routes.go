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
		jobPost.GET("", middlewares.AuthMiddleware("job_seeker", "employer"), controllers.GetJobPosts)
		jobPost.GET("/:id", controllers.GetJobPostById)
	}

	employerRoutes := router.Group("/employer")
	{
		employerRoutes.GET("/jobs", middlewares.AuthMiddleware("employer"), controllers.GetEmployerJobPosts)
	}
}
