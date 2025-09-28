package models

import (
	"time"

	"gorm.io/gorm"
)

// Experience represents a profile's work experience.
// It can be linked to an official Employer or a NewEmployer suggestion
// if the company is not yet in the system.
type Experience struct {
	gorm.Model
	ProfileID   uint       `json:"profile_id" gorm:"column:profile_id;not null"`
	JobTitle    string     `json:"job_title" gorm:"column:job_title;not null"`
	Description string     `json:"description" gorm:"column:description"`
	StartDate   time.Time  `json:"start_date" gorm:"column:start_date;not null"`
	EndDate     *time.Time `json:"end_date" gorm:"column:end_date"`

	// An experience record is linked to either an existing employer or a suggested one.
	EmployerID    *uint `json:"employer_id,omitempty" gorm:"column:employer_id"`
	NewEmployerID *uint `json:"new_employer_id,omitempty" gorm:"column:new_employer_id"`

	// Relationships
	Employer    *Employer    `json:"employer,omitempty"`
	NewEmployer *NewEmployer `json:"new_employer,omitempty"`
}
