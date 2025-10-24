package controllers

import (
	"net/http"

	"github.com/Alvinferdeveloper/summa-backend/config"
	"github.com/Alvinferdeveloper/summa-backend/dto"
	"github.com/Alvinferdeveloper/summa-backend/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateUniversitySuggestion(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var req dto.NewUniversitySuggestionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	suggestion := models.UniversitySuggestion{
		SuggestedName: req.SuggestedName,
		Country:       req.Country,
		Website:       req.Website,
		SuggestedBy:   userID.(uuid.UUID),
	}

	if err := config.DB.Create(&suggestion).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo crear la sugerencia de universidad"})
		return
	}

	c.JSON(http.StatusCreated, suggestion)
}
