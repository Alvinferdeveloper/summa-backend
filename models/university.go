package models

import "gorm.io/gorm"

// University represents an educational institution.
type University struct {
	gorm.Model
	Name    string `json:"name" gorm:"column:name;unique;not null"`
	Country string `json:"country" gorm:"column:country;not null"`
	Website string `json:"website" gorm:"column:website"`
	Address string `json:"address" gorm:"column:address"`
	LogoURL string `json:"logo_url" gorm:"column:logo_url"`

	// Relationships
	// A university can have many education records from candidates.
	Educations []ProfileEducation `json:"-" gorm:"foreignKey:UniversityID"`
}
