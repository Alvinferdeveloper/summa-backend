package controllers

import (
	"net/http"
	"strconv"

	"github.com/Alvinferdeveloper/summa-backend/dto"
	"github.com/Alvinferdeveloper/summa-backend/services"
	"github.com/Alvinferdeveloper/summa-backend/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetConversations(c *gin.Context) {
	participantID, participantType, err := utils.GetParticipantFromContext(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	conversationsDTO, err := services.GetConversations(participantID, participantType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve conversations"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": conversationsDTO})
}

func GetOrCreateConversation(c *gin.Context) {
	participantID, participantType, err := utils.GetParticipantFromContext(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	otherParticipantIDStr := c.Param("otherParticipantId")
	otherParticipantID, err := uuid.Parse(otherParticipantIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid participant ID"})
		return
	}

	var userID, employerID uuid.UUID
	switch participantType {
	case "user":
		userID = participantID
		employerID = otherParticipantID
	case "employer":
		userID = otherParticipantID
		employerID = participantID
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid participant type"})
		return
	}

	conversation, err := services.GetOrCreateConversation(userID, employerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get or create conversation"})
		return
	}

	conversationDTO := dto.ConvertConversationToDTO(*conversation, 0)

	c.JSON(http.StatusOK, gin.H{"data": conversationDTO})
}

func GetMessagesForConversation(c *gin.Context) {
	conversationIDStr := c.Param("conversationId")
	conversationID, err := strconv.ParseUint(conversationIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid conversation ID"})
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))

	messages, total, err := services.GetMessagesForConversation(uint(conversationID), page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve messages"})
		return
	}

	var dtoMessages []dto.MessageResponseDTO = make([]dto.MessageResponseDTO, 0)
	for _, message := range messages {
		dtoMessages = append(dtoMessages, *dto.ConvertMessageToDTO(message))
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  dtoMessages,
		"total": total,
		"page":  page,
		"limit": limit,
	})
}

func MarkConversationAsRead(c *gin.Context) {
	userType, _ := c.Get("user_type")
	var userID uuid.UUID
	if userType == "job_seeker" {
		if val, exists := c.Get("user_id"); exists {
			userID = val.(uuid.UUID)
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "user_id no encontrado"})
			return
		}
	} else {
		if val, exists := c.Get("employer_id"); exists {
			userID = val.(uuid.UUID)
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "employer_id no encontrado"})
			return
		}
	}

	conversationID, err := strconv.Atoi(c.Param("conversationId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de conversación inválido"})
		return
	}

	if err := services.MarkConversationAsRead(uint(conversationID), userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al marcar la conversación como leída"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Conversación marcada como leída"})
}
