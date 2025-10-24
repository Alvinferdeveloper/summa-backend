package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// User represents a job seeker who authenticates via an OAuth provider.
type User struct {
	ID         uuid.UUID `gorm:"column:id;primaryKey"`
	Email      string    `json:"email" gorm:"column:email;unique;not null"`
	Provider   string    `json:"provider" gorm:"column:provider;not null"`       // e.g., "google", "linkedin"
	ProviderID string    `json:"provider_id" gorm:"column:provider_id;not null"` // Unique ID from the provider

	// A user has one job seeker profile
	Profile Profile `json:"profile,omitempty" gorm:"foreignKey:UserID"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}
