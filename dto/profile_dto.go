package dto

import "time"

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
	Description string `json:"description"`
}

type UpdateDisabilityInfoRequest struct {
	DisabilityInfoConsent  bool   `json:"disability_info_consent"`
	DetailedAccommodations string `json:"detailed_accommodations"`
	DisabilityTypeIDs      []uint `json:"disability_type_ids"`
	AccessibilityNeedIDs   []uint `json:"accessibility_need_ids"`
}

type UpdateExperienceRequest struct {
	EmployerID    *uint      `json:"employer_id"`
	NewEmployerID *uint      `json:"new_employer_id"`
	JobTitle      string     `json:"job_title" binding:"required"`
	Description   string     `json:"description"`
	StartDate     time.Time  `json:"start_date" binding:"required"`
	EndDate       *time.Time `json:"end_date"`
}

type CreateEducationRequest struct {
	UniversityID           *uint      `json:"university_id"`
	UniversitySuggestionID *uint      `json:"university_suggestion_id"`
	Degree                 string     `json:"degree" binding:"required"`
	FieldOfStudy           string     `json:"field_of_study"`
	StartDate              time.Time  `json:"start_date" binding:"required"`
	EndDate                *time.Time `json:"end_date"`
}

type UpdateEducationRequest struct {
	UniversityID           *uint      `json:"university_id"`
	UniversitySuggestionID *uint      `json:"university_suggestion_id"`
	Degree                 string     `json:"degree" binding:"required"`
	FieldOfStudy           string     `json:"field_of_study"`
	StartDate              time.Time  `json:"start_date" binding:"required"`
	EndDate                *time.Time `json:"end_date"`
}

type UpdateSkillsRequest struct {
	SkillNames []string `json:"skill_names" binding:"required,min=1"`
}

// SuggestNewEmployerRequest defines the data for suggesting a new employer.
type SuggestNewEmployerRequest struct {
	CompanyName string `json:"company_name" binding:"required"`
	Website     string `json:"website"`
}
