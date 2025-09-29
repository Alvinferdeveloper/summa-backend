package models

import "gorm.io/gorm"

// AccessibilityNeed represents a specific accommodation or requirement for a job.
// This is the primary model for matching profiles with suitable job positions.
type AccessibilityNeed struct {
	gorm.Model
	Name     string `json:"name" gorm:"column:name;unique;not null"`
	Category string `json:"category" gorm:"column:category;not null"` // e.g., "Physical Environment", "Software", "Support"

	// Relationships
	Profiles     []Profile     `json:"-" gorm:"many2many:profile_accessibility_needs;"`
	JobPositions []JobPosition `json:"-" gorm:"many2many:job_position_accessibility_features;"`
}
