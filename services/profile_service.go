package services

import (
	"github.com/Alvinferdeveloper/summa-backend/config"
	"github.com/Alvinferdeveloper/summa-backend/dto"
	"github.com/Alvinferdeveloper/summa-backend/models"
)

func CompleteOnboarding(req *dto.OnboardingRequest, userID uint) (*models.Profile, error) {
	var profile models.Profile
	if err := config.DB.Where("user_id = ?", userID).First(&profile).Error; err != nil {
		return nil, err
	}

	profile.FirstName = req.FirstName
	profile.LastName = req.LastName
	profile.OnboardingCompleted = true

	var disabilityTypes []models.DisabilityType
	if err := config.DB.Where(req.DisabilityTypeIDs).Find(&disabilityTypes).Error; err != nil {
		return nil, err
	}

	if err := config.DB.Model(&profile).Association("DisabilityTypes").Replace(&disabilityTypes); err != nil {
		return nil, err
	}

	if err := config.DB.Save(&profile).Error; err != nil {
		return nil, err
	}

	return &profile, nil
}

func GetDisabilityTypes() ([]models.DisabilityType, error) {
	var disabilityTypes []models.DisabilityType
	if err := config.DB.Find(&disabilityTypes).Error; err != nil {
		return nil, err
	}
	return disabilityTypes, nil
}
