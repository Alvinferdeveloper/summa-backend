package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Notification struct {
	gorm.Model
	UserID     *uuid.UUID `json:"user_id,omitempty" gorm:"column:user_id"`
	EmployerID *uuid.UUID `json:"employer_id,omitempty" gorm:"column:employer_id"`
	Message    string     `json:"message" gorm:"column:message;not null"`
	IsRead     bool       `json:"is_read" gorm:"column:is_read;default:false"`
	Link       string     `json:"link,omitempty" gorm:"column:link"` //optional link to redirect the user
}
