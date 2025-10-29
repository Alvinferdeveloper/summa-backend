package dto

import (
	"time"

	"github.com/Alvinferdeveloper/summa-backend/models"
)

type ApplicantSummaryResponse struct {
	ProfileID         uint   `json:"profile_id"`
	FirstName         string `json:"first_name"`
	LastName          string `json:"last_name"`
	Email             string `json:"email"`
	ResumeURL         string `json:"resume_url"`
	ProfilePictureURL string `json:"profile_picture_url"`
}

type JobPostSummaryResponse struct {
	ID          uint                 `json:"id"`
	Title       string               `json:"title"`
	CompanyName string               `json:"company_name"`
	Employer    *EmployerResponseDTO `json:"employer"`
}

type UpdateApplicationStatusRequest struct {
	Status string `json:"status" binding:"required"`
}

type JobApplicationResponse struct {
	ID                     uint      `json:"id"`
	CreatedAt              time.Time `json:"created_at"`
	Status                 string    `json:"status"`
	CoverLetter            string    `json:"cover_letter"`
	ResumeURLAtApplication string    `json:"resume_url_at_application"`

	Applicant *ApplicantSummaryResponse `json:"applicant"`
	JobPost   *JobPostSummaryResponse   `json:"job_post"`
	Interview *InterviewResponse        `json:"interview,omitempty"`
}

type InterviewResponse struct {
	ID                      uint      `json:"id"`
	ScheduledAt             time.Time `json:"scheduled_at"`
	Format                  string    `json:"format"`
	Notes                   string    `json:"notes"`
	CandidateResponseStatus string    `json:"candidate_response_status"`
	RequestedAccommodations string    `json:"requested_accommodations"`
}

func ConvertJobApplicationToDTO(app models.JobApplication) JobApplicationResponse {
	applicantSummary := ApplicantSummaryResponse{
		ProfileID:         app.Profile.ID,
		FirstName:         app.Profile.FirstName,
		LastName:          app.Profile.LastName,
		Email:             app.Profile.Email,
		ResumeURL:         app.Profile.ResumeURL,
		ProfilePictureURL: app.Profile.ProfilePictureURL,
	}

	jobPostSummary := JobPostSummaryResponse{
		ID:          app.JobPost.ID,
		Title:       app.JobPost.Title,
		CompanyName: app.JobPost.Employer.CompanyName,
		Employer:    ConvertEmployerToDTO(app.JobPost.Employer),
	}

	var interviewDTO *InterviewResponse
	if app.Interview != nil {
		interviewDTO = &InterviewResponse{
			ID:                      app.Interview.ID,
			ScheduledAt:             app.Interview.ScheduledAt,
			Format:                  app.Interview.Format,
			Notes:                   app.Interview.Notes,
			CandidateResponseStatus: app.Interview.CandidateResponseStatus,
			RequestedAccommodations: app.Interview.RequestedAccommodations,
		}
	}

	return JobApplicationResponse{
		ID:                     app.ID,
		CreatedAt:              app.CreatedAt,
		Status:                 app.Status,
		CoverLetter:            app.CoverLetter,
		ResumeURLAtApplication: app.ResumeURLAtApplication,
		Applicant:              &applicantSummary,
		JobPost:                &jobPostSummary,
		Interview:              interviewDTO,
	}
}
