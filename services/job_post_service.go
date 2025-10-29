package services

import (
	"fmt"

	"github.com/Alvinferdeveloper/summa-backend/config"
	"github.com/Alvinferdeveloper/summa-backend/dto"
	"github.com/Alvinferdeveloper/summa-backend/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateJobPost(req *dto.CreateJobPostRequest, employerID uuid.UUID) (*models.JobPost, error) {
	jobPost := models.JobPost{
		EmployerID:        employerID,
		CategoryID:        req.CategoryID,
		Title:             req.Title,
		Location:          req.Location,
		IsUrgent:          req.IsUrgent,
		WorkModelID:       req.WorkModelID,
		WorkScheduleID:    req.WorkScheduleID,
		ContractTypeID:    req.ContractTypeID,
		ExperienceLevelID: req.ExperienceLevelID,
		Salary:            req.Salary,
		Description:       req.Description,
		Responsibilities:  req.Responsibilities,
		Requirements:      req.Requirements,
	}

	err := config.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&jobPost).Error; err != nil {
			return fmt.Errorf("failed to create job post: %w", err)
		}
		if len(req.AccessibilityNeedIDs) > 0 {
			var accessibilityNeeds []models.AccessibilityNeed
			if err := tx.Where("id IN ?", req.AccessibilityNeedIDs).Find(&accessibilityNeeds).Error; err != nil {
				return fmt.Errorf("failed to find accessibility needs: %w", err)
			}
			if err := tx.Model(&jobPost).Association("AccessibilityNeeds").Append(&accessibilityNeeds); err != nil {
				return fmt.Errorf("failed to associate accessibility needs: %w", err)
			}
		}

		if len(req.DisabilityTypeIDs) > 0 {
			var disabilityTypes []models.DisabilityType
			if err := tx.Where("id IN ?", req.DisabilityTypeIDs).Find(&disabilityTypes).Error; err != nil {
				return fmt.Errorf("failed to find disability types: %w", err)
			}
			if err := tx.Model(&jobPost).Association("DisabilityTypes").Append(&disabilityTypes); err != nil {
				return fmt.Errorf("failed to associate disability types: %w", err)
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &jobPost, nil
}

func GetJobPosts(page, limit int, userID *uuid.UUID, filters map[string]string) ([]dto.JobPostResponse, int64, bool, error) {
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
			baseQuery = baseQuery.Where("job_posts.is_urgent = ?", value == "true")
		case "date_posted":
			baseQuery = baseQuery.Where("job_posts.created_at >= ?", value)
		case "category_id":
			baseQuery = baseQuery.Where("categories.id = ?", value)
		case "work_schedule_id":
			baseQuery = baseQuery.Where("work_schedules.id = ?", value)
		case "experience_level_id":
			baseQuery = baseQuery.Where("experience_levels.id = ?", value)
		case "contract_type_id":
			baseQuery = baseQuery.Where("contract_types.id = ?", value)
		case "work_model_id":
			baseQuery = baseQuery.Where("work_models.id = ?", value)
		case "disability_type_id":
			baseQuery = baseQuery.Joins("JOIN job_post_disability_types ON job_post_disability_types.job_post_id = job_posts.id").Where("job_post_disability_types.disability_type_id = ?", value)
		}
	}

	// Create a separate query for counting distinct job posts
	countQuery := baseQuery.Session(&gorm.Session{}).Distinct("job_posts.id")
	if err := countQuery.Count(&total).Error; err != nil {
		return nil, 0, false, fmt.Errorf("failed to count job posts: %w", err)
	}

	// Perform the find query with pagination and preloads
	result := baseQuery.Preload("Employer").Preload("Category").Preload("ContractType").Preload("ExperienceLevel").Preload("WorkSchedule").Preload("WorkModel").Preload("AccessibilityNeeds").Preload("DisabilityTypes").Limit(limit).Offset(offset).Order("job_posts.created_at desc").Find(&jobPosts)

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
	if err := config.DB.Preload("Employer").Preload("Category").Preload("ContractType").Preload("ExperienceLevel").Preload("WorkSchedule").Preload("WorkModel").Preload("AccessibilityNeeds").Preload("DisabilityTypes").First(&jobPost, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("job post not found")
		}
		return nil, fmt.Errorf("failed to fetch job post: %w", err)
	}

	return &jobPost, nil
}

func GetJobPostsByEmployerID(employerID uuid.UUID) ([]dto.JobPostResponse, error) {
	var jobPosts []models.JobPost
	if err := config.DB.Where("employer_id = ?", employerID).Preload("WorkModel").Order("created_at desc").Find(&jobPosts).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch employer job posts: %w", err)
	}

	var jobPostDTOs []dto.JobPostResponse
	for _, jobPost := range jobPosts {
		jobPostDTOs = append(jobPostDTOs, dto.ConvertJobPostToDTO(jobPost))
	}

	return jobPostDTOs, nil
}

func UpdateJobPostStatus(jobID uint, employerID uuid.UUID, status string) (*models.JobPost, error) {
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

func UpdateJobPost(jobID uint, employerID uuid.UUID, req *dto.UpdateJobPostRequest) (*models.JobPost, error) {
	var jobPost models.JobPost

	err := config.DB.Transaction(func(tx *gorm.DB) error {
		// 1. Find the job post and verify ownership
		if err := tx.First(&jobPost, jobID).Error; err != nil {
			return fmt.Errorf("job post not found")
		}
		if jobPost.EmployerID != employerID {
			return fmt.Errorf("unauthorized to modify this job post")
		}

		// 2. Update primary fields
		jobPost.Title = req.Title
		jobPost.Location = req.Location
		jobPost.IsUrgent = req.IsUrgent
		jobPost.WorkModelID = req.WorkModelID
		jobPost.WorkScheduleID = req.WorkScheduleID
		jobPost.ContractTypeID = req.ContractTypeID
		jobPost.ExperienceLevelID = req.ExperienceLevelID
		jobPost.Salary = req.Salary
		jobPost.CategoryID = req.CategoryID
		jobPost.Description = req.Description
		jobPost.Responsibilities = req.Responsibilities
		jobPost.Requirements = req.Requirements

		// 3. Update Accessibility Needs (replace existing associations)
		var accessibilityNeeds []models.AccessibilityNeed
		if len(req.AccessibilityNeedIDs) > 0 {
			if err := tx.Where("id IN ?", req.AccessibilityNeedIDs).Find(&accessibilityNeeds).Error; err != nil {
				return err
			}
		}
		if err := tx.Model(&jobPost).Association("AccessibilityNeeds").Replace(&accessibilityNeeds); err != nil {
			return err
		}

		// 4. Update Disability Types (replace existing associations)
		var disabilityTypes []models.DisabilityType
		if len(req.DisabilityTypeIDs) > 0 {
			if err := tx.Where("id IN ?", req.DisabilityTypeIDs).Find(&disabilityTypes).Error; err != nil {
				return err
			}
		}
		if err := tx.Model(&jobPost).Association("DisabilityTypes").Replace(&disabilityTypes); err != nil {
			return err
		}

		// 5. Save the updated job post
		if err := tx.Save(&jobPost).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &jobPost, nil
}

func CalculateCompatibility(profileID uint, jobPostID uint) (*dto.CompatibilityResponse, error) {
	var profile models.Profile
	var jobPost models.JobPost

	if err := config.DB.Preload("AccessibilityNeeds").First(&profile, profileID).Error; err != nil {
		return nil, fmt.Errorf("profile not found")
	}
	if err := config.DB.Preload("AccessibilityNeeds").First(&jobPost, jobPostID).Error; err != nil {
		return nil, fmt.Errorf("job post not found")
	}

	profileNeeds := make(map[uint]string)
	for _, need := range profile.AccessibilityNeeds {
		profileNeeds[need.ID] = need.Name
	}

	jobFeatures := make(map[uint]string)
	for _, feature := range jobPost.AccessibilityNeeds {
		jobFeatures[feature.ID] = feature.Name
	}

	var metNeeds []string = []string{}
	var unmetNeeds []string = []string{}
	metCount := 0

	for id, name := range profileNeeds {
		if _, found := jobFeatures[id]; found {
			metCount++
			metNeeds = append(metNeeds, name)
		} else {
			unmetNeeds = append(unmetNeeds, name)
		}
	}

	var score float64
	if len(profileNeeds) > 0 {
		score = (float64(metCount) / float64(len(profileNeeds))) * 100
	} else {
		score = 100
	}

	return &dto.CompatibilityResponse{
		Score:               score,
		MetNeeds:            metNeeds,
		UnmetNeeds:          unmetNeeds,
		TotalCandidateNeeds: len(profileNeeds),
	}, nil
}
