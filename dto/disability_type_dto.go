package dto

import "github.com/Alvinferdeveloper/summa-backend/models"

type DisabilityTypeResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func ConvertDisabilityTypeToDTO(dt models.DisabilityType) DisabilityTypeResponse {
	return DisabilityTypeResponse{
		ID:          dt.ID,
		Name:        dt.Name,
		Description: dt.Description,
	}
}
