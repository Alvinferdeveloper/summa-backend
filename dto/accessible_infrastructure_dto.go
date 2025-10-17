package dto

import "github.com/Alvinferdeveloper/summa-backend/models"

// AccessibleInfrastructureDTO is the DTO for accessible infrastructures.
type AccessibleInfrastructureDTO struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

// ConvertAccessibleInfrastructureToDTO converts an AccessibleInfrastructure model to its DTO.
func ConvertAccessibleInfrastructureToDTO(infra models.AccessibleInfrastructure) AccessibleInfrastructureDTO {
	return AccessibleInfrastructureDTO{
		ID:   infra.ID,
		Name: infra.Name,
	}
}
