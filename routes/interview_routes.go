package routes

import (
	"github.com/Alvinferdeveloper/summa-backend/controllers"
	"github.com/Alvinferdeveloper/summa-backend/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupInterviewRoutes(router *gin.RouterGroup) {
	interviews := router.Group("/interviews")
	{
		interviews.POST("", middlewares.AuthMiddleware("employer"), controllers.ScheduleInterview)
		interviews.PUT("/:id/respond", middlewares.AuthMiddleware("job_seeker"), controllers.RespondToInterview)
		interviews.GET("/:id/download-ics", middlewares.AuthMiddleware("job_seeker", "employer"), controllers.DownloadICS)
	}
}
