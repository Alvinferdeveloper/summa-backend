package models

import "gorm.io/gorm"

// JobPosition represents a job posting by an Employer.
type JobPosition struct {
	gorm.Model
	EmployerID  uint   `json:"employer_id" gorm:"column:employer_id;not null"`
	Title       string `json:"title" gorm:"column:title;not null"`
	Description string `json:"description" gorm:"column:description"`
	Location    string `json:"location" gorm:"column:location"`
	WorkModel   string `json:"work_model" gorm:"column:work_model"` // e.g., OnSite, Hybrid, Remote
	Status      string `json:"status" gorm:"column:status;not null;default:'open'"`
	IsExclusivelyForDisabled bool `json:"is_exclusively_for_disabled" gorm:"column:is_exclusively_for_disabled;default:true"`
	AccessibilityDescription string `json:"accessibility_description" gorm:"type:text"`

	// Relationships
	Employer Employer `json:"employer,omitempty"`
	AccessibilityFeatures []AccessibilityNeed `json:"accessibility_features,omitempty" gorm:"many2many:job_position_accessibility_features;"`
}
