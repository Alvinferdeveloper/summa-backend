package services

import (
	"fmt"

	"github.com/Alvinferdeveloper/summa-backend/config"
	"github.com/Alvinferdeveloper/summa-backend/models"
	"gorm.io/gorm"
)

func GetCandidates(page, limit int, filters map[string]string) ([]models.Profile, int64, error) {
	offset := (page - 1) * limit

	var profiles []models.Profile
	var total int64

	baseQuery := config.DB.Model(&models.Profile{}).Where("onboarding_completed = ?", true)

	for key, value := range filters {
		if value == "" {
			continue
		}
		switch key {
		case "country":
			baseQuery = baseQuery.Where("country = ?", value)
		case "disability_type_id":
			baseQuery = baseQuery.Joins("JOIN profile_disability_types ON profile_disability_types.profile_id = profiles.id").Where("profile_disability_types.disability_type_id = ?", value)
		}
	}

	countQuery := baseQuery.Session(&gorm.Session{})
	if err := countQuery.Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to count candidates: %w", err)
	}

	if err := baseQuery.Preload("Skills").Preload("DisabilityTypes").Limit(limit).Offset(offset).Order("created_at desc").Find(&profiles).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to fetch candidates: %w", err)
	}

	return profiles, total, nil
}
