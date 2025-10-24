package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Conversation struct {
	gorm.Model
	UserID     uuid.UUID `json:"user_id" gorm:"not null"`
	EmployerID uuid.UUID `json:"employer_id" gorm:"not null"`

	User     User     `gorm:"foreignKey:UserID"`
	Employer Employer `gorm:"foreignKey:EmployerID"`

	Messages []Message `gorm:"foreignKey:ConversationID"`
}
