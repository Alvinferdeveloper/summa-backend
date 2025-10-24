package dto

import (
	"time"

	"github.com/Alvinferdeveloper/summa-backend/models"
	"github.com/google/uuid"
)

type CreateExperienceRequest struct {
	JobTitle      string     `json:"job_title" binding:"required"`
	Description   string     `json:"description"`
	StartDate     time.Time  `json:"start_date" binding:"required"`
	EndDate       *time.Time `json:"end_date"`
	EmployerID    *uuid.UUID `json:"employer_id"`
	NewEmployerID *uint      `json:"new_employer_id"`
}

type UpdateExperienceRequest struct {
	JobTitle      string     `json:"job_title" binding:"required"`
	Description   string     `json:"description"`
	StartDate     time.Time  `json:"start_date" binding:"required"`
	EndDate       *time.Time `json:"end_date"`
	EmployerID    *uuid.UUID `json:"employer_id"`
	NewEmployerID *uint      `json:"new_employer_id"`
}

type ExperienceResponseDTO struct {
	ID          uint                    `json:"id"`
	JobTitle    string                  `json:"job_title"`
	Description string                  `json:"description"`
	StartDate   time.Time               `json:"start_date"`
	EndDate     *time.Time              `json:"end_date"`
	Employer    *EmployerResponseDTO    `json:"employer,omitempty"`
	NewEmployer *NewEmployerResponseDTO `json:"new_employer,omitempty"`
}

func ConvertExperienceToDTO(exp models.Experience) ExperienceResponseDTO {
	dto := ExperienceResponseDTO{
		ID:          exp.ID,
		JobTitle:    exp.JobTitle,
		Description: exp.Description,
		StartDate:   exp.StartDate,
		EndDate:     exp.EndDate,
	}
	if exp.Employer != nil {
		dto.Employer = ConvertEmployerToDTO(*exp.Employer)
	}
	if exp.NewEmployer != nil {
		dto.NewEmployer = ConverNewEmployerToDTO(*exp.NewEmployer)
	}
	return dto
}
