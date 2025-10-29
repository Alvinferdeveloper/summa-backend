package services

import (
	"fmt"
	"log"

	"github.com/Alvinferdeveloper/summa-backend/config"
	"github.com/Alvinferdeveloper/summa-backend/dto"
	"github.com/Alvinferdeveloper/summa-backend/models"
	"github.com/Alvinferdeveloper/summa-backend/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func ScheduleInterview(employerID uuid.UUID, req *dto.ScheduleInterviewRequest) (*models.Interview, error) {
	var application models.JobApplication
	if err := config.DB.Joins("JOIN job_posts ON job_posts.id = job_applications.job_post_id").Where("job_applications.id = ? AND job_posts.employer_id = ?", req.JobApplicationID, employerID).First(&application).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("postulación no encontrada o no autorizada")
		}
		return nil, err
	}

	interview := models.Interview{
		JobApplicationID: req.JobApplicationID,
		ScheduledAt:      req.ScheduledAt,
		Format:           req.Format,
		Notes:            req.Notes,
	}

	if err := config.DB.Create(&interview).Error; err != nil {
		return nil, err
	}

	application.Status = "Entrevista"
	if err := config.DB.Save(&application).Error; err != nil {
		return nil, err
	}

	return &interview, nil
}

func RespondToInterview(profileID uint, interviewID uint, req *dto.RespondToInterviewRequest) (*models.Interview, error) {
	var interview models.Interview
	if err := config.DB.Joins("JOIN job_applications ON job_applications.id = interviews.job_application_id").Preload("JobApplication.Profile").Preload("JobApplication.JobPost.Employer").Where("interviews.id = ? AND job_applications.profile_id = ?", interviewID, profileID).First(&interview).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("entrevista no encontrada o no autorizada")
		}
		return nil, err
	}

	interview.CandidateResponseStatus = req.Status
	interview.RequestedAccommodations = req.RequestedAccommodations

	if err := config.DB.Save(&interview).Error; err != nil {
		return nil, err
	}

	go func() {
		data := gin.H{
			"EmployerName":   interview.JobApplication.JobPost.Employer.CompanyName,
			"ApplicantName":  interview.JobApplication.Profile.FirstName + " " + interview.JobApplication.Profile.LastName,
			"JobTitle":       interview.JobApplication.JobPost.Title,
			"ResponseStatus": req.Status,
			"Accommodations": req.RequestedAccommodations,
		}

		body, err := utils.ParseTemplate("interview_response_notification.html", data)
		if err != nil {
			log.Printf("Failed to parse interview response email template: %v", err)
			return
		}

		task := &EmailTask{
			To:      interview.JobApplication.JobPost.Employer.Email,
			Subject: "Respuesta de " + interview.JobApplication.Profile.FirstName + " a la Invitación de Entrevista",
			Body:    body,
		}

		if err := GlobalRabbitMQService.PublishEmailTask(task); err != nil {
			log.Printf("Failed to publish interview response email task: %v", err)
		}
	}()

	return &interview, nil
}
