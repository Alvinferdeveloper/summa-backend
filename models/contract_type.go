package models

import "gorm.io/gorm"

type ContractType struct {
	gorm.Model
	Name string `json:"name" gorm:"unique;not null"`
}
