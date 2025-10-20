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
	WorkModelID           uint      `json:"work_model_id" gorm:"column:work_model_id;not null"`
	WorkScheduleID        uint      `json:"work_schedule_id" gorm:"column:work_schedule_id;not null"`
	ContractTypeID        uint      `json:"contract_type_id" gorm:"column:contract_type_id;not null"`
	ExperienceLevelID     uint      `json:"experience_level_id" gorm:"column:experience_level_id;not null"`
	Salary                string    `json:"salary" gorm:"column:salary"` // Ej: $1000 - $1500, A convenir
	Description           string    `json:"description" gorm:"type:text;not null"`
	Responsibilities      string    `json:"responsibilities" gorm:"type:text;not null"`
	Requirements          string    `json:"requirements" gorm:"type:text;not null"`
	AccessibilityFeatures string    `json:"accessibility_features" gorm:"type:text"`
	Status                string    `json:"status" gorm:"column:status;not null;default:'open'"` // e.g., open, closed

	// Relationships
	Employer        Employer        `json:"employer,omitempty"`
	Category        Category        `json:"category,omitempty"`
	ContractType    ContractType    `json:"contract_type,omitempty"`
	ExperienceLevel ExperienceLevel `json:"experience_level,omitempty"`
	WorkSchedule    WorkSchedule    `json:"work_schedule,omitempty"`
	WorkModel       WorkModel       `json:"work_model,omitempty"`
}
