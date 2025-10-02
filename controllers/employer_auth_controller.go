package controllers

import (
	"net/http"
	"regexp"

	"github.com/Alvinferdeveloper/summa-backend/dto"
	"github.com/Alvinferdeveloper/summa-backend/services"
	"github.com/Alvinferdeveloper/summa-backend/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterEmployer(c *gin.Context) {
	var req dto.EmployerRegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !isPasswordStrong(req.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password must contain at least one uppercase letter, one lowercase letter, one digit, and one special character."})
		return
	}

	if _, err := services.FindEmployerByEmail(req.Email); err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Email already registered"})
		return
	}

	input := dto.EmployerRegisterRequest{
		CompanyName:    req.CompanyName,
		Email:          req.Email,
		Password:       req.Password,
		PhoneNumber:    req.PhoneNumber,
		Country:        req.Country,
		FoundationDate: req.FoundationDate,
		Industry:       req.Industry,
		Size:           req.Size,
		Description:    req.Description,
		Address:        req.Address,
		Website:        req.Website,
	}

	_, err := services.RegisterEmployer(&input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register employer"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Employer registered successfully"})
}

func LoginEmployer(c *gin.Context) {
	var req dto.EmployerLoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	employer, err := services.LoginEmployer(req.Email, req.Password)
	if err != nil {
		if err == gorm.ErrRecordNotFound || err == gorm.ErrInvalidData {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "An unexpected error occurred"})
		return
	}

	token, err := utils.GenerateEmployerJWT(employer.ID, employer.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "accessToken": token})
}

func isPasswordStrong(password string) bool {
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
	hasDigit := regexp.MustCompile(`[0-9]`).MatchString(password)
	hasSpecial := regexp.MustCompile(`[!@#$%^&*()_+\-=\[\]{};':"\\|,.<>/?~]`).MatchString(password)

	return hasUpper && hasLower && hasDigit && hasSpecial
}
