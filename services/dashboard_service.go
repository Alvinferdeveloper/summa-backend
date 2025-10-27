package services

import (
	"time"

	"github.com/Alvinferdeveloper/summa-backend/config"
	"github.com/Alvinferdeveloper/summa-backend/dto"
	"github.com/Alvinferdeveloper/summa-backend/models"
	"github.com/google/uuid"
)

func GetDashboardStats(employerID uuid.UUID) (*dto.DashboardStats, error) {
	var stats dto.DashboardStats

	if err := config.DB.Model(&models.JobPost{}).Where("employer_id = ?", employerID).Count(&stats.ActiveJobs).Error; err != nil {
		return nil, err
	}

	jobIDsSubQuery := config.DB.Model(&models.JobPost{}).Select("id").Where("employer_id = ?", employerID)

	if err := config.DB.Model(&models.JobApplication{}).Where("job_post_id IN (?)", jobIDsSubQuery).Count(&stats.TotalApplicants).Error; err != nil {
		return nil, err
	}

	sevenDaysAgo := time.Now().AddDate(0, 0, -7)
	if err := config.DB.Model(&models.JobApplication{}).Where("job_post_id IN (?) AND created_at >= ?", jobIDsSubQuery, sevenDaysAgo).Count(&stats.NewApplicants7d).Error; err != nil {
		return nil, err
	}

	return &stats, nil
}

func GetPipeline(employerID uuid.UUID) ([]dto.PipelineStep, error) {
	var pipeline []dto.PipelineStep

	jobIDsSubQuery := config.DB.Model(&models.JobPost{}).Select("id").Where("employer_id = ?", employerID)

	if err := config.DB.Model(&models.JobApplication{}).Select("status, count(*) as count").Where("job_post_id IN (?)", jobIDsSubQuery).Group("status").Order("count desc").Scan(&pipeline).Error; err != nil {
		return nil, err
	}

	return pipeline, nil
}

func GetCandidateSkillInsights(employerID uuid.UUID) ([]dto.SkillInsight, error) {
	var skills []dto.SkillInsight

	jobIDsSubQuery := config.DB.Model(&models.JobPost{}).Select("id").Where("employer_id = ?", employerID)
	profileIDsSubQuery := config.DB.Model(&models.JobApplication{}).Select("profile_id").Where("job_post_id IN (?)", jobIDsSubQuery)

	if err := config.DB.Table("skills").
		Select("skills.name, count(profile_skills.profile_id) as count").
		Joins("join profile_skills on skills.id = profile_skills.skill_id").
		Where("profile_skills.profile_id IN (?)", profileIDsSubQuery).
		Group("skills.name").
		Order("count desc").
		Limit(5).
		Scan(&skills).Error; err != nil {
		return nil, err
	}

	return skills, nil
}

func GetDisabilityInsights(employerID uuid.UUID) ([]dto.DisabilityInsight, error) {
	var insights []dto.DisabilityInsight

	jobIDsSubQuery := config.DB.Model(&models.JobPost{}).Select("id").Where("employer_id = ?", employerID)
	profileIDsSubQuery := config.DB.Model(&models.JobApplication{}).Select("profile_id").Where("job_post_id IN (?)", jobIDsSubQuery)

	if err := config.DB.Table("disability_types").
		Select("disability_types.name, count(profile_disability_types.profile_id) as count").
		Joins("join profile_disability_types on disability_types.id = profile_disability_types.disability_type_id").
		Where("profile_disability_types.profile_id IN (?)", profileIDsSubQuery).
		Group("disability_types.name").
		Order("count desc").
		Scan(&insights).Error; err != nil {
		return nil, err
	}

	return insights, nil
}
