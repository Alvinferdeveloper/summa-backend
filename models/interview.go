package models

import (
	"time"

	"gorm.io/gorm"
)

type Interview struct {
	gorm.Model
	JobApplicationID        uint      `json:"job_application_id" gorm:"column:job_application_id;not null"`
	ScheduledAt             time.Time `json:"scheduled_at" gorm:"column:scheduled_at;not null"`
	Format                  string    `json:"format" gorm:"column:format;not null"`                                                  // Ej: "Virtual", "Presencial", "Telefónica"
	Notes                   string    `json:"notes" gorm:"type:text"`                                                                // Notas del empleador
	CandidateResponseStatus string    `json:"candidate_response_status" gorm:"column:candidate_response_status;default:'Pendiente'"` // Ej: Pendiente, Aceptada, Rechazada
	RequestedAccommodations string    `json:"requested_accommodations" gorm:"type:text"`                                             // Adaptaciones solicitadas por el candidato

	// Relationships
	JobApplication *JobApplication `json:"job_application,omitempty"`
}
