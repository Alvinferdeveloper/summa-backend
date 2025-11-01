package dto

import (
	"time"

	"github.com/Alvinferdeveloper/summa-backend/models"
)

// DisabilityTypeResponse define el DTO para el tipo de discapacidad.
type DisabilityTypeResponse struct {
	ID          uint      `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}

// ConvertDisabilityTypeToDTO convierte un modelo DisabilityType a su DTO.
func ConvertDisabilityTypeToDTO(dt models.DisabilityType) DisabilityTypeResponse {
	return DisabilityTypeResponse{
		ID:          dt.ID,
		CreatedAt:   dt.CreatedAt,
		UpdatedAt:   dt.UpdatedAt,
		Name:        dt.Name,
		Description: dt.Description,
	}
}