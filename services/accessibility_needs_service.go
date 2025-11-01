
package services

import (
	"github.com/Alvinferdeveloper/summa-backend/config"
	"github.com/Alvinferdeveloper/summa-backend/models"
)

func GetAllAccessibilityNeeds() ([]models.AccessibilityNeed, error) {
	var accessibilityNeeds []models.AccessibilityNeed
	if err := config.DB.Find(&accessibilityNeeds).Error; err != nil {
		return nil, err
	}
	return accessibilityNeeds, nil
}

func CreateAccessibilityNeed(name, category string) (*models.AccessibilityNeed, error) {
	accessibilityNeed := models.AccessibilityNeed{Name: name, Category: category}
	if err := config.DB.Create(&accessibilityNeed).Error; err != nil {
		return nil, err
	}
	return &accessibilityNeed, nil
}

func UpdateAccessibilityNeed(id uint, name, category string) (*models.AccessibilityNeed, error) {
	var accessibilityNeed models.AccessibilityNeed
	if err := config.DB.First(&accessibilityNeed, id).Error; err != nil {
		return nil, err
	}
	accessibilityNeed.Name = name
	accessibilityNeed.Category = category
	if err := config.DB.Save(&accessibilityNeed).Error; err != nil {
		return nil, err
	}
	return &accessibilityNeed, nil
}

func DeleteAccessibilityNeed(id uint) error {
	return config.DB.Delete(&models.AccessibilityNeed{}, id).Error
}
