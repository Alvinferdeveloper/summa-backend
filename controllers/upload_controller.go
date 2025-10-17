package controllers

import (
	"net/http"

	"github.com/Alvinferdeveloper/summa-backend/config"
	"github.com/Alvinferdeveloper/summa-backend/models"
	"github.com/Alvinferdeveloper/summa-backend/services"
	"github.com/gin-gonic/gin"
)

func UploadEmployerLogo(c *gin.Context) {
	employerID, _ := c.Get("employer_id")

	file, err := c.FormFile("logo")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No se ha subido ningún archivo"})
		return
	}

	logoURL, err := services.UploadFile(file, "logos")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Model(&models.Employer{}).Where("id = ?", employerID).Update("logo_url", logoURL).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo actualizar el logo del empleador"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Logo subido exitosamente", "logo_url": logoURL})
}

func UploadProfilePicture(c *gin.Context) {
	userID, _ := c.Get("user_id")

	file, err := c.FormFile("picture")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No se ha subido ningún archivo"})
		return
	}

	pictureURL, err := services.UploadFile(file, "avatars")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Model(&models.Profile{}).Where("user_id = ?", userID).Update("profile_picture_url", pictureURL).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo actualizar la foto de perfil"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Foto de perfil subida exitosamente", "profile_picture_url": pictureURL})
}

func UploadProfileBanner(c *gin.Context) {
	userID, _ := c.Get("user_id")

	file, err := c.FormFile("banner")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No se ha subido ningún archivo"})
		return
	}

	bannerURL, err := services.UploadFile(file, "banners")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Model(&models.Profile{}).Where("user_id = ?", userID).Update("banner_url", bannerURL).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo actualizar el banner del perfil"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Banner subido exitosamente", "banner_url": bannerURL})
}

func UploadCV(c *gin.Context) {
	userID, _ := c.Get("user_id")

	file, err := c.FormFile("cv")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No se ha subido ningún archivo"})
		return
	}

	resumeURL, err := services.UploadFile(file, "resumes")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Model(&models.Profile{}).Where("user_id = ?", userID).Update("resume_url", resumeURL).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo actualizar el CV del perfil"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "CV subido exitosamente", "resume_url": resumeURL})
}
