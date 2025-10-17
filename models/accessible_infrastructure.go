package models

import "gorm.io/gorm"

// AccessibleInfrastructure represents a single accessibility feature.
type AccessibleInfrastructure struct {
	gorm.Model
	Name string `json:"name" gorm:"unique;not null"`
}
