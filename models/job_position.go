package models

import "gorm.io/gorm"

// JobPosition represents a job posting by an Employer.
type JobPosition struct {
	gorm.Model
	EmployerID  uint   `json:"employer_id" gorm:"column:employer_id;not null"`
	Title       string `json:"title" gorm:"column:title;not null"`
	Description string `json:"description" gorm:"column:description"`
	Location    string `json:"location" gorm:"column:location"`
	Status      string `json:"status" gorm:"column:status;not null;default:'open'"`

	// Relationships
	Employer Employer `json:"employer,omitempty"`
}
