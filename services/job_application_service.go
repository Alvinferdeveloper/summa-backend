package services

import (
	"fmt"
	"log"

	"github.com/Alvinferdeveloper/summa-backend/config"
	"github.com/Alvinferdeveloper/summa-backend/models"
	"github.com/Alvinferdeveloper/summa-backend/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateJobApplication(profile *models.Profile, jobID uint, coverLetter string) (*models.JobApplication, error) {
	var existingApplication models.JobApplication
	err := config.DB.Where("profile_id = ? AND job_post_id = ?", profile.ID, jobID).First(&existingApplication).Error
	if err == nil {
		return nil, fmt.Errorf("already applied to this job")
	}
	if err != gorm.ErrRecordNotFound {
		return nil, fmt.Errorf("error checking for existing application: %w", err)
	}

	application := &models.JobApplication{
		ProfileID:              profile.ID,
		JobPostID:              jobID,
		Status:                 "Postulado",
		CoverLetter:            coverLetter,
		ResumeURLAtApplication: profile.ResumeURL,
	}

	if err := config.DB.Create(application).Error; err != nil {
		return nil, fmt.Errorf("could not create application: %w", err)
	}

	// Fetch job post and employer details for email
	var jobPost models.JobPost
	if err := config.DB.Preload("Employer").First(&jobPost, jobID).Error; err != nil {
		// Log error but don't block application creation
		log.Printf("Failed to fetch job post for email confirmation: %v", err)
	}

	// Fetch user details for email
	var user models.User
	if err := config.DB.Preload("Profile").First(&user, profile.UserID).Error; err != nil {
		log.Printf("Failed to fetch user for email confirmation: %v", err)
	}

	// Send email confirmation via RabbitMQ
	go func() {
		data := gin.H{
			"ApplicantName": profile.FirstName,
			"JobTitle":      jobPost.Title,
			"CompanyName":   jobPost.Employer.CompanyName,
		}
		body, err := utils.ParseTemplate("job_application_confirmation.html", data)
		if err != nil {
			log.Printf("Failed to parse job application confirmation email template: %v", err)
			return
		}

		task := &EmailTask{
			To:      user.Email,
			Subject: "Confirmación de Postulación a " + jobPost.Title,
			Body:    body,
		}
		if err := GlobalRabbitMQService.PublishEmailTask(task); err != nil {
			log.Printf("Failed to publish job application confirmation email task: %v", err)
		}
	}()

	return application, nil
}

func GetMyApplications(profileID uint) ([]models.JobApplication, error) {
	var applications []models.JobApplication
	if err := config.DB.Preload("JobPost.Employer").Preload("Interview").Where("profile_id = ?", profileID).Order("created_at desc").Find(&applications).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch applications: %w", err)
	}
	return applications, nil
}

func GetJobApplicants(jobID uint, employerID uuid.UUID, page int, limit int) ([]models.JobApplication, int64, error) {
	var jobPost models.JobPost
	if err := config.DB.Where("id = ? AND employer_id = ?", jobID, employerID).First(&jobPost).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, 0, fmt.Errorf("job post not found or does not belong to employer")
		}
		return nil, 0, fmt.Errorf("error verifying job post: %w", err)
	}

	var applications []models.JobApplication
	var total int64

	offset := (page - 1) * limit

	// First, get the total count without pagination
	if err := config.DB.Model(&models.JobApplication{}).Where("job_post_id = ?", jobID).Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to count job applications: %w", err)
	}

	// Then, get the paginated results
	if err := config.DB.Preload("Profile").Preload("Interview").Where("job_post_id = ?", jobID).Order("created_at desc").Limit(limit).Offset(offset).Find(&applications).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to fetch job applications: %w", err)
	}

	return applications, total, nil
}

func UpdateApplicationStatus(applicationID uint, employerID uuid.UUID, status string) (*models.JobApplication, error) {
	var application models.JobApplication
	if err := config.DB.Preload("JobPost.Employer").Preload("Profile").First(&application, applicationID).Error; err != nil {
		return nil, fmt.Errorf("postulación no encontrada")
	}
	if application.JobPost.EmployerID != employerID {
		return nil, fmt.Errorf("no autorizado para modificar esta postulación")
	}

	application.Status = status
	if err := config.DB.Save(&application).Error; err != nil {
		return nil, fmt.Errorf("no se pudo actualizar el estado: %w", err)
	}

	// Send email notification via RabbitMQ
	go func() {
		var templateName string
		var subject string

		switch status {
		case "En revisión":
			templateName = "application_status_in_review.html"
			subject = "Tu postulación para " + application.JobPost.Title + " está en revisión"
		case "Aceptado":
			templateName = "application_status_accepted.html"
			subject = "¡Felicitaciones! Has avanzado en el proceso para " + application.JobPost.Title
		case "Rechazado":
			templateName = "application_status_rejected.html"
			subject = "Actualización sobre tu postulación para " + application.JobPost.Title
		default:
			return
		}

		data := gin.H{
			"ApplicantName": application.Profile.FirstName,
			"JobTitle":      application.JobPost.Title,
			"CompanyName":   application.JobPost.Employer.CompanyName,
		}

		body, err := utils.ParseTemplate(templateName, data)
		if err != nil {
			log.Printf("Failed to parse email template %s: %v", templateName, err)
			return
		}

		task := &EmailTask{
			To:      application.Profile.Email,
			Subject: subject,
			Body:    body,
		}

		if err := GlobalRabbitMQService.PublishEmailTask(task); err != nil {
			log.Printf("Failed to publish status update email task: %v", err)
		}
	}()

	return &application, nil
}
