package dto

import "github.com/Alvinferdeveloper/summa-backend/models"

// ExperienceLevelDTO is the DTO for experience levels.
type ExperienceLevelDTO struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

// ConvertExperienceLevelToDTO converts an ExperienceLevel model to its DTO.
func ConvertExperienceLevelToDTO(el models.ExperienceLevel) ExperienceLevelDTO {
	return ExperienceLevelDTO{
		ID:   el.ID,
		Name: el.Name,
	}
}
