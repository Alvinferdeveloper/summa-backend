package dto

import "github.com/Alvinferdeveloper/summa-backend/models"

type NewEmployerRequest struct {
	CompanyName string `json:"company_name" binding:"required"`
	Website     string `json:"website"`
}

type NewEmployerResponseDTO struct {
	ID          uint   `json:"id"`
	CompanyName string `json:"company_name"`
	Website     string `json:"website"`
	Status      string `json:"status"`
}

func ConverNewEmployerToDTO(newEmployer models.NewEmployer) *NewEmployerResponseDTO {
	return &NewEmployerResponseDTO{
		ID:          newEmployer.ID,
		CompanyName: newEmployer.CompanyName,
		Website:     newEmployer.Website,
		Status:      newEmployer.Status,
	}
}
