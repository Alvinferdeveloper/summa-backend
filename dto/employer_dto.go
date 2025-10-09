package dto

import (
	"github.com/Alvinferdeveloper/summa-backend/models"
)

type EmployerRegisterRequest struct {
	CompanyName    string `json:"company_name" binding:"required,min=3"`
	Email          string `json:"email" binding:"required,email"`
	Password       string `json:"password" binding:"required"`
	PhoneNumber    string `json:"phone_number"`
	Country        string `json:"country"`
	FoundationDate string `json:"foundation_date"` // String to parse into time.Time
	Industry       string `json:"industry"`
	Size           string `json:"size"`
	Description    string `json:"description"`
	Dedication     string `json:"dedication"`
	Address        string `json:"address"`
	Website        string `json:"website"`
}

type EmployerResponseDTO struct {
	ID          uint   `json:"id"`
	CompanyName string `json:"company_name"`
	LogoURL     string `json:"logo_url"`
	Industry    string `json:"industry"`
	Email       string `json:"email"`
	Address     string `json:"address"`
}

func ConvertEmployerToDTO(employer models.Employer) *EmployerResponseDTO {
	return &EmployerResponseDTO{
		ID:          employer.ID,
		CompanyName: employer.CompanyName,
		LogoURL:     employer.LogoURL,
		Industry:    employer.Industry,
		Email:       employer.Email,
	}
}
