
package controllers

import (
	"net/http"
	"strconv"

	"github.com/Alvinferdeveloper/summa-backend/config"
	"github.com/Alvinferdeveloper/summa-backend/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// ApplyToJobRequest define el cuerpo de la solicitud para postular a un empleo.
type ApplyToJobRequest struct {
	CoverLetter string `json:"cover_letter"`
}

// ApplyToJob maneja la lógica para que un candidato postule a un empleo.
func ApplyToJob(c *gin.Context) {
	// 1. Obtener el ID del candidato desde el token JWT
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// 2. Obtener el ID del empleo desde la URL
	jobID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de empleo inválido"})
		return
	}

	// 3. Obtener el perfil del candidato
	var profile models.Profile
	if err := config.DB.Where("user_id = ?", userID).First(&profile).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Perfil de candidato no encontrado"})
		return
	}

	// 4. Validar que no haya postulado previamente
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

	// 5. Procesar el cuerpo de la solicitud
	var req ApplyToJobRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 6. Crear la postulación
	application := models.JobApplication{
		ProfileID:             profile.ID,
		JobPostID:             uint(jobID),
		Status:                "Postulado", // Estado inicial
		CoverLetter:           req.CoverLetter,
		ResumeURLAtApplication: profile.ResumeURL, // Snapshot del CV actual
	}

	if err := config.DB.Create(&application).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo procesar la postulación"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Postulación exitosa"})
}
