package routes

import (
	"github.com/Alvinferdeveloper/summa-backend/controllers"
	"github.com/Alvinferdeveloper/summa-backend/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupJobApplicationRoutes(router *gin.RouterGroup) {
	router.POST("/jobs/:id/apply", middlewares.AuthMiddleware("job_seeker"), controllers.ApplyToJob)
	router.GET("/applications", middlewares.AuthMiddleware("job_seeker"), controllers.GetMyApplications)
}
