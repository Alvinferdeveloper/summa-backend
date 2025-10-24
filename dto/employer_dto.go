package dto

import (
	"time"

	"github.com/Alvinferdeveloper/summa-backend/models"
	"github.com/google/uuid"
)

type EmployerRegisterRequest struct {
	CompanyName                 string   `json:"company_name"`
	Email                       string   `json:"email"`
	Password                    string   `json:"password"`
	PhoneNumber                 string   `json:"phone_number"`
	Country                     string   `json:"country"`
	FoundationDate              string   `json:"foundation_date"`
	Industry                    string   `json:"industry"`
	Size                        string   `json:"size"`
	Description                 string   `json:"description"`
	Dedication                  string   `json:"dedication"`
	Address                     string   `json:"address"`
	Website                     string   `json:"website"`
	LogoURL                     string   `json:"logo"`
	AccessibleInfrastructureIDs []string `json:"accessible_infrastructure_ids"`
	InclusiveProgramIDs         []string `json:"inclusive_program_ids"`
}

type EmployerResponseDTO struct {
	ID             uuid.UUID  `json:"id"`
	CompanyName    string     `json:"company_name"`
	LogoURL        string     `json:"logo_url"`
	Industry       string     `json:"industry"`
	Email          string     `json:"email"`
	Address        string     `json:"address"`
	FoundationDate *time.Time `json:"foundation_date"`
	PhoneNumber    string     `json:"phone_number"`
	Country        string     `json:"country"`
	Size           string     `json:"size"`
	Description    string     `json:"description"`
	Dedication     string     `json:"dedication"`
	Website        string     `json:"website"`
}

func ConvertEmployerToDTO(employer models.Employer) *EmployerResponseDTO {
	return &EmployerResponseDTO{
		ID:             employer.ID,
		CompanyName:    employer.CompanyName,
		LogoURL:        employer.LogoURL,
		Industry:       employer.Industry,
		Email:          employer.Email,
		Address:        employer.Address,
		FoundationDate: employer.FoundationDate,
		PhoneNumber:    employer.PhoneNumber,
		Country:        employer.Country,
		Size:           employer.Size,
		Description:    employer.Description,
		Dedication:     employer.Dedication,
		Website:        employer.Website,
	}
}

type UpdateEmployerProfileRequest struct {
	CompanyName    string     `json:"company_name" binding:"required,min=3"`
	Email          string     `json:"email" binding:"required,email"`
	PhoneNumber    string     `json:"phone_number"`
	Country        string     `json:"country"`
	FoundationDate *time.Time `json:"foundation_date"` // String to parse into time.Time
	Industry       string     `json:"industry"`
	Size           string     `json:"size"`
	Description    string     `json:"description"`
	Dedication     string     `json:"dedication"`
	Address        string     `json:"address"`
	Website        string     `json:"website"`
}
