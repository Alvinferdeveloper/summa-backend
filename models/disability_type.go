package models

import "gorm.io/gorm"

// DisabilityType represents a category of disability.
// This provides a structured way for profiles to self-identify if they choose to.
type DisabilityType struct {
	gorm.Model
	Name        string `json:"name" gorm:"column:name;unique;not null"`
	Description string `json:"description" gorm:"type:text"`

	// Relationships
	Profiles []Profile `json:"-" gorm:"many2many:profile_disability_types;"`
}
