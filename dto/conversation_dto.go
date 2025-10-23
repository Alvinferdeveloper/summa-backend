package dto

import "github.com/Alvinferdeveloper/summa-backend/models"

type UserSummary struct {
	ID                uint   `json:"id"`
	FirstName         string `json:"first_name"`
	LastName          string `json:"last_name"`
	Email             string `json:"email"`
	ProfilePictureURL string `json:"profile_picture_url"`
}

type EmployerSummary struct {
	ID          uint   `json:"id"`
	CompanyName string `json:"company_name"`
	Email       string `json:"email"`
	LogoURL     string `json:"logo_url"`
}
type ConversationResponseDTO struct {
	ID         uint            `json:"id"`
	UserID     uint            `json:"user_id"`
	EmployerID uint            `json:"employer_id"`
	User       UserSummary     `json:"user"`
	Employer   EmployerSummary `json:"employer"`
}

func ConvertConversationToDTO(conversation models.Conversation) *ConversationResponseDTO {
	return &ConversationResponseDTO{
		ID:         conversation.ID,
		UserID:     conversation.UserID,
		EmployerID: conversation.EmployerID,
		User: UserSummary{
			ID:                conversation.User.ID,
			FirstName:         conversation.User.Profile.FirstName,
			LastName:          conversation.User.Profile.LastName,
			Email:             conversation.User.Email,
			ProfilePictureURL: conversation.User.Profile.ProfilePictureURL,
		},
		Employer: EmployerSummary{
			ID:          conversation.Employer.ID,
			CompanyName: conversation.Employer.CompanyName,
			Email:       conversation.Employer.Email,
			LogoURL:     conversation.Employer.LogoURL,
		},
	}
}
