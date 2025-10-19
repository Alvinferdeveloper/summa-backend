package dto

import "github.com/Alvinferdeveloper/summa-backend/models"

type ContractTypeDTO struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func ConvertContractTypeToDTO(ct models.ContractType) ContractTypeDTO {
	return ContractTypeDTO{
		ID:   ct.ID,
		Name: ct.Name,
	}
}
