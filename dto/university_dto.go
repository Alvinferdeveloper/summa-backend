package dto

import "github.com/Alvinferdeveloper/summa-backend/models"

type UniversityResponse struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Country string `json:"country"`
	Website string `json:"website"`
	LogoURL string `json:"logo_url"`
	Address string `json:"address"`
}

func ConvertUniversityToDTO(university models.University) UniversityResponse {
	return UniversityResponse{
		ID:      university.ID,
		Name:    university.Name,
		Country: university.Country,
		Website: university.Website,
		LogoURL: university.LogoURL,
		Address: university.Address,
	}
}
