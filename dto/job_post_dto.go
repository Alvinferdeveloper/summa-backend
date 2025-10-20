package dto

import (
	"time"

	"github.com/Alvinferdeveloper/summa-backend/models"
)

type CreateJobPostRequest struct {
	Title                string `json:"title" binding:"required"`
	Location             string `json:"location" binding:"required"`
	IsUrgent             bool   `json:"is_urgent"`
	WorkModelID          uint   `json:"workModelId" binding:"required"`
	WorkScheduleID       uint   `json:"workScheduleId" binding:"required"`
	ContractTypeID       uint   `json:"contractTypeId" binding:"required"`
	ExperienceLevelID    uint   `json:"experienceLevelId" binding:"required"`
	Salary               string `json:"salary"`
	CategoryID           uint   `json:"category_id" binding:"required"`
	Description          string `json:"description" binding:"required"`
	Responsibilities     string `json:"responsibilities" binding:"required"`
	Requirements         string `json:"requirements" binding:"required"`
	AccessibilityNeedIDs []uint `json:"accessibilityNeedIds"`
	DisabilityTypeIDs    []uint `json:"disabilityTypeIds"`
}

type JobPostResponse struct {
	ID                 uint                        `json:"id"`
	CreatedAt          time.Time                   `json:"created_at"`
	UpdatedAt          time.Time                   `json:"updated_at"`
	Title              string                      `json:"title"`
	Location           string                      `json:"location"`
	IsUrgent           bool                        `json:"is_urgent"`
	WorkModel          string                      `json:"work_model"`
	WorkSchedule       string                      `json:"work_schedule"`
	ContractType       string                      `json:"contract_type"`
	ExperienceLevel    string                      `json:"experience_level"`
	Salary             string                      `json:"salary"`
	Description        string                      `json:"description"`
	Status             string                      `json:"status"`
	Responsibilities   string                      `json:"responsibilities"`
	Requirements       string                      `json:"requirements"`
	AccessibilityNeeds []AccessibilityNeedResponse `json:"accessibility_needs,omitempty"`
	DisabilityTypes    []DisabilityTypeResponse    `json:"disability_types,omitempty"`
	Employer           *EmployerResponseDTO        `json:"employer"`
	Category           *CategoryResponseDTO        `json:"category"`
	ApplicantCount     int64                       `json:"applicant_count"`
	HasApplied         bool                        `json:"has_applied"`
}

type UpdateJobPostStatusRequest struct {
	Status string `json:"status" binding:"required,oneof=open closed"`
}

// ConvertJobPostToDTO converts a JobPost model to its DTO response.
func ConvertJobPostToDTO(jobPost models.JobPost) JobPostResponse {
	return JobPostResponse{
		ID:              jobPost.ID,
		CreatedAt:       jobPost.CreatedAt,
		UpdatedAt:       jobPost.UpdatedAt,
		Title:           jobPost.Title,
		Location:        jobPost.Location,
		IsUrgent:        jobPost.IsUrgent,
		WorkModel:       jobPost.WorkModel.Name,
		WorkSchedule:    jobPost.WorkSchedule.Name,
		ContractType:    jobPost.ContractType.Name,
		ExperienceLevel: jobPost.ExperienceLevel.Name,
		AccessibilityNeeds: func() []AccessibilityNeedResponse {
			var dtos []AccessibilityNeedResponse
			for _, an := range jobPost.AccessibilityNeeds {
				dtos = append(dtos, ConvertAccessibilityNeedToDTO(an))
			}
			return dtos
		}(),
		DisabilityTypes: func() []DisabilityTypeResponse {
			var dtos []DisabilityTypeResponse
			for _, dt := range jobPost.DisabilityTypes {
				dtos = append(dtos, ConvertDisabilityTypeToDTO(dt))
			}
			return dtos
		}(),
		Salary:           jobPost.Salary,
		Status:           jobPost.Status,
		Description:      jobPost.Description,
		Responsibilities: jobPost.Responsibilities,
		Requirements:     jobPost.Requirements,
		Employer:         ConvertEmployerToDTO(jobPost.Employer),
		Category:         ConvertCategoryToDTO(jobPost.Category),
	}
}
