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

func GetJobPosts(page, limit int, userID *uint, filters map[string]string) ([]dto.JobPostResponse, int64, bool, error) {
	offset := (page - 1) * limit

	var jobPosts []models.JobPost
	var total int64

	db := config.DB.Model(&models.JobPost{}).Where("status = ?", "open")

	for key, value := range filters {
		if value == "" {
			continue
		}
		switch key {
		case "is_urgent":
			db = db.Where("is_urgent = ?", value == "true")
		case "date_posted":
			db = db.Where("created_at >= ?", value)
		case "category_id":
			db = db.Where("category_id = ?", value)
		case "work_schedule":
			db = db.Where("work_schedule = ?", value)
		case "experience_level":
			db = db.Where("experience_level = ?", value)
		case "contract_type":
			db = db.Where("contract_type = ?", value)
		}
	}

	db.Count(&total)

	result := db.Preload("Employer").Preload("Category").Limit(limit).Offset(offset).Order("created_at desc").Find(&jobPosts)

	if result.Error != nil {
		return nil, 0, false, fmt.Errorf("failed to fetch job posts: %w", result.Error)
	}

	var profileID uint
	if userID != nil {
		var profile models.Profile
		config.DB.Where("user_id = ?", *userID).First(&profile)
		profileID = profile.ID
	}

	var jobPostDTOs []dto.JobPostResponse
	for _, jobPost := range jobPosts {
		dto := dto.ConvertJobPostToDTO(jobPost)
		if profileID != 0 {
			var count int64
			config.DB.Model(&models.JobApplication{}).Where("profile_id = ? AND job_post_id = ?", profileID, jobPost.ID).Count(&count)
			dto.HasApplied = count > 0
		}
		jobPostDTOs = append(jobPostDTOs, dto)
	}

	return jobPostDTOs, total, page*limit < int(total), nil
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
	if err := config.DB.Where("employer_id = ?", employerID).Order("created_at desc").Find(&jobPosts).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch employer job posts: %w", err)
	}

	var jobPostDTOs []dto.JobPostResponse
	for _, jobPost := range jobPosts {
		jobPostDTOs = append(jobPostDTOs, dto.ConvertJobPostToDTO(jobPost))
	}

	return jobPostDTOs, nil
}

func UpdateJobPostStatus(jobID uint, employerID uint, status string) (*models.JobPost, error) {
	var jobPost models.JobPost
	if err := config.DB.First(&jobPost, jobID).Error; err != nil {
		return nil, fmt.Errorf("job post not found")
	}

	// Security check: ensure the employer owns the job post
	if jobPost.EmployerID != employerID {
		return nil, fmt.Errorf("unauthorized to modify this job post")
	}

	jobPost.Status = status
	if err := config.DB.Save(&jobPost).Error; err != nil {
		return nil, fmt.Errorf("could not update job post status: %w", err)
	}

	return &jobPost, nil
}
