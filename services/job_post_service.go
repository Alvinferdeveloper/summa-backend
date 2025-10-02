package services

import (
	"encoding/json"

	"github.com/Alvinferdeveloper/summa-backend/config"
	"github.com/Alvinferdeveloper/summa-backend/dto"
	"github.com/Alvinferdeveloper/summa-backend/models"
)

func CreateJobPost(req *dto.CreateJobPostRequest, employerID uint) (*models.JobPost, error) {
	accessibilityFeaturesJSON, err := json.Marshal(req.AccessibilityFeatures)
	if err != nil {
		return nil, err
	}

	jobPost := &models.JobPost{
		EmployerID:            employerID,
		Title:                 req.Title,
		Location:              req.Location,
		WorkModel:             req.WorkModel,
		ContractType:          req.ContractType,
		Description:           req.Description,
		Responsibilities:      req.Responsibilities,
		Requirements:          req.Requirements,
		AccessibilityFeatures: string(accessibilityFeaturesJSON),
	}

	if err := config.DB.Create(jobPost).Error; err != nil {
		return nil, err
	}

	return jobPost, nil
}

func ListJobPosts(page, limit int) ([]models.JobPost, int64, error) {
	offset := (page - 1) * limit

	var jobPosts []models.JobPost
	var total int64

	// Get total count
	if err := config.DB.Model(&models.JobPost{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Get paginated results with employer preloading
	if err := config.DB.Preload("Employer").Limit(limit).Offset(offset).Order("created_at desc").Find(&jobPosts).Error; err != nil {
		return nil, 0, err
	}

	return jobPosts, total, nil
}
