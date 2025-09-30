package services

import (
	"github.com/Alvinferdeveloper/summa-backend/config"
	"github.com/Alvinferdeveloper/summa-backend/models"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
func CreateEmployer(employer *models.Employer) error {
	return config.DB.Create(employer).Error
}

func FindEmployerByEmail(email string) (*models.Employer, error) {
	var employer models.Employer
	if err := config.DB.Where("email = ?", email).First(&employer).Error; err != nil {
		return nil, err
	}
	return &employer, nil
}
