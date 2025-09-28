package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName string `json:"first_name" gorm:"column:first_name;not null"`
	LastName  string `json:"last_name" gorm:"column:last_name;not null"`
	Email     string `json:"email" gorm:"column:email;unique;not null"`
	Password  string `json:"-" gorm:"column:password;not null"`
	Role      string `json:"role" gorm:"column:role;not null;default:'job_seeker'"`

	// A user can have one job seeker profile
	Profile Profile `json:"profile,omitempty" gorm:"foreignKey:UserID"`

	// A user can have one employer profile
	Employer Employer `json:"employer,omitempty" gorm:"foreignKey:UserID"`
}
