package dto

import (
	"time"

	"github.com/Alvinferdeveloper/summa-backend/models"
)

type JobApplicationResponse struct {
	ID                     uint      `json:"id"`
	CreatedAt              time.Time `json:"created_at"`
	Status                 string    `json:"status"`
	CoverLetter            string    `json:"cover_letter"`
	ResumeURLAtApplication string    `json:"resume_url_at_application"`

	Applicant *ApplicantSummaryResponse `json:"applicant"`
	JobPost   *JobPostSummaryResponse   `json:"job_post"`
}

type ApplicantSummaryResponse struct {
	ProfileID      uint   `json:"profile_id"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Email          string `json:"email"`
	ProfilePicture string `json:"profile_picture"`
}

type JobPostSummaryResponse struct {
	ID          uint                 `json:"id"`
	Title       string               `json:"title"`
	CompanyName string               `json:"company_name"`
	Employer    *EmployerResponseDTO `json:"employer"`
}

func ConvertJobApplicationToDTO(app models.JobApplication) JobApplicationResponse {
	applicantSummary := ApplicantSummaryResponse{
		ProfileID:      app.Profile.ID,
		FirstName:      app.Profile.FirstName,
		LastName:       app.Profile.LastName,
		ProfilePicture: app.Profile.ProfilePicture,
	}

	jobPostSummary := JobPostSummaryResponse{
		ID:          app.JobPost.ID,
		Title:       app.JobPost.Title,
		CompanyName: app.JobPost.Employer.CompanyName,
		Employer:    ConvertEmployerToDTO(app.JobPost.Employer),
	}

	return JobApplicationResponse{
		ID:                     app.ID,
		CreatedAt:              app.CreatedAt,
		Status:                 app.Status,
		CoverLetter:            app.CoverLetter,
		ResumeURLAtApplication: app.ResumeURLAtApplication,
		Applicant:              &applicantSummary,
		JobPost:                &jobPostSummary,
	}
}
