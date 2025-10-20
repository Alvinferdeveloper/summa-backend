package models

import "gorm.io/gorm"

// WorkModel represents a work model option (e.g., On-site, Hybrid, Remote).
type WorkModel struct {
	gorm.Model
	Name string `json:"name" gorm:"unique;not null"`
}
