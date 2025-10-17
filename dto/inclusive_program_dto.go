package dto

import "github.com/Alvinferdeveloper/summa-backend/models"

// InclusiveProgramDTO is the DTO for inclusive programs.
type InclusiveProgramDTO struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

// ConvertInclusiveProgramToDTO converts an InclusiveProgram model to its DTO.
func ConvertInclusiveProgramToDTO(prog models.InclusiveProgram) InclusiveProgramDTO {
	return InclusiveProgramDTO{
		ID:   prog.ID,
		Name: prog.Name,
	}
}
