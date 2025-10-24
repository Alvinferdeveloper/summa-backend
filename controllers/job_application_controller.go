package controllers

import (
	"net/http"
	"strconv"

	"github.com/Alvinferdeveloper/summa-backend/dto"
	"github.com/Alvinferdeveloper/summa-backend/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ApplyToJobRequest struct {
	CoverLetter string `json:"cover_letter"`
}

func ApplyToJob(c *gin.Context) {
	userID, _ := c.Get("user_id")

	jobID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de empleo inválido"})
		return
	}

	profile, err := services.GetFullProfile(userID.(uuid.UUID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Perfil de candidato no encontrado"})
		return
	}

	var req ApplyToJobRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = services.CreateJobApplication(profile, uint(jobID), req.CoverLetter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Postulación exitosa"})
}

func GetMyApplications(c *gin.Context) {
	userID, _ := c.Get("user_id")

	profile, err := services.GetFullProfile(userID.(uuid.UUID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Perfil de candidato no encontrado"})
		return
	}

	applications, err := services.GetMyApplications(profile.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var applicationDTOs []dto.JobApplicationResponse
	for _, app := range applications {
		applicationDTOs = append(applicationDTOs, dto.ConvertJobApplicationToDTO(app))
	}

	c.JSON(http.StatusOK, applicationDTOs)
}

func GetJobApplicants(c *gin.Context) {
	employerID, _ := c.Get("employer_id")

	jobID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de empleo inválido"})
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	applications, total, err := services.GetJobApplicants(uint(jobID), employerID.(uuid.UUID), page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var applicationDTOs []dto.JobApplicationResponse
	for _, app := range applications {
		applicationDTOs = append(applicationDTOs, dto.ConvertJobApplicationToDTO(app))
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  applicationDTOs,
		"total": total,
		"page":  page,
		"limit": limit,
	})
}

func UpdateApplicationStatus(c *gin.Context) {
	employerID, _ := c.Get("employer_id")

	applicationID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de postulación inválido"})
		return
	}

	var req dto.UpdateApplicationStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	application, err := services.UpdateApplicationStatus(uint(applicationID), employerID.(uuid.UUID), req.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Estado de la postulación actualizado", "application": application})
}
