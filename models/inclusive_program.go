package models

import "gorm.io/gorm"

// InclusiveProgram represents a single inclusion program or benefit.
type InclusiveProgram struct {
	gorm.Model
	Name string `json:"name" gorm:"unique;not null"`
}
