package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	ConversationID uint      `json:"conversation_id" gorm:"not null"`
	SenderID       uuid.UUID `json:"sender_id" gorm:"not null"`
	SenderType     string    `json:"sender_type" gorm:"not null"` // "user" or "employer"
	RecipientID    uuid.UUID `json:"recipient_id" gorm:"not null"`
	RecipientType  string    `json:"recipient_type" gorm:"not null"` // "user" or "employer"
	Content        string    `json:"content" gorm:"type:text;not null"`
	Read           bool      `json:"read" gorm:"default:false"`
}
