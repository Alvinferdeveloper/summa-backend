package dto

import "github.com/Alvinferdeveloper/summa-backend/models"

type CategoryResponseDTO struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func ConvertCategoryToDTO(category models.Category) *CategoryResponseDTO {
	return &CategoryResponseDTO{
		ID:   category.ID,
		Name: category.Name,
	}
}
