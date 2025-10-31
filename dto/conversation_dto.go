package dto

import (
	"github.com/Alvinferdeveloper/summa-backend/models"
	"github.com/google/uuid"
)

type UserSummary struct {
	ID                uuid.UUID `json:"id"`
	FirstName         string    `json:"first_name"`
	LastName          string    `json:"last_name"`
	Email             string    `json:"email"`
	ProfilePictureURL string    `json:"profile_picture_url"`
}

type EmployerSummary struct {
	ID          uuid.UUID `json:"id"`
	CompanyName string    `json:"company_name"`
	Email       string    `json:"email"`
	LogoURL     string    `json:"logo_url"`
}
type ConversationResponseDTO struct {
	ID         uint            `json:"id"`
	EmployerID uuid.UUID       `json:"employer_id"`
	User       UserSummary     `json:"user"`
	Employer   EmployerSummary `json:"employer"`
	UnreadCount int64           `json:"unread_count"`
}

func ConvertConversationToDTO(conversation models.Conversation, unreadCount int64) *ConversationResponseDTO {
	return &ConversationResponseDTO{
		ID:         conversation.ID,
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
		UnreadCount: unreadCount,
	}
}
