package services

import (
	"github.com/Alvinferdeveloper/summa-backend/config"
	"github.com/Alvinferdeveloper/summa-backend/dto"
	"github.com/Alvinferdeveloper/summa-backend/models"
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

func UpdateExperience(profileID uint, experienceID uint, req *dto.UpdateExperienceRequest) (*models.Experience, error) {
	var experience models.Experience
	if err := config.DB.Where("profile_id = ? AND id = ?", profileID, experienceID).First(&experience).Error; err != nil {
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

func DeleteExperience(profileID uint, experienceID uint) error {
	return config.DB.Where("profile_id = ?", profileID).Delete(&models.Experience{}, experienceID).Error
}

func CreateNewEmployer(newEmployer *dto.NewEmployerRequest, suggestedBy uint) error {
	if err := config.DB.Create(&models.NewEmployer{
		CompanyName: newEmployer.CompanyName,
		Website:     newEmployer.Website,
		SuggestedBy: suggestedBy,
		Status:      "pending",
	}).Error; err != nil {
		return err
	}

	return nil
}
