package services

import (
	"fmt"

	"github.com/Alvinferdeveloper/summa-backend/config"
	"github.com/Alvinferdeveloper/summa-backend/models"
	"gorm.io/gorm"
)

func CreateJobApplication(profile *models.Profile, jobID uint, coverLetter string) (*models.JobApplication, error) {
	var existingApplication models.JobApplication
	err := config.DB.Where("profile_id = ? AND job_post_id = ?", profile.ID, jobID).First(&existingApplication).Error
	if err == nil {
		return nil, fmt.Errorf("already applied to this job")
	}
	if err != gorm.ErrRecordNotFound {
		return nil, fmt.Errorf("error checking for existing application: %w", err)
	}

	application := &models.JobApplication{
		ProfileID:              profile.ID,
		JobPostID:              jobID,
		Status:                 "Postulado",
		CoverLetter:            coverLetter,
		ResumeURLAtApplication: profile.ResumeURL,
	}

	if err := config.DB.Create(application).Error; err != nil {
		return nil, fmt.Errorf("could not create application: %w", err)
	}

	return application, nil
}

func GetMyApplications(profileID uint) ([]models.JobApplication, error) {
	var applications []models.JobApplication
	if err := config.DB.Preload("JobPost.Employer").Where("profile_id = ?", profileID).Order("created_at desc").Find(&applications).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch applications: %w", err)
	}
	return applications, nil
}

func GetJobApplicants(jobID uint, employerID uint) ([]models.JobApplication, error) {
	var jobPost models.JobPost
	if err := config.DB.Where("id = ? AND employer_id = ?", jobID, employerID).First(&jobPost).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("job post not found or does not belong to employer")
		}
		return nil, fmt.Errorf("error verifying job post: %w", err)
	}

	var applications []models.JobApplication
	if err := config.DB.Preload("Profile").Where("job_post_id = ?", jobID).Order("created_at desc").Find(&applications).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch job applications: %w", err)
	}
	return applications, nil
}
