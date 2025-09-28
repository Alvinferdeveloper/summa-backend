package models

import "gorm.io/gorm"

type Skill struct {
	gorm.Model
	Name     string    `json:"name" gorm:"column:name;unique;not null"`
	Profiles []Profile `json:"-" gorm:"many2many:profile_skills;"`
}
