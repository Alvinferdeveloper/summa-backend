package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// JobPost represents a job posting by an Employer.
type JobPost struct {
	gorm.Model
	EmployerID        uuid.UUID `json:"employer_id" gorm:"column:employer_id;not null"`
	CategoryID        uint      `json:"category_id" gorm:"column:category_id;not null"`
	Title             string    `json:"title" gorm:"column:title;not null"`
	Location          string    `json:"location" gorm:"column:location;not null"`
	Latitude          float64   `json:"latitude" gorm:"column:latitude"`
	Longitude         float64   `json:"longitude" gorm:"column:longitude"`
	IsUrgent          bool      `json:"is_urgent" gorm:"column:is_urgent;default:false"`
	WorkModelID       uint      `json:"work_model_id" gorm:"column:work_model_id;not null"`
	WorkScheduleID    uint      `json:"work_schedule_id" gorm:"column:work_schedule_id;not null"`
	ContractTypeID    uint      `json:"contract_type_id" gorm:"column:contract_type_id;not null"`
	ExperienceLevelID uint      `json:"experience_level_id" gorm:"column:experience_level_id;not null"`
	Salary            string    `json:"salary" gorm:"column:salary"` // Ej: $1000 - $1500, A convenir
	Description       string    `json:"description" gorm:"type:text;not null"`
	Responsibilities  string    `json:"responsibilities" gorm:"type:text;not null"`
	Requirements      string    `json:"requirements" gorm:"type:text;not null"`
	Status            string    `json:"status" gorm:"column:status;not null;default:'open'"` // e.g., open, closed

	// Relationships
	Employer           Employer            `json:"employer,omitempty"`
	Category           Category            `json:"category,omitempty"`
	ContractType       ContractType        `json:"contract_type,omitempty"`
	ExperienceLevel    ExperienceLevel     `json:"experience_level,omitempty"`
	WorkSchedule       WorkSchedule        `json:"work_schedule,omitempty"`
	WorkModel          WorkModel           `json:"work_model,omitempty"`
	AccessibilityNeeds []AccessibilityNeed `json:"accessibility_needs,omitempty" gorm:"many2many:job_post_accessibility_needs;"`
	DisabilityTypes    []DisabilityType    `json:"disability_types,omitempty" gorm:"many2many:job_post_disability_types;"`
}
