package controllers

import (
	"net/http"

	"github.com/Alvinferdeveloper/summa-backend/config"
	"github.com/Alvinferdeveloper/summa-backend/dto"
	"github.com/Alvinferdeveloper/summa-backend/models"
	"github.com/Alvinferdeveloper/summa-backend/services"
	"github.com/gin-gonic/gin"
)

func CreateNewEmployer(c *gin.Context) {
	userID, _ := c.Get("user_id")

	var req dto.NewEmployerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.CreateNewEmployer(&req, userID.(uint)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo crear la sugerencia de empresa"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Sugerencia de empresa creada exitosamente"})
}

func CreateExperience(c *gin.Context) {
	userID, _ := c.Get("user_id")

	var profile models.Profile
	if err := config.DB.Where("user_id = ?", userID).First(&profile).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Perfil de usuario no encontrado"})
		return
	}

	var req dto.CreateExperienceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	experience, err := services.CreateExperience(profile.ID, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo crear la experiencia"})
		return
	}

	experienceDTO := dto.ConvertExperienceToDTO(*experience)

	c.JSON(http.StatusCreated, experienceDTO)
}
