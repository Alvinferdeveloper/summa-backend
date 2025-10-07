package controllers

import (
	"net/http"
	"strconv"

	"github.com/Alvinferdeveloper/summa-backend/config"
	"github.com/Alvinferdeveloper/summa-backend/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

	var profile models.Profile
	if err := config.DB.Where("user_id = ?", userID).First(&profile).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Perfil de candidato no encontrado"})
		return
	}

	var existingApplication models.JobApplication
	err = config.DB.Where("profile_id = ? AND job_post_id = ?", profile.ID, jobID).First(&existingApplication).Error
	if err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Ya has postulado a este empleo"})
		return
	}
	if err != gorm.ErrRecordNotFound {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al verificar la postulación"})
		return
	}

	var req ApplyToJobRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	application := models.JobApplication{
		ProfileID:              profile.ID,
		JobPostID:              uint(jobID),
		Status:                 "Postulado",
		CoverLetter:            req.CoverLetter,
		ResumeURLAtApplication: profile.ResumeURL,
	}

	if err := config.DB.Create(&application).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo procesar la postulación"})
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

	var profile models.Profile
	if err := config.DB.Where("user_id = ?", userID).First(&profile).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Perfil de candidato no encontrado"})
		return
	}

	var applications []models.JobApplication
	if err := config.DB.Preload("JobPost.Employer").Where("profile_id = ?", profile.ID).Order("created_at desc").Find(&applications).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener las postulaciones"})
		return
	}

	c.JSON(http.StatusOK, applications)
}
