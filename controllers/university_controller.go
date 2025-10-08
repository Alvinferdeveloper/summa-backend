package controllers

import (
	"net/http"
	"strings"

	"github.com/Alvinferdeveloper/summa-backend/config"
	"github.com/Alvinferdeveloper/summa-backend/dto"
	"github.com/Alvinferdeveloper/summa-backend/models"
	"github.com/gin-gonic/gin"
)

func SearchUniversities(c *gin.Context) {
	query := c.Query("q")
	if len(query) < 2 {
		c.JSON(http.StatusOK, []models.University{})
		return
	}

	var universities []models.University
	if err := config.DB.Where("LOWER(name) LIKE ?", "%"+strings.ToLower(query)+"%").Limit(10).Find(&universities).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al buscar universidades"})
		return
	}

	var universityDTOs []dto.UniversityResponse
	for _, university := range universities {
		universityDTOs = append(universityDTOs, dto.ConvertUniversityToDTO(university))
	}

	c.JSON(http.StatusOK, universityDTOs)
}
