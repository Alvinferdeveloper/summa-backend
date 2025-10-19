package models

import "gorm.io/gorm"

type JobApplication struct {
	gorm.Model
	ProfileID              uint   `json:"profile_id" gorm:"column:profile_id;not null"`
	JobPostID              uint   `json:"job_post_id" gorm:"column:job_post_id;not null"`
	Status                 string `json:"status" gorm:"column:status;not null;default:'Postulado'"`
	CoverLetter            string `json:"cover_letter" gorm:"type:text"`
	ResumeURLAtApplication string `json:"resume_url_at_application" gorm:"column:resume_url_at_application"`

	// Relationships
	Profile Profile `json:"profile,omitempty"`
	JobPost JobPost `json:"job_post,omitempty"`
}
