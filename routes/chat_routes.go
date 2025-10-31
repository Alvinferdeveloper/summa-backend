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
		conversationRoutes := chatRoutes.Group("/conversations")
		{
			conversationRoutes.GET("", authMiddleware, controllers.GetConversations)
			conversationRoutes.POST("/with/:otherParticipantId", authMiddleware, controllers.GetOrCreateConversation)
			conversationRoutes.GET("/:conversationId/messages", authMiddleware, controllers.GetMessagesForConversation)
			conversationRoutes.POST("/:conversationId/mark-as-read", authMiddleware, controllers.MarkConversationAsRead)
		}
	}
}
