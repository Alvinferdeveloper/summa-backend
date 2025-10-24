package dto

import (
	"github.com/Alvinferdeveloper/summa-backend/models"
	"github.com/google/uuid"
)

type MessageResponseDTO struct {
	ID             uint      `json:"id"`
	ConversationID uint      `json:"conversation_id"`
	SenderID       uuid.UUID `json:"sender_id"`
	SenderType     string    `json:"sender_type"`
	RecipientID    uuid.UUID `json:"recipient_id"`
	RecipientType  string    `json:"recipient_type"`
	Content        string    `json:"content"`
	Read           bool      `json:"read"`
}

func ConvertMessageToDTO(message models.Message) *MessageResponseDTO {
	return &MessageResponseDTO{
		ID:             message.ID,
		ConversationID: message.ConversationID,
		SenderID:       message.SenderID,
		SenderType:     message.SenderType,
		RecipientID:    message.RecipientID,
		RecipientType:  message.RecipientType,
		Content:        message.Content,
		Read:           message.Read,
	}
}
