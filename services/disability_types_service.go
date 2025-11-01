package services

import (
	"github.com/Alvinferdeveloper/summa-backend/config"
	"github.com/Alvinferdeveloper/summa-backend/models"
)

func GetDisabilityTypes() ([]models.DisabilityType, error) {
	var disabilityTypes []models.DisabilityType
	if err := config.DB.Find(&disabilityTypes).Error; err != nil {
		return nil, err
	}
	return disabilityTypes, nil
}

func CreateDisabilityType(name, description string) (*models.DisabilityType, error) {
	disabilityType := models.DisabilityType{Name: name, Description: description}
	if err := config.DB.Create(&disabilityType).Error; err != nil {
		return nil, err
	}
	return &disabilityType, nil
}

func UpdateDisabilityType(id uint, name, description string) (*models.DisabilityType, error) {
	var disabilityType models.DisabilityType
	if err := config.DB.First(&disabilityType, id).Error; err != nil {
		return nil, err
	}
	disabilityType.Name = name
	disabilityType.Description = description
	if err := config.DB.Save(&disabilityType).Error; err != nil {
		return nil, err
	}
	return &disabilityType, nil
}

func DeleteDisabilityType(id uint) error {
	return config.DB.Delete(&models.DisabilityType{}, id).Error
}
