package dto

import "github.com/Alvinferdeveloper/summa-backend/models"

type SkillResponseDTO struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func ConvertSkillToDTO(skill models.Skill) SkillResponseDTO {
	return SkillResponseDTO{
		ID:   skill.ID,
		Name: skill.Name,
	}
}
