package services

import (
	"fmt"

	"github.com/Alvinferdeveloper/summa-backend/config"
	"github.com/Alvinferdeveloper/summa-backend/dto"
	"github.com/Alvinferdeveloper/summa-backend/models"
)

func CompleteOnboarding(req *dto.OnboardingRequest, userID uint) (*models.Profile, error) {
	var profile models.Profile
	if err := config.DB.Where("user_id = ?", userID).First(&profile).Error; err != nil {
		return nil, err
	}

	profile.FirstName = req.FirstName
	profile.LastName = req.LastName
	profile.OnboardingCompleted = true

	var disabilityTypes []models.DisabilityType
	if err := config.DB.Where(req.DisabilityTypeIDs).Find(&disabilityTypes).Error; err != nil {
		return nil, err
	}

	if err := config.DB.Model(&profile).Association("DisabilityTypes").Replace(&disabilityTypes); err != nil {
		return nil, err
	}

	if err := config.DB.Save(&profile).Error; err != nil {
		return nil, err
	}

	return &profile, nil
}

func GetDisabilityTypes() ([]models.DisabilityType, error) {
	var disabilityTypes []models.DisabilityType
	if err := config.DB.Find(&disabilityTypes).Error; err != nil {
		return nil, err
	}
	return disabilityTypes, nil
}

func GetFullProfile(userID uint) (*models.Profile, error) {
	var profile models.Profile
	if err := config.DB.Where("user_id = ?", userID).
		Preload("Skills").
		Preload("Experiences.Employer").
		Preload("Experiences.NewEmployer").
		Preload("Educations.University").
		Preload("DisabilityTypes").
		Preload("AccessibilityNeeds").
		First(&profile).Error; err != nil {
		return nil, err
	}
	return &profile, nil
}

// GetAccessibilityNeeds fetches all available accessibility needs.
func GetAccessibilityNeeds() ([]models.AccessibilityNeed, error) {
	var accessibilityNeeds []models.AccessibilityNeed
	if err := config.DB.Find(&accessibilityNeeds).Error; err != nil {
		return nil, err
	}
	return accessibilityNeeds, nil
}

// UpdatePersonalInfo updates the basic personal information of a profile.
func UpdatePersonalInfo(userID uint, req *dto.UpdatePersonalInfoRequest) (*models.Profile, error) {
	var profile models.Profile
	if err := config.DB.Where("user_id = ?", userID).First(&profile).Error; err != nil {
		return nil, err
	}
	profile.FirstName = req.FirstName
	profile.LastName = req.LastName
	if err := config.DB.Save(&profile).Error; err != nil {
		return nil, err
	}
	return &profile, nil
}

// UpdateContactInfo updates the contact information of a profile.
func UpdateContactInfo(userID uint, req *dto.UpdateContactInfoRequest) (*models.Profile, error) {
	var profile models.Profile
	if err := config.DB.Where("user_id = ?", userID).First(&profile).Error; err != nil {
		return nil, err
	}
	profile.PhoneNumber = req.PhoneNumber
	profile.City = req.City
	profile.Country = req.Country
	profile.Address = req.Address
	profile.LinkedIn = req.LinkedIn
	profile.ResumeURL = req.ResumeURL
	profile.ProfilePicture = req.ProfilePicture
	if err := config.DB.Save(&profile).Error; err != nil {
		return nil, err
	}
	return &profile, nil
}

// UpdateDescription updates the personal description of a profile.
func UpdateDescription(userID uint, req *dto.UpdateDescriptionRequest) (*models.Profile, error) {
	var profile models.Profile
	if err := config.DB.Where("user_id = ?", userID).First(&profile).Error; err != nil {
		return nil, err
	}
	profile.Description = req.Description
	if err := config.DB.Save(&profile).Error; err != nil {
		return nil, err
	}
	return &profile, nil
}

// UpdateDisabilityInfo updates disability-related information and associations.
func UpdateDisabilityInfo(userID uint, req *dto.UpdateDisabilityInfoRequest) (*models.Profile, error) {
	var profile models.Profile
	if err := config.DB.Where("user_id = ?", userID).First(&profile).Error; err != nil {
		return nil, err
	}
	profile.DisabilityInfoConsent = req.DisabilityInfoConsent
	profile.DetailedAccommodations = req.DetailedAccommodations

	// Update DisabilityTypes association
	var disabilityTypes []models.DisabilityType
	if len(req.DisabilityTypeIDs) > 0 {
		if err := config.DB.Where(req.DisabilityTypeIDs).Find(&disabilityTypes).Error; err != nil {
			return nil, err
		}
	}
	if err := config.DB.Model(&profile).Association("DisabilityTypes").Replace(&disabilityTypes); err != nil {
		return nil, err
	}

	// Update AccessibilityNeeds association
	var accessibilityNeeds []models.AccessibilityNeed
	if len(req.AccessibilityNeedIDs) > 0 {
		if err := config.DB.Where(req.AccessibilityNeedIDs).Find(&accessibilityNeeds).Error; err != nil {
			return nil, err
		}
	}
	if err := config.DB.Model(&profile).Association("AccessibilityNeeds").Replace(&accessibilityNeeds); err != nil {
		return nil, err
	}

	if err := config.DB.Save(&profile).Error; err != nil {
		return nil, err
	}
	return &profile, nil
}

// CreateExperience creates a new experience for a profile.
func CreateExperience(profileID uint, req *dto.CreateExperienceRequest) (*models.Experience, error) {
	fmt.Println(req)
	experience := &models.Experience{
		ProfileID:     profileID,
		EmployerID:    req.EmployerID,
		NewEmployerID: req.NewEmployerID,
		JobTitle:      req.JobTitle,
		Description:   req.Description,
		StartDate:     req.StartDate,
		EndDate:       req.EndDate,
	}
	if err := config.DB.Create(experience).Error; err != nil {
		return nil, err
	}
	return experience, nil
}

// UpdateExperience updates an existing experience for a profile.
func UpdateExperience(profileID uint, experienceID uint, req *dto.UpdateExperienceRequest) (*models.Experience, error) {
	var experience models.Experience
	if err := config.DB.Where("profile_id = ? AND id = ?", profileID, experienceID).First(&experience).Error; err != nil {
		return nil, err
	}
	experience.EmployerID = req.EmployerID
	experience.NewEmployerID = req.NewEmployerID
	experience.JobTitle = req.JobTitle
	experience.Description = req.Description
	experience.StartDate = req.StartDate
	experience.EndDate = req.EndDate
	if err := config.DB.Save(&experience).Error; err != nil {
		return nil, err
	}
	return &experience, nil
}

// DeleteExperience deletes an experience from a profile.
func DeleteExperience(profileID uint, experienceID uint) error {
	return config.DB.Where("profile_id = ?", profileID).Delete(&models.Experience{}, experienceID).Error
}

// CreateEducation creates a new education entry for a profile.
func CreateEducation(profileID uint, req *dto.CreateEducationRequest) (*models.ProfileEducation, error) {
	education := &models.ProfileEducation{
		ProfileID:              profileID,
		UniversityID:           req.UniversityID,
		UniversitySuggestionID: req.UniversitySuggestionID,
		Degree:                 req.Degree,
		FieldOfStudy:           req.FieldOfStudy,
		StartDate:              req.StartDate,
		EndDate:                req.EndDate,
	}
	if err := config.DB.Create(education).Error; err != nil {
		return nil, err
	}
	return education, nil
}

// UpdateEducation updates an existing education entry for a profile.
func UpdateEducation(profileID uint, educationID uint, req *dto.UpdateEducationRequest) (*models.ProfileEducation, error) {
	var education models.ProfileEducation
	if err := config.DB.Where("profile_id = ? AND id = ?", profileID, educationID).First(&education).Error; err != nil {
		return nil, err
	}
	education.UniversityID = req.UniversityID
	education.UniversitySuggestionID = req.UniversitySuggestionID
	education.Degree = req.Degree
	education.FieldOfStudy = req.FieldOfStudy
	education.StartDate = req.StartDate
	education.EndDate = req.EndDate
	if err := config.DB.Save(&education).Error; err != nil {
		return nil, err
	}
	return &education, nil
}

func DeleteEducation(profileID uint, educationID uint) error {
	return config.DB.Where("profile_id = ?", profileID).Delete(&models.ProfileEducation{}, educationID).Error
}
func UpdateSkills(profileID uint, req *dto.UpdateSkillsRequest) (*models.Profile, error) {
	var profile models.Profile
	if err := config.DB.Where("user_id = ?", profileID).First(&profile).Error; err != nil {
		return nil, err
	}

	var skills []models.Skill
	// Find existing skills by name or create new ones
	for _, skillName := range req.SkillNames {
		var skill models.Skill
		if err := config.DB.FirstOrCreate(&skill, models.Skill{Name: skillName}).Error; err != nil {
			return nil, err
		}
		skills = append(skills, skill)
	}

	// Replace existing skills association
	if err := config.DB.Model(&profile).Association("Skills").Replace(&skills); err != nil {
		return nil, err
	}

	if err := config.DB.Save(&profile).Error; err != nil {
		return nil, err
	}
	return &profile, nil
}
