package dto

import "github.com/Alvinferdeveloper/summa-backend/models"

type CandidateResponseDTO struct {
	ID                     uint   `json:"id"`
	FirstName              string `json:"first_name"`
	LastName               string `json:"last_name"`
	PhoneNumber            string `json:"phone_number"`
	City                   string `json:"city"`
	Country                string `json:"country"`
	ProfilePictureURL      string `json:"profile_picture_url"`
	BannerURL              string `json:"banner_url"`
	Address                string `json:"address"`
	LinkedIn               string `json:"linked_in"`
	ResumeURL              string `json:"resume_url"`
	Description            string `json:"description"`
	DetailedAccommodations string `json:"detailed_accommodations"`
}

func ConvertProfileToCandidateDTO(profile models.Profile) CandidateResponseDTO {
	return CandidateResponseDTO{
		ID:                     profile.ID,
		FirstName:              profile.FirstName,
		LastName:               profile.LastName,
		PhoneNumber:            profile.PhoneNumber,
		City:                   profile.City,
		Country:                profile.Country,
		ProfilePictureURL:      profile.ProfilePictureURL,
		BannerURL:              profile.BannerURL,
		Address:                profile.Address,
		LinkedIn:               profile.LinkedIn,
		ResumeURL:              profile.ResumeURL,
		Description:            profile.Description,
		DetailedAccommodations: profile.DetailedAccommodations,
	}
}
