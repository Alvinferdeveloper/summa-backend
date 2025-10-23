package routes

import (
	"github.com/Alvinferdeveloper/summa-backend/controllers"
	"github.com/Alvinferdeveloper/summa-backend/middlewares"
	"github.com/Alvinferdeveloper/summa-backend/websocket"
	"github.com/gin-gonic/gin"
)

func SetupWebSocketRoutes(router *gin.RouterGroup, hub *websocket.Hub) {
	wsRoutes := router.Group("/ws")
	{
		authMiddleware := middlewares.AuthMiddleware("job_seeker", "employer")

		wsRoutes.GET("/chat", authMiddleware, controllers.ServeWebSocket(hub))
	}
}
