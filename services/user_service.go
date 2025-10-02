package services

import (
	"github.com/Alvinferdeveloper/summa-backend/config"
	"github.com/Alvinferdeveloper/summa-backend/models"
)

func FindOrCreateUser(provider, providerID, email string) (*models.User, *models.Profile, error) {
	var user models.User
	result := config.DB.Where(models.User{Provider: provider, ProviderID: providerID}).First(&user)

	if result.Error != nil {
		newUser := models.User{
			Provider:   provider,
			ProviderID: providerID,
			Email:      email,
		}

		if err := config.DB.Create(&newUser).Error; err != nil {
			return nil, nil, err
		}

		profile := models.Profile{UserID: newUser.ID}
		if err := config.DB.Create(&profile).Error; err != nil {
			// Consider rolling back in future feature
			return nil, nil, err
		}

		return &newUser, &profile, nil
	}

	var profile models.Profile
	if err := config.DB.Where("user_id = ?", user.ID).First(&profile).Error; err != nil {
		return nil, nil, err
	}

	return &user, &profile, nil
}
