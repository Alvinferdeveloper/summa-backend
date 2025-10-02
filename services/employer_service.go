package services

import (
	"time"

	"github.com/Alvinferdeveloper/summa-backend/config"
	"github.com/Alvinferdeveloper/summa-backend/dto"
	"github.com/Alvinferdeveloper/summa-backend/models"
	"github.com/Alvinferdeveloper/summa-backend/utils"
	"gorm.io/gorm"
)

func RegisterEmployer(req *dto.EmployerRegisterRequest) (*models.Employer, error) {
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	employer := &models.Employer{
		CompanyName: req.CompanyName,
		Email:       req.Email,
		Password:    hashedPassword,
		Role:        "employer",
		PhoneNumber: req.PhoneNumber,
		Country:     req.Country,
		Industry:    req.Industry,
		Size:        req.Size,
		Description: req.Description,
		Address:     req.Address,
		Website:     req.Website,
	}

	if req.FoundationDate != "" {
		parsedDate, err := time.Parse("2006-01-02", req.FoundationDate)
		if err == nil {
			employer.FoundationDate = &parsedDate
		}
	}

	if err := config.DB.Create(employer).Error; err != nil {
		return nil, err
	}

	return employer, nil
}

func LoginEmployer(email, password string) (*models.Employer, error) {
	employer, err := FindEmployerByEmail(email)
	if err != nil {
		return nil, err
	}

	if !utils.CheckPasswordHash(password, employer.Password) {
		return nil, gorm.ErrInvalidData
	}

	return employer, nil
}

func FindEmployerByEmail(email string) (*models.Employer, error) {
	var employer models.Employer
	if err := config.DB.Where("email = ?", email).First(&employer).Error; err != nil {
		return nil, err
	}
	return &employer, nil
}
