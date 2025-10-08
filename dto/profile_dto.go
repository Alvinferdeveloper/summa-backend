package dto

import (
	"time"

	"github.com/Alvinferdeveloper/summa-backend/models"
)

// --- Request DTOs ---

type UpdatePersonalInfoRequest struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
}

type UpdateContactInfoRequest struct {
	PhoneNumber    string `json:"phone_number"`
	City           string `json:"city"`
	Country        string `json:"country"`
	Address        string `json:"address"`
	LinkedIn       string `json:"linked_in"`
	ResumeURL      string `json:"resume_url"`
	ProfilePicture string `json:"profile_picture"`
}

type UpdateDescriptionRequest struct {
	Description string `json:"description" binding:"required"`
}

type UpdateDisabilityInfoRequest struct {
	DisabilityInfoConsent  bool   `json:"disability_info_consent"`
	DetailedAccommodations string `json:"detailed_accommodations"`
	DisabilityTypeIDs      []uint `json:"disability_type_ids"`
	AccessibilityNeedIDs   []uint `json:"accessibility_need_ids"`
}

type CreateEducationRequest struct {
	Degree                 string     `json:"degree" binding:"required"`
	FieldOfStudy           string     `json:"field_of_study"`
	StartDate              time.Time  `json:"start_date" binding:"required"`
	EndDate                *time.Time `json:"end_date"`
	UniversityID           *uint      `json:"university_id"`
	UniversitySuggestionID *uint      `json:"university_suggestion_id"`
}

type UpdateEducationRequest struct {
	Degree                 string     `json:"degree" binding:"required"`
	FieldOfStudy           string     `json:"field_of_study"`
	StartDate              time.Time  `json:"start_date" binding:"required"`
	EndDate                *time.Time `json:"end_date"`
	UniversityID           *uint      `json:"university_id"`
	UniversitySuggestionID *uint      `json:"university_suggestion_id"`
}

type UpdateSkillsRequest struct {
	SkillNames []string `json:"skill_names" binding:"required"`
}

type SuggestNewEmployerRequest struct {
	CompanyName string `json:"company_name" binding:"required"`
	Website     string `json:"website"`
}

// --- Response DTOs ---

// FullProfileResponseDTO representa el perfil completo de un usuario para la API.
type FullProfileResponseDTO struct {
	ID                     uint                        `json:"id"`
	FirstName              string                      `json:"first_name"`
	LastName               string                      `json:"last_name"`
	PhoneNumber            string                      `json:"phone_number"`
	City                   string                      `json:"city"`
	Country                string                      `json:"country"`
	ProfilePicture         string                      `json:"profile_picture"`
	Address                string                      `json:"address"`
	LinkedIn               string                      `json:"linked_in"`
	ResumeURL              string                      `json:"resume_url"`
	Description            string                      `json:"description"`
	DisabilityInfoConsent  bool                        `json:"disability_info_consent"`
	DetailedAccommodations string                      `json:"detailed_accommodations"`
	Experiences            []ExperienceResponse        `json:"experiences"`
	Educations             []EducationResponse         `json:"educations"`
	Skills                 []SkillResponse             `json:"skills"`
	DisabilityTypes        []DisabilityTypeResponse    `json:"disability_types"`
	AccessibilityNeeds     []AccessibilityNeedResponse `json:"accessibility_needs"`
}

type ExperienceResponse struct {
	ID          uint                        `json:"id"`
	JobTitle    string                      `json:"job_title"`
	Description string                      `json:"description"`
	StartDate   time.Time                   `json:"start_date"`
	EndDate     *time.Time                  `json:"end_date"`
	Employer    *EmployerResponseDTO        `json:"employer,omitempty"`
	NewEmployer *NewEmployerSummaryResponse `json:"new_employer,omitempty"`
}
type EducationResponse struct {
	ID                   uint                                 `json:"id"`
	Degree               string                               `json:"degree"`
	FieldOfStudy         string                               `json:"field_of_study"`
	StartDate            time.Time                            `json:"start_date"`
	EndDate              *time.Time                           `json:"end_date"`
	University           *UniversitySummaryResponse           `json:"university,omitempty"`
	UniversitySuggestion *UniversitySuggestionSummaryResponse `json:"university_suggestion,omitempty"`
}

type SkillResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
type DisabilityTypeResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
type AccessibilityNeedResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type NewEmployerSummaryResponse struct {
	ID          uint   `json:"id"`
	CompanyName string `json:"company_name"`
}
type UniversitySummaryResponse struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}
type UniversitySuggestionSummaryResponse struct {
	ID            uint   `json:"id"`
	SuggestedName string `json:"suggested_name"`
}

func ConvertProfileToFullDTO(profile models.Profile) FullProfileResponseDTO {
	// Convert Experiences
	experiences := make([]ExperienceResponse, len(profile.Experiences))
	for i, exp := range profile.Experiences {
		experiences[i] = ExperienceResponse{
			ID:          exp.ID,
			JobTitle:    exp.JobTitle,
			Description: exp.Description,
			StartDate:   exp.StartDate,
			EndDate:     exp.EndDate,
		}
		if exp.Employer != nil {
			experiences[i].Employer = &EmployerResponseDTO{
				ID:          exp.Employer.ID,
				CompanyName: exp.Employer.CompanyName,
				LogoURL:     exp.Employer.LogoURL,
			}
		}
		if exp.NewEmployer != nil {
			experiences[i].NewEmployer = &NewEmployerSummaryResponse{
				ID:          exp.NewEmployer.ID,
				CompanyName: exp.NewEmployer.CompanyName,
			}
		}
	}

	educations := make([]EducationResponse, len(profile.Educations))
	for i, edu := range profile.Educations {
		educations[i] = EducationResponse{
			ID:           edu.ID,
			Degree:       edu.Degree,
			FieldOfStudy: edu.FieldOfStudy,
			StartDate:    edu.StartDate,
			EndDate:      edu.EndDate,
		}
		if edu.University != nil {
			educations[i].University = &UniversitySummaryResponse{
				ID:      edu.University.ID,
				Name:    edu.University.Name,
				Address: edu.University.Address,
			}
		}
		if edu.UniversitySuggestion != nil {
			educations[i].UniversitySuggestion = &UniversitySuggestionSummaryResponse{
				ID:            edu.UniversitySuggestion.ID,
				SuggestedName: edu.UniversitySuggestion.SuggestedName,
			}
		}
	}

	skills := make([]SkillResponse, len(profile.Skills))
	for i, skill := range profile.Skills {
		skills[i] = SkillResponse{ID: skill.ID, Name: skill.Name}
	}

	disabilityTypes := make([]DisabilityTypeResponse, len(profile.DisabilityTypes))
	for i, dt := range profile.DisabilityTypes {
		disabilityTypes[i] = DisabilityTypeResponse{ID: dt.ID, Name: dt.Name}
	}
	accessibilityNeeds := make([]AccessibilityNeedResponse, len(profile.AccessibilityNeeds))
	for i, an := range profile.AccessibilityNeeds {
		accessibilityNeeds[i] = AccessibilityNeedResponse{ID: an.ID, Name: an.Name}
	}

	return FullProfileResponseDTO{
		ID:                     profile.ID,
		FirstName:              profile.FirstName,
		LastName:               profile.LastName,
		PhoneNumber:            profile.PhoneNumber,
		City:                   profile.City,
		Country:                profile.Country,
		ProfilePicture:         profile.ProfilePicture,
		Address:                profile.Address,
		LinkedIn:               profile.LinkedIn,
		ResumeURL:              profile.ResumeURL,
		Description:            profile.Description,
		DisabilityInfoConsent:  profile.DisabilityInfoConsent,
		DetailedAccommodations: profile.DetailedAccommodations,
		Experiences:            experiences,
		Educations:             educations,
		Skills:                 skills,
		DisabilityTypes:        disabilityTypes,
		AccessibilityNeeds:     accessibilityNeeds,
	}
}
