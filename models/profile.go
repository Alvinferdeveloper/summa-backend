package models

import "gorm.io/gorm"

type Profile struct {
	gorm.Model
	UserID         uint   `json:"user_id" gorm:"column:user_id;unique;not null"`
	PhoneNumber    string `json:"phone_number" gorm:"column:phone_number"`
	City           string `json:"city" gorm:"column:city"`
	Country        string `json:"country" gorm:"column:country"`
	ProfilePicture string `json:"profile_picture" gorm:"column:profile_picture"`
	Address        string `json:"address" gorm:"column:address"`
	ResumeURL      string `json:"resume_url" gorm:"column:resume_url"`

	// Relationships
	Skills      []Skill            `json:"skills,omitempty" gorm:"many2many:profile_skills;"`
	Experiences []Experience       `json:"experiences,omitempty" gorm:"foreignKey:ProfileID"`
	Educations  []ProfileEducation `json:"educations,omitempty" gorm:"foreignKey:ProfileID"`
}
