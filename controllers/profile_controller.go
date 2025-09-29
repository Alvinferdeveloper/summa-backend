package controllers

import (
	"net/http"

	"github.com/Alvinferdeveloper/summa-backend/config"
	"github.com/Alvinferdeveloper/summa-backend/models"
	"github.com/gin-gonic/gin"
)

// OnboardingRequest defines the data for the initial profile setup.
type OnboardingRequest struct {
	FirstName         string `json:"first_name" binding:"required"`
	LastName          string `json:"last_name" binding:"required"`
	DisabilityTypeIDs []uint `json:"disability_type_ids" binding:"required,min=1"`
}

// CompleteOnboarding updates the user's profile with essential information.
func CompleteOnboarding(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var req OnboardingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	var profile models.Profile
	if err := config.DB.Where("user_id = ?", userID).First(&profile).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Profile not found"})
		return
	}

	// Update the profile fields
	profile.FirstName = req.FirstName
	profile.LastName = req.LastName
	profile.OnboardingCompleted = true

	// Find the disability types to associate
	var disabilityTypes []models.DisabilityType
	if err := config.DB.Where(req.DisabilityTypeIDs).Find(&disabilityTypes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find disability types"})
		return
	}

	if err := config.DB.Model(&profile).Association("DisabilityTypes").Replace(&disabilityTypes); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update disability types"})
		return
	}

	if err := config.DB.Save(&profile).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save profile"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Profile updated successfully"})
}

func GetDisabilityTypes(c *gin.Context) {
	var disabilityTypes []models.DisabilityType
	if err := config.DB.Find(&disabilityTypes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch disability types"})
		return
	}
	c.JSON(http.StatusOK, disabilityTypes)
}
