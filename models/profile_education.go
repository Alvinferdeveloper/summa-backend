package models

import (
	"time"

	"gorm.io/gorm"
)

// ProfileEducation represents a profile's education history.
// It can be linked to an official University or a UniversitySuggestion
// if the university is not yet in the system.
type ProfileEducation struct {
	gorm.Model
	ProfileID    uint       `json:"profile_id" gorm:"column:profile_id;not null"`
	Degree       string     `json:"degree" gorm:"column:degree;not null"`
	FieldOfStudy string     `json:"field_of_study" gorm:"column:field_of_study"`
	StartDate    time.Time  `json:"start_date" gorm:"column:start_date"`
	EndDate      *time.Time `json:"end_date" gorm:"column:end_date"`

	// An education record is linked to either an existing university or a suggested one.
	UniversityID           *uint `json:"university_id,omitempty" gorm:"column:university_id"`
	UniversitySuggestionID *uint `json:"university_suggestion_id,omitempty" gorm:"column:university_suggestion_id"`

	University           *University           `json:"university,omitempty"`
	UniversitySuggestion *UniversitySuggestion `json:"university_suggestion,omitempty"`
}
