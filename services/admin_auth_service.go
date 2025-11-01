package services

import (
	"errors"

	"github.com/Alvinferdeveloper/summa-backend/config"
	"github.com/Alvinferdeveloper/summa-backend/models"
	"github.com/Alvinferdeveloper/summa-backend/utils"
)

func LoginAdmin(email string, password string) (string, error) {
	var admin models.Admin
	if err := config.DB.Where("email = ?", email).First(&admin).Error; err != nil {
		return "", err
	}

	if !utils.CheckPasswordHash(password, admin.Password) {
		return "", errors.New("Credenciales inválidas")
	}

	token, err := utils.GenerateAdminJWT(admin.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}
