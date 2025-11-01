
package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Admin represents a system administrator.
type Admin struct {
	ID       uuid.UUID `gorm:"type:uuid;primary_key;"`
	Email    string    `gorm:"unique;not null"`
	Password string    `gorm:"not null"`
}

func (admin *Admin) BeforeCreate(tx *gorm.DB) (err error) {
	admin.ID = uuid.New()
	return
}
