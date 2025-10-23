package routes

import (
	"github.com/Alvinferdeveloper/summa-backend/controllers"
	"github.com/Alvinferdeveloper/summa-backend/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupChatRoutes(router *gin.RouterGroup) {
	chatRoutes := router.Group("/chat")
	{
		authMiddleware := middlewares.AuthMiddleware("job_seeker", "employer")

		chatRoutes.GET("/conversations", authMiddleware, controllers.GetConversations)
		chatRoutes.POST("/conversations/with/:otherParticipantId", authMiddleware, controllers.GetOrCreateConversation)
		chatRoutes.GET("/conversations/:conversationId/messages", authMiddleware, controllers.GetMessagesForConversation)
	}
}
