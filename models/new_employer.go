package models

import "gorm.io/gorm"

// NewEmployer is a suggestion table for when a candidate's employer is not in the main Employer table.
type NewEmployer struct {
	gorm.Model
	CompanyName string `json:"company_name" gorm:"column:company_name;not null"`
	Website     string `json:"website" gorm:"column:website"`
	SuggestedBy uint   `json:"suggested_by" gorm:"column:suggested_by;not null"`
	Status      string `json:"status" gorm:"column:status;not null;default:'pending'"` // 'pending', 'approved', 'rejected'

	// Relationships
	User User `json:"user,omitempty" gorm:"foreignKey:SuggestedBy"`
}
