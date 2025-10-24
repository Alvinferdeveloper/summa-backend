package services

import (
	"fmt"

	"github.com/Alvinferdeveloper/summa-backend/config"
	"github.com/Alvinferdeveloper/summa-backend/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetOrCreateConversation(userID, employerID uuid.UUID) (*models.Conversation, error) {
	var conversation models.Conversation

	err := config.DB.Where("user_id = ? AND employer_id = ?", userID, employerID).First(&conversation).Error

	if err == gorm.ErrRecordNotFound {
		newConversation := models.Conversation{
			UserID:     userID,
			EmployerID: employerID,
		}
		if err = config.DB.Create(&newConversation).Error; err != nil {
			return nil, err
		}
		// Preload the associations for the new conversation
		config.DB.Preload("User.Profile").Preload("Employer").First(&newConversation, newConversation.ID)
		return &newConversation, nil
	} else if err != nil {
		return nil, err
	}

	// Preload associations for existing conversation
	config.DB.Preload("User.Profile").Preload("Employer").First(&conversation, conversation.ID)
	return &conversation, nil
}

func GetConversations(participantID uuid.UUID, participantType string) ([]models.Conversation, error) {
	var conversations []models.Conversation
	var query *gorm.DB

	switch participantType {
	case "user":
		query = config.DB.Where("user_id = ?", participantID)
	case "employer":
		query = config.DB.Where("employer_id = ?", participantID)
	default:
		return nil, fmt.Errorf("invalid participant type")
	}

	err := query.Preload("User.Profile").Preload("Employer").Order("updated_at desc").Find(&conversations).Error
	return conversations, err
}

func GetMessagesForConversation(conversationID uint, page, limit int) ([]models.Message, int64, error) {
	var messages []models.Message
	var total int64

	offset := (page - 1) * limit

	query := config.DB.Model(&models.Message{}).Where("conversation_id = ?", conversationID)

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Order("created_at desc").Limit(limit).Offset(offset).Find(&messages).Error; err != nil {
		return nil, 0, err
	}

	return messages, total, nil
}

func CreateMessage(conversationID uint, senderID uuid.UUID, senderType string, recipientID uuid.UUID, recipientType string, content string) (*models.Message, error) {
	message := models.Message{
		ConversationID: conversationID,
		SenderID:       senderID,
		SenderType:     senderType,
		RecipientID:    recipientID,
		RecipientType:  recipientType,
		Content:        content,
		Read:           false,
	}

	if err := config.DB.Create(&message).Error; err != nil {
		return nil, err
	}

	if err := config.DB.Model(&models.Conversation{}).Where("id = ?", conversationID).Update("updated_at", gorm.Expr("NOW()")).Error; err != nil {
		return nil, err
	}

	return &message, nil
}
