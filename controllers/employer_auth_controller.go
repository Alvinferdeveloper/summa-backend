package controllers

import (
	"net/http"
	"regexp"

	"github.com/Alvinferdeveloper/summa-backend/models"
	"github.com/Alvinferdeveloper/summa-backend/services"
	"github.com/Alvinferdeveloper/summa-backend/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type EmployerRegisterRequest struct {
	CompanyName string `json:"company_name" binding:"required,min=3"`
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password" binding:"required,min=8,max=60"` // Min 8 chars, max 60 for bcrypt
}

type EmployerLoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func RegisterEmployer(c *gin.Context) {
	var req EmployerRegisterRequest
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
	hashedPassword, err := services.HashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	employer := models.Employer{
		CompanyName: req.CompanyName,
		Email:       req.Email,
		Password:    hashedPassword,
		Role:        "employer",
	}

	if err := services.CreateEmployer(&employer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register employer"})
		return
	}

	token, err := utils.GenerateEmployerJWT(employer.ID, employer.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Employer registered successfully", "accessToken": token})
}

func LoginEmployer(c *gin.Context) {
	var req EmployerLoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	employer, err := services.FindEmployerByEmail(req.Email)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find employer"})
		return
	}

	if !services.CheckPasswordHash(req.Password, employer.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
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
