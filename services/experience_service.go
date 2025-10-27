package services

import (
	"github.com/Alvinferdeveloper/summa-backend/config"
	"github.com/Alvinferdeveloper/summa-backend/dto"
	"github.com/Alvinferdeveloper/summa-backend/models"
	"github.com/google/uuid"
)

func CreateExperience(profileID uint, req *dto.CreateExperienceRequest) (*models.Experience, error) {
	experience := &models.Experience{
		ProfileID:     profileID,
		EmployerID:    req.EmployerID,
		NewEmployerID: req.NewEmployerID,
		JobTitle:      req.JobTitle,
		Description:   req.Description,
		StartDate:     req.StartDate,
		EndDate:       req.EndDate,
	}
	if err := config.DB.Create(experience).Error; err != nil {
		return nil, err
	}

	config.DB.Preload("Employer").Preload("NewEmployer").First(&experience, experience.ID)

	return experience, nil
}

func UpdateExperience(userID uuid.UUID, experienceID uint, req *dto.UpdateExperienceRequest) (*models.Experience, error) {
	var profile models.Profile
	if err := config.DB.Where("user_id = ?", userID).First(&profile).Error; err != nil {
		return nil, err
	}
	var experience models.Experience
	if err := config.DB.Where("profile_id = ? AND id = ?", profile.ID, experienceID).First(&experience).Error; err != nil {
		return nil, err
	}
	experience.EmployerID = req.EmployerID
	experience.NewEmployerID = req.NewEmployerID
	experience.JobTitle = req.JobTitle
	experience.Description = req.Description
	experience.StartDate = req.StartDate
	experience.EndDate = req.EndDate
	if err := config.DB.Save(&experience).Error; err != nil {
		return nil, err
	}

	config.DB.Preload("Employer").Preload("NewEmployer").First(&experience, experience.ID)

	return &experience, nil
}

func DeleteExperience(userID uuid.UUID, experienceID uint) error {
	var profile models.Profile
	if err := config.DB.Where("user_id = ?", userID).First(&profile).Error; err != nil {
		return err
	}
	return config.DB.Where("profile_id = ? AND id = ?", profile.ID, experienceID).Delete(&models.Experience{}).Error
}

func CreateNewEmployer(newEmployer *dto.NewEmployerRequest, suggestedBy uuid.UUID) (*models.NewEmployer, error) {
	employer := &models.NewEmployer{
		CompanyName: newEmployer.CompanyName,
		Website:     newEmployer.Website,
		SuggestedBy: suggestedBy,
		Status:      "pending",
	}

	if err := config.DB.Create(employer).Error; err != nil {
		return nil, err
	}

	return employer, nil
}
