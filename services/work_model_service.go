
package services

import (
	"github.com/Alvinferdeveloper/summa-backend/config"
	"github.com/Alvinferdeveloper/summa-backend/models"
)

func GetAllWorkModels() ([]models.WorkModel, error) {
	var workModels []models.WorkModel
	if err := config.DB.Find(&workModels).Error; err != nil {
		return nil, err
	}
	return workModels, nil
}

func CreateWorkModel(name string) (*models.WorkModel, error) {
	workModel := models.WorkModel{Name: name}
	if err := config.DB.Create(&workModel).Error; err != nil {
		return nil, err
	}
	return &workModel, nil
}

func UpdateWorkModel(id uint, name string) (*models.WorkModel, error) {
	var workModel models.WorkModel
	if err := config.DB.First(&workModel, id).Error; err != nil {
		return nil, err
	}
	workModel.Name = name
	if err := config.DB.Save(&workModel).Error; err != nil {
		return nil, err
	}
	return &workModel, nil
}

func DeleteWorkModel(id uint) error {
	return config.DB.Delete(&models.WorkModel{}, id).Error
}
