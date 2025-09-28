package models

import "gorm.io/gorm"

// UniversitySuggestion is for users to suggest new universities.
type UniversitySuggestion struct {
	gorm.Model
	SuggestedName string `json:"suggested_name" gorm:"column:suggested_name;not null"`
	Country       string `json:"country" gorm:"column:country"`
	Website       string `json:"website" gorm:"column:website"`
	Status        string `json:"status" gorm:"column:status;not null;default:'pending'"` // 'pending', 'approved', 'rejected'
	SuggestedBy   uint   `json:"suggested_by" gorm:"column:suggested_by;not null"`

	// Relationships
	User User `json:"user,omitempty" gorm:"foreignKey:SuggestedBy"`
}
