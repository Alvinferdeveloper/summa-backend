
package services

import (
	"github.com/Alvinferdeveloper/summa-backend/config"
	"github.com/Alvinferdeveloper/summa-backend/models"
)

func GetAllCategories() ([]models.Category, error) {
	var categories []models.Category
	if err := config.DB.Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func CreateCategory(name string) (*models.Category, error) {
	category := models.Category{Name: name}
	if err := config.DB.Create(&category).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func UpdateCategory(id uint, name string) (*models.Category, error) {
	var category models.Category
	if err := config.DB.First(&category, id).Error; err != nil {
		return nil, err
	}
	category.Name = name
	if err := config.DB.Save(&category).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func DeleteCategory(id uint) error {
	return config.DB.Delete(&models.Category{}, id).Error
}
