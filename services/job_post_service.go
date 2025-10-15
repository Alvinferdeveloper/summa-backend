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
		CategoryID:            req.CategoryID,
		Title:                 req.Title,
		Location:              req.Location,
		IsUrgent:              req.IsUrgent,
		WorkModel:             req.WorkModel,
		WorkSchedule:          req.WorkSchedule,
		ContractType:          req.ContractType,
		ExperienceLevel:       req.ExperienceLevel,
		Salary:                req.Salary,
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

func GetJobPosts(page, limit int) ([]models.JobPost, int64, bool, error) {
	offset := (page - 1) * limit

	var jobPosts []models.JobPost
	var total int64

	config.DB.Model(&models.JobPost{}).Count(&total)

	result := config.DB.Preload("Employer").Preload("Category").Limit(limit).Offset(offset).Order("created_at desc").Find(&jobPosts)

	if result.Error != nil {
		return nil, 0, false, fmt.Errorf("failed to fetch job posts: %w", result.Error)
	}

	return jobPosts, total, page*limit < int(total), nil
}

func GetJobPostById(id uint) (*models.JobPost, error) {
	var jobPost models.JobPost
	if err := config.DB.Preload("Employer").Preload("Category").First(&jobPost, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("job post not found")
		}
		return nil, fmt.Errorf("failed to fetch job post: %w", err)
	}

	return &jobPost, nil
}

func GetJobPostsByEmployerID(employerID uint) ([]dto.JobPostResponse, error) {
	var jobPosts []models.JobPost
	if err := config.DB.Preload("Employer").Preload("Category").Where("employer_id = ?", employerID).Order("created_at desc").Find(&jobPosts).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch employer job posts: %w", err)
	}

	var jobPostDTOs []dto.JobPostResponse
	for _, jobPost := range jobPosts {
		var applicantCount int64
		config.DB.Model(&models.JobApplication{}).Where("job_post_id = ?", jobPost.ID).Count(&applicantCount)
		dto := dto.ConvertJobPostToDTO(jobPost)
		dto.ApplicantCount = applicantCount
		jobPostDTOs = append(jobPostDTOs, dto)
	}

	return jobPostDTOs, nil
}
