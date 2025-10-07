package services

import (
	"encoding/json"
	"fmt"

	"github.com/Alvinferdeveloper/summa-backend/config"
	"github.com/Alvinferdeveloper/summa-backend/dto"
	"github.com/Alvinferdeveloper/summa-backend/models"
	"gorm.io/gorm"
)

func CreateJobPost(req *dto.CreateJobPostRequest, employerID uint) (*models.JobPost, error) {
	accessibilityFeaturesJSON, err := json.Marshal(req.AccessibilityFeatures)
	if err != nil {
		return nil, fmt.Errorf("failed to process accessibility features: %w", err)
	}

	jobPost := models.JobPost{
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

	if err := config.DB.Create(&jobPost).Error; err != nil {
		return nil, fmt.Errorf("failed to create job post: %w", err)
	}

	return &jobPost, nil
}

func GetJobPosts(page, limit int) ([]dto.JobPostResponse, int64, bool, error) {
	offset := (page - 1) * limit

	var jobPosts []models.JobPost
	var total int64

	config.DB.Model(&models.JobPost{}).Count(&total)

	result := config.DB.Preload("Employer").Limit(limit).Offset(offset).Order("created_at desc").Find(&jobPosts)

	if result.Error != nil {
		return nil, 0, false, fmt.Errorf("failed to fetch job posts: %w", result.Error)
	}

	var jobPostDTOs []dto.JobPostResponse
	for _, jobPost := range jobPosts {
		jobPostDTOs = append(jobPostDTOs, dto.ConvertJobPostToDTO(jobPost))
	}

	return jobPostDTOs, total, page*limit < int(total), nil
}

func GetJobPostById(id uint) (*dto.JobPostResponse, error) {
	var jobPost models.JobPost
	if err := config.DB.Preload("Employer").First(&jobPost, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("job post not found")
		}
		return nil, fmt.Errorf("failed to fetch job post: %w", err)
	}

	jobPostDTO := dto.ConvertJobPostToDTO(jobPost)
	return &jobPostDTO, nil
}

func GetJobPostsByEmployerID(employerID uint) ([]dto.JobPostResponse, error) {
	var jobPosts []models.JobPost
	if err := config.DB.Preload("Employer").Where("employer_id = ?", employerID).Order("created_at desc").Find(&jobPosts).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch employer job posts: %w", err)
	}

	var jobPostDTOs []dto.JobPostResponse
	for _, jobPost := range jobPosts {
		jobPostDTOs = append(jobPostDTOs, dto.ConvertJobPostToDTO(jobPost))
	}

	return jobPostDTOs, nil
}
