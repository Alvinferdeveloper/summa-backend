package routes

import (
	"github.com/Alvinferdeveloper/summa-backend/controllers"
	"github.com/Alvinferdeveloper/summa-backend/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupJobApplicationRoutes(router *gin.RouterGroup) {
	router.POST("/jobs/:id/apply", middlewares.AuthMiddleware("job_seeker"), controllers.ApplyToJob)
	router.GET("/applications", middlewares.AuthMiddleware("job_seeker"), controllers.GetMyApplications)
	router.GET("/jobs/:id/applicants", middlewares.AuthMiddleware("employer"), controllers.GetJobApplicants)
	router.GET("/jobs/:id/all-applicants", middlewares.AuthMiddleware("employer"), controllers.GetAllJobApplicants)

	router.PUT("/applications/:id/status", middlewares.AuthMiddleware("employer"), controllers.UpdateApplicationStatus)
}
