package dto

import "github.com/Alvinferdeveloper/summa-backend/models"

type AccessibilityNeedResponse struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Category string `json:"category"`
}

func ConvertAccessibilityNeedToDTO(accessibilityNeed models.AccessibilityNeed) AccessibilityNeedResponse {
	return AccessibilityNeedResponse{
		ID:       accessibilityNeed.ID,
		Name:     accessibilityNeed.Name,
		Category: accessibilityNeed.Category,
	}
}
