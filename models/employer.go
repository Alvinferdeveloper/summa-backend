package models

import "gorm.io/gorm"

type Employer struct {
	gorm.Model
	UserID      uint   `json:"user_id" gorm:"column:user_id;unique;not null"`
	CompanyName string `json:"company_name" gorm:"column:company_name;unique;not null"`
	Website     string `json:"website" gorm:"column:website"`
	LogoURL     string `json:"logo_url" gorm:"column:logo_url"`

	// Relationships
	JobPositions []JobPosition `json:"job_positions,omitempty" gorm:"foreignKey:EmployerID"`
	// Experiences are linked to employers, but the primary relation is from the candidate side.
	// This field can be used to query all candidates that worked for an employer.
	Experiences []Experience `json:"-" gorm:"foreignKey:EmployerID"`
}
