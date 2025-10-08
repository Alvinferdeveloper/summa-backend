package dto

import (
	"time"

	"github.com/Alvinferdeveloper/summa-backend/models"
)

type CreateJobPostRequest struct {
	Title                 string   `json:"title" binding:"required"`
	Location              string   `json:"location" binding:"required"`
	WorkModel             string   `json:"workModel" binding:"required"`
	ContractType          string   `json:"contractType" binding:"required"`
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
	WorkModel             string               `json:"work_model"`
	ContractType          string               `json:"contract_type"`
	Description           string               `json:"description"`
	Responsibilities      string               `json:"responsibilities"`
	Requirements          string               `json:"requirements"`
	AccessibilityFeatures string               `json:"accessibility_features"`
	Employer              *EmployerResponseDTO `json:"employer"`
}

// ConvertJobPostToDTO converts a JobPost model to its DTO response.
func ConvertJobPostToDTO(jobPost models.JobPost) JobPostResponse {
	return JobPostResponse{
		ID:                    jobPost.ID,
		CreatedAt:             jobPost.CreatedAt,
		UpdatedAt:             jobPost.UpdatedAt,
		Title:                 jobPost.Title,
		Location:              jobPost.Location,
		WorkModel:             jobPost.WorkModel,
		ContractType:          jobPost.ContractType,
		Description:           jobPost.Description,
		Responsibilities:      jobPost.Responsibilities,
		Requirements:          jobPost.Requirements,
		AccessibilityFeatures: jobPost.AccessibilityFeatures,
		Employer:              ConvertEmployerToDTO(jobPost.Employer),
	}
}
