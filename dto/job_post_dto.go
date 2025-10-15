package dto

import (
	"time"

	"github.com/Alvinferdeveloper/summa-backend/models"
)

type CreateJobPostRequest struct {
	Title                 string   `json:"title" binding:"required"`
	Location              string   `json:"location" binding:"required"`
	IsUrgent              bool     `json:"is_urgent"`
	WorkModel             string   `json:"workModel" binding:"required"`
	WorkSchedule          string   `json:"workSchedule" binding:"required"`
	ContractType          string   `json:"contractType" binding:"required"`
	ExperienceLevel       string   `json:"experienceLevel" binding:"required"`
	Salary                string   `json:"salary"`
	CategoryID            uint     `json:"category_id" binding:"required"`
	Description           string   `json:"description" binding:"required"`
	Responsibilities      string   `json:"responsibilities" binding:"required"`
	Requirements          string   `json:"requirements" binding:"required"`
	AccessibilityFeatures []string `json:"accessibilityFeatures"`
}

type JobPostResponse struct {
	ID                    uint                 `json:"id"`
	CreatedAt             time.Time            `json:"created_at"`
	UpdatedAt             time.Time            `json:"updated_at"`
	Title                 string               `json:"title"`
	Location              string               `json:"location"`
	IsUrgent              bool                 `json:"is_urgent"`
	WorkModel             string               `json:"work_model"`
	WorkSchedule          string               `json:"work_schedule"`
	ContractType          string               `json:"contract_type"`
	ExperienceLevel       string               `json:"experience_level"`
	Salary                string               `json:"salary"`
	Description           string               `json:"description"`
	Responsibilities      string               `json:"responsibilities"`
	Requirements          string               `json:"requirements"`
	AccessibilityFeatures string               `json:"accessibility_features"`
	Employer              *EmployerResponseDTO `json:"employer"`
	Category              *CategoryResponseDTO `json:"category"`
	ApplicantCount        int64                `json:"applicant_count"`
	HasApplied            bool                 `json:"has_applied"`
}

// ConvertJobPostToDTO converts a JobPost model to its DTO response.
func ConvertJobPostToDTO(jobPost models.JobPost) JobPostResponse {
	return JobPostResponse{
		ID:                    jobPost.ID,
		CreatedAt:             jobPost.CreatedAt,
		UpdatedAt:             jobPost.UpdatedAt,
		Title:                 jobPost.Title,
		Location:              jobPost.Location,
		IsUrgent:              jobPost.IsUrgent,
		WorkModel:             jobPost.WorkModel,
		WorkSchedule:          jobPost.WorkSchedule,
		ContractType:          jobPost.ContractType,
		ExperienceLevel:       jobPost.ExperienceLevel,
		Salary:                jobPost.Salary,
		Description:           jobPost.Description,
		Responsibilities:      jobPost.Responsibilities,
		Requirements:          jobPost.Requirements,
		AccessibilityFeatures: jobPost.AccessibilityFeatures,
		Employer:              ConvertEmployerToDTO(jobPost.Employer),
		Category:              ConvertCategoryToDTO(jobPost.Category),
	}
}
