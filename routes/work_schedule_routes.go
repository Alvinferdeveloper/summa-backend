package routes

import (
	"github.com/Alvinferdeveloper/summa-backend/controllers"
	"github.com/gin-gonic/gin"
)

// SetupWorkScheduleRoutes sets up the routes for work schedules.
func SetupWorkScheduleRoutes(router *gin.RouterGroup) {
	workScheduleRoutes := router.Group("/work-schedules")
	{
		workScheduleRoutes.GET("", controllers.GetAllWorkSchedules)
	}
}
