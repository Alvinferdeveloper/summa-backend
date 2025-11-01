
package services

import (
	"github.com/Alvinferdeveloper/summa-backend/config"
	"github.com/Alvinferdeveloper/summa-backend/models"
)

func GetAllExperienceLevels() ([]models.ExperienceLevel, error) {
	var experienceLevels []models.ExperienceLevel
	if err := config.DB.Find(&experienceLevels).Error; err != nil {
		return nil, err
	}
	return experienceLevels, nil
}

func CreateExperienceLevel(name string) (*models.ExperienceLevel, error) {
	experienceLevel := models.ExperienceLevel{Name: name}
	if err := config.DB.Create(&experienceLevel).Error; err != nil {
		return nil, err
	}
	return &experienceLevel, nil
}

func UpdateExperienceLevel(id uint, name string) (*models.ExperienceLevel, error) {
	var experienceLevel models.ExperienceLevel
	if err := config.DB.First(&experienceLevel, id).Error; err != nil {
		return nil, err
	}
	experienceLevel.Name = name
	if err := config.DB.Save(&experienceLevel).Error; err != nil {
		return nil, err
	}
	return &experienceLevel, nil
}

func DeleteExperienceLevel(id uint) error {
	return config.DB.Delete(&models.ExperienceLevel{}, id).Error
}
