package dto

type NewUniversitySuggestionRequest struct {
	SuggestedName string `json:"suggested_name" binding:"required"`
	Country       string `json:"country"`
	Website       string `json:"website"`
}
