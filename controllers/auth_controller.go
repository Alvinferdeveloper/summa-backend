package controllers

import (
	"net/http"

	"github.com/Alvinferdeveloper/summa-backend/config"
	"github.com/Alvinferdeveloper/summa-backend/models"
	"github.com/Alvinferdeveloper/summa-backend/utils"
	"github.com/gin-gonic/gin"
)

// OAuthRequest defines the expected request body from any OAuth callback.
type OAuthRequest struct {
	Provider   string `json:"provider" binding:"required"`
	ProviderID string `json:"provider_id" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
}

// OAuthCallbackHandler handles the sign-in/sign-up logic for users via any OAuth provider.
func OAuthCallbackHandler(c *gin.Context) {
	var req OAuthRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	var user models.User
	result := config.DB.Where(models.User{Provider: req.Provider, ProviderID: req.ProviderID}).First(&user)

	if result.Error != nil {
		newUser := models.User{
			Provider:   req.Provider,
			ProviderID: req.ProviderID,
			Email:      req.Email,
		}

		createResult := config.DB.Create(&newUser)
		if createResult.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
			return
		}

		profile := models.Profile{UserID: newUser.ID}
		if err := config.DB.Create(&profile).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user profile"})
			return
		}

		user = newUser
	} else {
		// If user exists, update their email in case it changed.
		config.DB.Model(&user).Updates(models.User{Email: req.Email})
	}

	jwtToken, err := utils.GenerateJWT(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"accessToken": jwtToken})
}
