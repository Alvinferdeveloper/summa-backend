package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Alvinferdeveloper/summa-backend/config"
	"github.com/Alvinferdeveloper/summa-backend/models"
	"github.com/gin-gonic/gin"
)

type CreateJobPostRequest struct {
	Title                 string   `json:"title" binding:"required"`
	Location              string   `json:"location" binding:"required"`
	WorkModel             string   `json:"workModel" binding:"required"`
	ContractType          string   `json:"contractType" binding:"required"`
	Description           string   `json:"description" binding:"required"`
	Responsibilities      string   `json:"responsibilities" binding:"required"`
	Requirements          string   `json:"requirements" binding:"required"`
	AccessibilityFeatures []string `json:"accessibilityFeatures"`
}

func CreateJobPost(c *gin.Context) {
	employerID, exists := c.Get("employer_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var req CreateJobPostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	accessibilityFeaturesJSON, err := json.Marshal(req.AccessibilityFeatures)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process accessibility features"})
		return
	}

	jobPost := models.JobPost{
		EmployerID:            employerID.(uint),
		Title:                 req.Title,
		Location:              req.Location,
		WorkModel:             req.WorkModel,
		ContractType:          req.ContractType,
		Description:           req.Description,
		Responsibilities:      req.Responsibilities,
		Requirements:          req.Requirements,
		AccessibilityFeatures: string(accessibilityFeaturesJSON),
	}

	if err := config.DB.Create(&jobPost).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create job post"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Job post created successfully", "jobPost": jobPost})
}

func GetJobPosts(c *gin.Context) {
	var page, limit int
	page, _ = strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ = strconv.Atoi(c.DefaultQuery("limit", "10"))

	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	offset := (page - 1) * limit

	var jobPosts []models.JobPost
	var total int64

	config.DB.Model(&models.JobPost{}).Count(&total)

	result := config.DB.Preload("Employer").Limit(limit).Offset(offset).Order("created_at desc").Find(&jobPosts)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch job posts"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":      jobPosts,
		"total":     total,
		"page":      page,
		"limit":     limit,
		"next_page": page*limit < int(total),
	})
}
