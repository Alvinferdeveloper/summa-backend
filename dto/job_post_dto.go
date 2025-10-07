package dto

import (
	"time"

	"github.com/Alvinferdeveloper/summa-backend/models"
)

// JobPostResponse represents the data structure of a job post for the API.
type JobPostResponse struct {
	ID                    uint      `json:"ID"`
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAt             time.Time `json:"updated_at"`
	Title                 string    `json:"title"`
	Location              string    `json:"location"`
	WorkModel             string    `json:"work_model"`
	ContractType          string    `json:"contract_type"`
	Description           string    `json:"description"`
	Responsibilities      string    `json:"responsibilities"`
	Requirements          string    `json:"requirements"`
	AccessibilityFeatures string    `json:"accessibility_features"`
	Employer              EmployerResponse `json:"employer"`
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
		Employer: EmployerResponse{
			ID:          jobPost.Employer.ID,
			CompanyName: jobPost.Employer.CompanyName,
			LogoURL:     jobPost.Employer.LogoURL,
			Industry:    jobPost.Employer.Industry,
			Email:       jobPost.Employer.Email,
		},
	}
}
