package controllers

import (
	"net/http"
	"strconv"

	"github.com/Alvinferdeveloper/summa-backend/dto"
	"github.com/Alvinferdeveloper/summa-backend/services"
	"github.com/gin-gonic/gin"
)

type ApplyToJobRequest struct {
	CoverLetter string `json:"cover_letter"`
}

func ApplyToJob(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	jobID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de empleo inválido"})
		return
	}

	profile, err := services.GetFullProfile(userID.(uint))
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
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	profile, err := services.GetFullProfile(userID.(uint))
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
	employerID, exists := c.Get("employer_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	jobID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de empleo inválido"})
		return
	}

	applications, err := services.GetJobApplicants(uint(jobID), employerID.(uint))
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
