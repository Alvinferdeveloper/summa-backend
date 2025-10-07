package services

import (
	"fmt"

	"github.com/Alvinferdeveloper/summa-backend/config"
	"github.com/Alvinferdeveloper/summa-backend/dto"
	"github.com/Alvinferdeveloper/summa-backend/models"
)

func GetJobApplicants(jobID uint) ([]dto.JobApplicationResponse, error) {
	var applications []models.JobApplication
	if err := config.DB.Preload("Profile").Preload("JobPost").Where("job_post_id = ?", jobID).Order("created_at desc").Find(&applications).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch job applications: %w", err)
	}

	var applicationDTOs []dto.JobApplicationResponse
	for _, app := range applications {
		applicationDTOs = append(applicationDTOs, dto.ConvertJobApplicationToDTO(app))
	}

	return applicationDTOs, nil
}
