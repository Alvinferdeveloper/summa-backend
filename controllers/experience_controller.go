package controllers

import (
	"net/http"

	"github.com/Alvinferdeveloper/summa-backend/config"
	"github.com/Alvinferdeveloper/summa-backend/dto"
	"github.com/Alvinferdeveloper/summa-backend/models"
	"github.com/gin-gonic/gin"
)

func CreateNewEmployer(c *gin.Context) {
	userID, _ := c.Get("user_id")

	var req dto.NewEmployerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newEmployer := models.NewEmployer{
		CompanyName: req.CompanyName,
		Website:     req.Website,
		SuggestedBy: userID.(uint),
	}

	if err := config.DB.Create(&newEmployer).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo crear la sugerencia de empresa"})
		return
	}

	c.JSON(http.StatusCreated, newEmployer)
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

	experience := models.Experience{
		ProfileID:     profile.ID,
		JobTitle:      req.JobTitle,
		Description:   req.Description,
		StartDate:     req.StartDate,
		EndDate:       req.EndDate,
		EmployerID:    req.EmployerID,
		NewEmployerID: req.NewEmployerID,
	}

	if err := config.DB.Create(&experience).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo crear la experiencia"})
		return
	}

	c.JSON(http.StatusCreated, experience)
}
