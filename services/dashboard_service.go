package services

import (
	"time"

	"github.com/Alvinferdeveloper/summa-backend/config"
	"github.com/Alvinferdeveloper/summa-backend/models"
	"github.com/google/uuid"
)

type DashboardStats struct {
	ActiveJobs      int64 `json:"active_jobs"`
	TotalApplicants int64 `json:"total_applicants"`
	NewApplicants7d int64 `json:"new_applicants_7d"`
}

type PipelineStep struct {
	Status string `json:"status"`
	Count  int64  `json:"count"`
}

type SkillInsight struct {
	Name  string `json:"name"`
	Count int64  `json:"count"`
}

func GetDashboardStats(employerID uuid.UUID) (*DashboardStats, error) {
	var stats DashboardStats

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

func GetPipeline(employerID uuid.UUID) ([]PipelineStep, error) {
	var pipeline []PipelineStep

	jobIDsSubQuery := config.DB.Model(&models.JobPost{}).Select("id").Where("employer_id = ?", employerID)

	if err := config.DB.Model(&models.JobApplication{}).Select("status, count(*) as count").Where("job_post_id IN (?)", jobIDsSubQuery).Group("status").Order("count desc").Scan(&pipeline).Error; err != nil {
		return nil, err
	}

	return pipeline, nil
}

func GetCandidateSkillInsights(employerID uuid.UUID) ([]SkillInsight, error) {
	var skills []SkillInsight

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
