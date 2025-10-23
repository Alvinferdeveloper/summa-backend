package models

import "gorm.io/gorm"

type Conversation struct {
	gorm.Model
	UserID     uint `json:"user_id" gorm:"not null"`
	EmployerID uint `json:"employer_id" gorm:"not null"`

	User     User     `gorm:"foreignKey:UserID"`
	Employer Employer `gorm:"foreignKey:EmployerID"`

	Messages []Message `gorm:"foreignKey:ConversationID"`
}