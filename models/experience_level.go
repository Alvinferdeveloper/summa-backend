package models

import "gorm.io/gorm"

type ExperienceLevel struct {
	gorm.Model
	Name string `json:"name" gorm:"unique;not null"`
}
