package models

import "gorm.io/gorm"

// WorkSchedule represents a work schedule option (e.g., Full-time, Part-time).
type WorkSchedule struct {
	gorm.Model
	Name string `json:"name" gorm:"unique;not null"`
}
