package controllers

import (
	"net/http"

	"github.com/Alvinferdeveloper/summa-backend/config"
	"github.com/Alvinferdeveloper/summa-backend/dto"
	"github.com/Alvinferdeveloper/summa-backend/models"
	"github.com/gin-gonic/gin"
)

// GetAllExperienceLevels returns all experience levels.
func GetAllExperienceLevels(c *gin.Context) {
	var experienceLevels []models.ExperienceLevel
	if err := config.DB.Find(&experienceLevels).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve experience levels"})
		return
	}

	var dtos []dto.ExperienceLevelDTO
	for _, el := range experienceLevels {
		dtos = append(dtos, dto.ConvertExperienceLevelToDTO(el))
	}

	c.JSON(http.StatusOK, dtos)
}
