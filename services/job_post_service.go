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
		WorkModelID:           req.WorkModelID,
		WorkScheduleID:        req.WorkScheduleID,
		ContractTypeID:        req.ContractTypeID,
		ExperienceLevelID:     req.ExperienceLevelID,
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

	// Base query with all joins and initial status filter
	baseQuery := config.DB.Model(&models.JobPost{}).
		Joins("JOIN employers ON employers.id = job_posts.employer_id").
		Joins("JOIN categories ON categories.id = job_posts.category_id").
		Joins("JOIN contract_types ON contract_types.id = job_posts.contract_type_id").
		Joins("JOIN experience_levels ON experience_levels.id = job_posts.experience_level_id").
		Joins("JOIN work_schedules ON work_schedules.id = job_posts.work_schedule_id").
		Joins("JOIN work_models ON work_models.id = job_posts.work_model_id").
		Where("job_posts.status = ?", "open")

	// Apply dynamic filters
	for key, value := range filters {
		if value == "" {
			continue
		}
		switch key {
		case "is_urgent":
			baseQuery = baseQuery.Where("is_urgent = ?", value == "true")
		case "date_posted":
			baseQuery = baseQuery.Where("created_at >= ?", value)
		case "category_id":
			baseQuery = baseQuery.Where("category_id = ?", value)
		case "work_schedule_id":
			baseQuery = baseQuery.Where("work_schedule_id = ?", value)
		case "experience_level_id":
			baseQuery = baseQuery.Where("experience_level_id = ?", value)
		case "contract_type_id":
			baseQuery = baseQuery.Where("contract_type_id = ?", value)
		case "work_model_id":
			baseQuery = baseQuery.Where("work_model_id = ?", value)
		}
	}

	if err := baseQuery.Count(&total).Error; err != nil {
		return nil, 0, false, fmt.Errorf("failed to count job posts: %w", err)
	}

	// Perform the find query with pagination and preloads
	result := baseQuery.Preload("Employer").Preload("Category").Preload("ContractType").Preload("ExperienceLevel").Preload("WorkSchedule").Preload("WorkModel").Limit(limit).Offset(offset).Order("job_posts.created_at desc").Find(&jobPosts)

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
	if err := config.DB.Preload("Employer").Preload("Category").Preload("ContractType").Preload("ExperienceLevel").Preload("WorkSchedule").Preload("WorkModel").First(&jobPost, id).Error; err != nil {
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
