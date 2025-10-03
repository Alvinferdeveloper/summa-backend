package dto

import "time"

type NewEmployerRequest struct {
	CompanyName string `json:"company_name" binding:"required"`
	Website     string `json:"website"`
}

type CreateExperienceRequest struct {
	JobTitle      string     `json:"job_title" binding:"required"`
	Description   string     `json:"description"`
	StartDate     time.Time  `json:"start_date" binding:"required"`
	EndDate       *time.Time `json:"end_date"`
	EmployerID    *uint      `json:"employer_id"`
	NewEmployerID *uint      `json:"new_employer_id"`
}
