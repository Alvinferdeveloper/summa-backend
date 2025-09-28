package models

import (
	"time"

	"gorm.io/gorm"
)

// ProfileSkill is the join table for the many-to-many relationship
// between Profile and Skill.
type ProfileSkill struct {
	ProfileID uint           `json:"profile_id" gorm:"primaryKey;column:profile_id"`
	SkillID   uint           `json:"skill_id" gorm:"primaryKey;column:skill_id"`
	CreatedAt time.Time      `json:"created_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
