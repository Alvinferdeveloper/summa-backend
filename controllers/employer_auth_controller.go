package controllers

import (
	"log"
	"net/http"
	"regexp"

	"github.com/Alvinferdeveloper/summa-backend/dto"
	"github.com/Alvinferdeveloper/summa-backend/services"
	"github.com/Alvinferdeveloper/summa-backend/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterEmployer(c *gin.Context) {
	const MAX_UPLOAD_SIZE = 10 << 30 // 10 MB
	if err := c.Request.ParseMultipartForm(MAX_UPLOAD_SIZE); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error parsing form data or file exceeds 10MB limit."})
		println(err.Error())
		return
	}

	req := dto.EmployerRegisterRequest{
		CompanyName:    c.Request.FormValue("company_name"),
		Password:       c.Request.FormValue("password"),
		Email:          c.Request.FormValue("email"),
		PhoneNumber:    c.Request.FormValue("phone_number"),
		Country:        c.Request.FormValue("country"),
		FoundationDate: c.Request.FormValue("foundation_date"),
		Industry:       c.Request.FormValue("industry"),
		Size:           c.Request.FormValue("size"),
		Description:    c.Request.FormValue("description"),
		Dedication:     c.Request.FormValue("dedication"),
		Address:        c.Request.FormValue("address"),
		Website:        c.Request.FormValue("website"),
	}

	if req.CompanyName == "" || req.Email == "" || req.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Los campos name, email y password son obligatorios"})
		return
	}

	if !isPasswordStrong(req.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password must contain at least one uppercase letter, one lowercase letter, one digit, and one special character."})
		return
	}

	employerName, err := services.FindEmployerByName(req.CompanyName)
	if employerName != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "El nombre de la empresa ya esta registrado"})
		return
	}

	if err != nil && err != gorm.ErrRecordNotFound {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "An unexpected error occurred"})
		return
	}

	employer, err := services.FindEmployerByEmail(req.Email)
	if employer != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "El correo electronico ya esta registrado"})
		return
	}

	if err != nil && err != gorm.ErrRecordNotFound {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "An unexpected error occurred"})
		return
	}

	file, err := c.FormFile("logo")
	var logoURL string
	if err == nil {
		logoURL, err = services.UploadFile(file, "logos")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload logo"})
			return
		}
	}

	req.LogoURL = logoURL

	_, err = services.RegisterEmployer(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register employer"})
		return
	}

	// Publish welcome email task
	go func() {
		body, err := utils.ParseTemplate("employer_welcome.html", gin.H{"CompanyName": req.CompanyName})
		if err != nil {
			log.Printf("Failed to parse welcome email template for %s: %v", req.Email, err)
			return
		}

		task := &services.EmailTask{
			To:      req.Email,
			Subject: "Bienvenido a Summa",
			Body:    body,
		}
		if err := services.GlobalRabbitMQService.PublishEmailTask(task); err != nil {
			log.Printf("Failed to publish welcome email task for %s: %v", req.Email, err)
		}
	}()

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
