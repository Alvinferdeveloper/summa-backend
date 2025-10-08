package services

import (
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
		Preload("Educations.UniversitySuggestion").
		Preload("DisabilityTypes").
		Preload("AccessibilityNeeds").
		First(&profile).Error; err != nil {
		return nil, err
	}

	return &profile, nil
}

func GetAccessibilityNeeds() ([]models.AccessibilityNeed, error) {
	var accessibilityNeeds []models.AccessibilityNeed
	if err := config.DB.Find(&accessibilityNeeds).Error; err != nil {
		return nil, err
	}
	return accessibilityNeeds, nil
}

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

func SuggestNewEmployer(userID uint, req *dto.SuggestNewEmployerRequest) (*models.NewEmployer, error) {
	newEmployer := &models.NewEmployer{
		CompanyName: req.CompanyName,
		Website:     req.Website,
		SuggestedBy: userID,
		Status:      "pending",
	}

	if err := config.DB.Create(newEmployer).Error; err != nil {
		return nil, err
	}
	return newEmployer, nil
}

func GetFullProfileByID(profileID uint) (*models.Profile, error) {
	var profile models.Profile
	if err := config.DB.Where("id = ?", profileID).
		Preload("Skills").
		Preload("Experiences.Employer").
		Preload("Experiences.NewEmployer").
		Preload("Educations.University").
		Preload("Educations.UniversitySuggestion").
		Preload("DisabilityTypes").
		Preload("AccessibilityNeeds").
		First(&profile).Error; err != nil {
		return nil, err
	}

	return &profile, nil
}
