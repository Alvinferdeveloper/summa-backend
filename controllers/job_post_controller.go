package controllers

import (
	"net/http"
	"strconv"

	"github.com/Alvinferdeveloper/summa-backend/dto"
	"github.com/Alvinferdeveloper/summa-backend/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateJobPost(c *gin.Context) {
	employerID, _ := c.Get("employer_id")

	var req dto.CreateJobPostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	jobPost, err := services.CreateJobPost(&req, employerID.(uuid.UUID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create job post"})
		return
	}

	jobPostDTO := dto.ConvertJobPostToDTO(*jobPost)

	c.JSON(http.StatusCreated, gin.H{"message": "Job post created successfully", "jobPost": jobPostDTO})
}

func GetJobPosts(c *gin.Context) {
	var page, limit int
	page, _ = strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ = strconv.Atoi(c.DefaultQuery("limit", "10"))

	var userID *uuid.UUID
	if id, exists := c.Get("user_id"); exists {
		uid := id.(uuid.UUID)
		userID = &uid
	}

	filters := make(map[string]string)
	filters["is_urgent"] = c.Query("is_urgent")
	filters["date_posted"] = c.Query("date_posted")
	filters["category_id"] = c.Query("category_id")
	filters["work_schedule_id"] = c.Query("work_schedule_id")
	filters["contract_type_id"] = c.Query("contract_type_id")
	filters["salary"] = c.Query("salary")
	filters["experience_level_id"] = c.Query("experience_level_id")
	filters["work_model_id"] = c.Query("work_model_id")
	filters["disability_type_id"] = c.Query("disability_type_id")

	jobPostDTOs, total, hasNextPage, err := services.GetJobPosts(page, limit, userID, filters)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":      jobPostDTOs,
		"total":     total,
		"page":      page,
		"limit":     limit,
		"next_page": hasNextPage,
	})
}

func GetJobPostById(c *gin.Context) {
	jobID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de empleo inválido"})
		return
	}

	jobPost, err := services.GetJobPostById(uint(jobID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	jobPostDTO := dto.ConvertJobPostToDTO(*jobPost)

	c.JSON(http.StatusOK, jobPostDTO)
}

func GetEmployerJobPosts(c *gin.Context) {
	employerID, _ := c.Get("employer_id")

	jobPost, err := services.GetJobPostsByEmployerID(employerID.(uuid.UUID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, jobPost)
}

func UpdateJobPostStatus(c *gin.Context) {
	employerID, _ := c.Get("employer_id")

	jobID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de empleo inválido"})
		return
	}

	var req dto.UpdateJobPostStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	jobPost, err := services.UpdateJobPostStatus(uint(jobID), employerID.(uuid.UUID), req.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	jobPostDTO := dto.ConvertJobPostToDTO(*jobPost)

	c.JSON(http.StatusOK, gin.H{"message": "Estado del empleo actualizado", "jobPost": jobPostDTO})
}

func UpdateJobPost(c *gin.Context) {
	employerID, _ := c.Get("employer_id")

	jobID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de empleo inválido"})
		return
	}

	var req dto.UpdateJobPostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	jobPost, err := services.UpdateJobPost(uint(jobID), employerID.(uuid.UUID), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	jobPostDTO := dto.ConvertJobPostToDTO(*jobPost)

	c.JSON(http.StatusOK, gin.H{"message": "Empleo actualizado exitosamente", "jobPost": jobPostDTO})
}
