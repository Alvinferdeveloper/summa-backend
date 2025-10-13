package models

import "gorm.io/gorm"

// JobPost represents a job posting by an Employer.
type JobPost struct {
	gorm.Model
	EmployerID            uint      `json:"employer_id" gorm:"column:employer_id;not null"`
	CategoryID            uint      `json:"category_id" gorm:"column:category_id;not null"`
	Title                 string    `json:"title" gorm:"column:title;not null"`
	Location              string    `json:"location" gorm:"column:location;not null"`
	IsUrgent              bool      `json:"is_urgent" gorm:"column:is_urgent;default:false"`
	WorkModel             string    `json:"work_model" gorm:"column:work_model;not null"`
	WorkSchedule          string    `json:"work_schedule" gorm:"column:work_schedule;not null"` // Ej: Tiempo completo, Medio Tiempo
	ContractType          string    `json:"contract_type" gorm:"column:contract_type;not null"` // Ej: Indefinido, Determinado
	ExperienceLevel       string    `json:"experience_level" gorm:"column:experience_level;not null"` // Ej: Sin experiencia, 1 año
	Salary                string    `json:"salary" gorm:"column:salary"` // Ej: $1000 - $1500, A convenir
	Description           string    `json:"description" gorm:"type:text;not null"`
	Responsibilities      string    `json:"responsibilities" gorm:"type:text;not null"`
	Requirements          string    `json:"requirements" gorm:"type:text;not null"`
	AccessibilityFeatures string    `json:"accessibility_features" gorm:"type:text"`

	// Relationships
	Employer Employer `json:"employer,omitempty"`
	Category Category `json:"category,omitempty"`
}
