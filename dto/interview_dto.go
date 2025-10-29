package dto

import "time"

type ScheduleInterviewRequest struct {
	JobApplicationID uint      `json:"job_application_id" binding:"required"`
	ScheduledAt      time.Time `json:"scheduled_at" binding:"required"`
	Format           string    `json:"format" binding:"required"`
	Notes            string    `json:"notes"`
}

type RespondToInterviewRequest struct {
	Status                  string `json:"status" binding:"required,oneof=Aceptada Rechazada"`
	RequestedAccommodations string `json:"requested_accommodations"`
}
