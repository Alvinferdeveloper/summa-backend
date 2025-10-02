package controllers

import (
	"net/http"

	"github.com/Alvinferdeveloper/summa-backend/services"
	"github.com/Alvinferdeveloper/summa-backend/utils"
	"github.com/gin-gonic/gin"
)

type OAuthRequest struct {
	Provider   string `json:"provider" binding:"required"`
	ProviderID string `json:"provider_id" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
}

func OAuthCallbackHandler(c *gin.Context) {
	var req OAuthRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	user, profile, err := services.FindOrCreateUser(req.Provider, req.ProviderID, req.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process user information"})
		return
	}

	jwtToken, err := utils.GenerateJWT(user.ID, profile.OnboardingCompleted)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"accessToken": jwtToken, "onboarding_completed": profile.OnboardingCompleted})
}
