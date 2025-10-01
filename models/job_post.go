package models

import "gorm.io/gorm"

// JobPost represents a job posting by an Employer.
type JobPost struct {
	gorm.Model
	EmployerID            uint   `json:"employer_id" gorm:"column:employer_id;not null"`
	Title                 string `json:"title" gorm:"column:title;not null"`
	Location              string `json:"location" gorm:"column:location;not null"`
	WorkModel             string `json:"work_model" gorm:"column:work_model;not null"`
	ContractType          string `json:"contract_type" gorm:"column:contract_type;not null"`
	Description           string `json:"description" gorm:"type:text;not null"`
	Responsibilities      string `json:"responsibilities" gorm:"type:text;not null"`
	Requirements          string `json:"requirements" gorm:"type:text;not null"`
	AccessibilityFeatures string `json:"accessibility_features" gorm:"type:text"`

	// Relationships
	Employer Employer `json:"employer,omitempty"`
}
