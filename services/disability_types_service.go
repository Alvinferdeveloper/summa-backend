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
