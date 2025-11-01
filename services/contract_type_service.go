
package services

import (
	"github.com/Alvinferdeveloper/summa-backend/config"
	"github.com/Alvinferdeveloper/summa-backend/models"
)

func GetAllContractTypes() ([]models.ContractType, error) {
	var contractTypes []models.ContractType
	if err := config.DB.Find(&contractTypes).Error; err != nil {
		return nil, err
	}
	return contractTypes, nil
}

func CreateContractType(name string) (*models.ContractType, error) {
	contractType := models.ContractType{Name: name}
	if err := config.DB.Create(&contractType).Error; err != nil {
		return nil, err
	}
	return &contractType, nil
}

func UpdateContractType(id uint, name string) (*models.ContractType, error) {
	var contractType models.ContractType
	if err := config.DB.First(&contractType, id).Error; err != nil {
		return nil, err
	}
	contractType.Name = name
	if err := config.DB.Save(&contractType).Error; err != nil {
		return nil, err
	}
	return &contractType, nil
}

func DeleteContractType(id uint) error {
	return config.DB.Delete(&models.ContractType{}, id).Error
}
