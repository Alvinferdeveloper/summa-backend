package models

import "gorm.io/gorm"

type Profile struct {
	gorm.Model
	UserID      uint   `json:"user_id" gorm:"column:user_id;unique;not null"`
	FirstName   string `json:"first_name" gorm:"column:first_name"`
	LastName    string `json:"last_name" gorm:"column:last_name"`
	PhoneNumber    string `json:"phone_number" gorm:"column:phone_number"`
	City           string `json:"city" gorm:"column:city"`
	Country        string `json:"country" gorm:"column:country"`
	ProfilePicture string `json:"profile_picture" gorm:"column:profile_picture"`
	Address        string `json:"address" gorm:"column:address"`
	LinkedIn       string `json:"linked_in" gorm:"column:linked_in"`
	ResumeURL      string `json:"resume_url" gorm:"column:resume_url"`
	Description    string `json:"description" gorm:"column:description"`

	// Disability and Accessibility Fields
	DisabilityInfoConsent  bool   `json:"disability_info_consent" gorm:"column:disability_info_consent;default:false"`
	DetailedAccommodations string `json:"detailed_accommodations,omitempty" gorm:"type:text"`

	// Relationships
	DisabilityTypes    []DisabilityType   `json:"disability_types,omitempty" gorm:"many2many:profile_disability_types;"`
	AccessibilityNeeds []AccessibilityNeed `json:"accessibility_needs,omitempty" gorm:"many2many:profile_accessibility_needs;"`
	Skills             []Skill            `json:"skills,omitempty" gorm:"many2many:profile_skills;"`
	Experiences        []Experience       `json:"experiences,omitempty" gorm:"foreignKey:ProfileID"`
	Educations         []ProfileEducation `json:"educations,omitempty" gorm:"foreignKey:ProfileID"`
}
