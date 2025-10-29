package routes

import (
	"github.com/Alvinferdeveloper/summa-backend/controllers"
	"github.com/Alvinferdeveloper/summa-backend/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupNotificationRoutes(router *gin.RouterGroup) {
	notifications := router.Group("/notifications")
	notifications.Use(middlewares.AuthMiddleware("job_seeker", "employer"))
	{
		notifications.GET("", controllers.GetNotifications)
		notifications.POST("/mark-as-read", controllers.MarkNotificationsAsRead)
	}
}
