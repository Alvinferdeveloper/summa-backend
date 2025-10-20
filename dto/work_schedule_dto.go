package dto

import "github.com/Alvinferdeveloper/summa-backend/models"

// WorkScheduleDTO is the DTO for work schedules.
type WorkScheduleDTO struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

// ConvertWorkScheduleToDTO converts a WorkSchedule model to its DTO.
func ConvertWorkScheduleToDTO(ws models.WorkSchedule) WorkScheduleDTO {
	return WorkScheduleDTO{
		ID:   ws.ID,
		Name: ws.Name,
	}
}
