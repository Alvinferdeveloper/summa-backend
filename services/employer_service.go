package services

import (
	"strings"
	"time"

	"github.com/Alvinferdeveloper/summa-backend/config"
	"github.com/Alvinferdeveloper/summa-backend/dto"
	"github.com/Alvinferdeveloper/summa-backend/models"
	"github.com/Alvinferdeveloper/summa-backend/utils"
	"github.com/google/uuid"
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
		Dedication:  req.Dedication,
		Address:     req.Address,
		Website:     req.Website,
		LogoURL:     req.LogoURL,
	}

	if req.FoundationDate != "" {
		parsedDate, err := time.Parse("2006-01-02", req.FoundationDate)
		if err == nil {
			employer.FoundationDate = &parsedDate
		}
	}

	err = config.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(employer).Error; err != nil {
			return err
		}
		if len(req.AccessibleInfrastructureIDs) > 0 {
			var infrastructures []models.AccessibleInfrastructure
			if err := tx.Where("id IN ?", req.AccessibleInfrastructureIDs).Find(&infrastructures).Error; err != nil {
				return err
			}
			if err := tx.Model(employer).Association("AccessibleInfrastructures").Append(&infrastructures); err != nil {
				return err
			}
		}

		if len(req.InclusiveProgramIDs) > 0 {
			var programs []models.InclusiveProgram
			if err := tx.Where("id IN ?", req.InclusiveProgramIDs).Find(&programs).Error; err != nil {
				return err
			}
			if err := tx.Model(employer).Association("InclusivePrograms").Append(&programs); err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
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

func SearchEmployers(query string) ([]models.Employer, error) {
	var employers []models.Employer
	if err := config.DB.Where("LOWER(company_name) LIKE ?", "%"+strings.ToLower(query)+"%").Find(&employers).Error; err != nil {
		return nil, err
	}
	return employers, nil
}

func FindEmployerByName(name string) (*models.Employer, error) {
	var employer models.Employer
	if err := config.DB.Where("company_name = ?", name).First(&employer).Error; err != nil {
		return nil, err
	}
	return &employer, nil
}

func FindEmployerByID(id uuid.UUID) (*models.Employer, error) {
	var employer models.Employer
	if err := config.DB.First(&employer, id).Error; err != nil {
		return nil, err
	}
	return &employer, nil
}

func UpdateEmployerProfile(id uuid.UUID, req *dto.UpdateEmployerProfileRequest) (*models.Employer, error) {
	var employer models.Employer
	if err := config.DB.First(&employer, id).Error; err != nil {
		return nil, err
	}

	employer.CompanyName = req.CompanyName
	employer.Email = req.Email
	employer.Website = req.Website
	employer.PhoneNumber = req.PhoneNumber
	employer.FoundationDate = req.FoundationDate
	employer.Dedication = req.Dedication
	employer.Country = req.Country
	employer.Industry = req.Industry
	employer.Size = req.Size
	employer.Description = req.Description
	employer.Address = req.Address
	employer.FoundationDate = req.FoundationDate

	if err := config.DB.Save(&employer).Error; err != nil {
		return nil, err
	}

	return &employer, nil
}
