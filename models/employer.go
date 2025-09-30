package models

import (
	"time"

	"gorm.io/gorm"
)

// Employer represents a company that can post jobs and authenticates with email and password.
type Employer struct {
	gorm.Model
	CompanyName                   string     `json:"company_name" gorm:"column:company_name;unique;not null"`
	Email                         string     `json:"email" gorm:"column:email;unique;not null"`
	Password                      string     `json:"-" gorm:"column:password;not null"`
	Role                          string     `json:"role" gorm:"column:role;not null;default:'employer'"`
	PhoneNumber                   string     `json:"phone_number" gorm:"column:phone_number"`
	Country                       string     `json:"country" gorm:"column:country"`
	FoundationDate                *time.Time `json:"foundation_date" gorm:"column:foundation_date"` // Use pointer for nullable date
	Industry                      string     `json:"industry" gorm:"column:industry"`
	Size                          string     `json:"size" gorm:"column:size"` // e.g., "1-10", "11-50", "51-200", "201-500", "500+"
	Description                   string     `json:"description" gorm:"type:text"`
	Address                       string     `json:"address" gorm:"column:address"`
	Website                       string     `json:"website" gorm:"column:website"`
	LogoURL                       string     `json:"logo_url" gorm:"column:logo_url"`
	DiversityInclusionPolicyURL   string     `json:"diversity_inclusion_policy_url" gorm:"column:diversity_inclusion_policy_url"`
	InclusionStatement            string     `json:"inclusion_statement" gorm:"type:text"`
	AccessibilityCertificationURL string     `json:"accessibility_certification_url" gorm:"column:accessibility_certification_url"`

	// Relationships
	JobPositions []JobPosition `json:"job_positions,omitempty" gorm:"foreignKey:EmployerID"`
	Experiences  []Experience  `json:"-" gorm:"foreignKey:EmployerID"`
}
