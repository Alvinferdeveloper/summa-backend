package controllers

import (
	"github.com/Alvinferdeveloper/summa-backend/utils"
	"github.com/Alvinferdeveloper/summa-backend/websocket"
	"github.com/gin-gonic/gin"
)

func ServeWebSocket(hub *websocket.Hub) gin.HandlerFunc {
	return func(c *gin.Context) {
		participantID, participantType, err := utils.GetParticipantFromContext(c)
		if err != nil {
			// If authentication fails, the WebSocket upgrade will be implicitly rejected
			// because we are not calling ServeWs. Gin will handle the response.
			return
		}

		websocket.ServeWs(hub, c.Writer, c.Request, participantID, participantType)
	}
}
