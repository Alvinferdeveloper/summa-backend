package dto

import "github.com/Alvinferdeveloper/summa-backend/models"

// WorkModelDTO is the DTO for work models.
type WorkModelDTO struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

// ConvertWorkModelToDTO converts a WorkModel model to its DTO.
func ConvertWorkModelToDTO(wm models.WorkModel) WorkModelDTO {
	return WorkModelDTO{
		ID:   wm.ID,
		Name: wm.Name,
	}
}
