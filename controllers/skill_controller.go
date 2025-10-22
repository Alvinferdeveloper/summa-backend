package controllers

import (
	"net/http"

	"github.com/Alvinferdeveloper/summa-backend/config"
	"github.com/Alvinferdeveloper/summa-backend/dto"
	"github.com/Alvinferdeveloper/summa-backend/models"
	"github.com/gin-gonic/gin"
)

func GetAllSkills(c *gin.Context) {
	var skills []models.Skill
	if err := config.DB.Find(&skills).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve skills"})
		return
	}

	var dtos []dto.SkillResponseDTO
	for _, s := range skills {
		dtos = append(dtos, dto.ConvertSkillToDTO(s))
	}

	c.JSON(http.StatusOK, dtos)
}
