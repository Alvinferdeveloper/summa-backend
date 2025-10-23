package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Alvinferdeveloper/summa-backend/config"
	"github.com/Alvinferdeveloper/summa-backend/models"
	"github.com/Alvinferdeveloper/summa-backend/services"
	"github.com/Alvinferdeveloper/summa-backend/utils"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file, using environment variables")
	}

	if err := services.InitEmailService(); err != nil {
		log.Fatalf("Failed to initialize email service: %v", err)
	}

	if err := services.InitRabbitMQService(); err != nil {
		log.Fatalf("Failed to initialize RabbitMQ service: %v", err)
	}
	msgs, err := services.GlobalRabbitMQService.Channel.Consume(
		services.GlobalRabbitMQService.ChatNotificationQueue.Name,
		"",    // consumer
		false, // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %s", err)
	}
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	go func() {
		for d := range msgs {
			log.Printf("Received a chat notification task: %s", d.Body)

			var task services.ChatMessageNotificationTask
			if err := json.Unmarshal(d.Body, &task); err != nil {
				log.Printf("Error unmarshalling chat notification task: %s", err)
				d.Nack(false, false) // Discard the message
				continue
			}

			var message models.Message
			if err := json.Unmarshal(task.OriginalMessage, &message); err != nil {
				log.Printf("Error unmarshalling original message: %s", err)
				d.Nack(false, false)
				continue
			}

			if err := processNotification(&message); err != nil {
				log.Printf("Error processing chat notification: %s", err)
				d.Nack(false, true) // Requeue the message
			} else {
				d.Ack(false) // Acknowledge the message
			}
		}
	}()
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-ctx.Done()
	log.Println("Shutting down chat notification worker...")
}

func processNotification(message *models.Message) error {
	var recipientEmail, recipientName, senderName string
	var err error

	// Get Recipient Info
	if message.RecipientType == "user" {
		var user models.User
		if err = config.DB.Preload("Profile").First(&user, message.RecipientID).Error; err == nil {
			recipientEmail = user.Email
			recipientName = user.Profile.FirstName
		} else {
			return fmt.Errorf("could not find recipient user with ID %d: %w", message.RecipientID, err)
		}
	} else if message.RecipientType == "employer" {
		var employer models.Employer
		if err = config.DB.First(&employer, message.RecipientID).Error; err == nil {
			recipientEmail = employer.Email
			recipientName = employer.CompanyName
		} else {
			return fmt.Errorf("could not find recipient employer with ID %d: %w", message.RecipientID, err)
		}
	} else {
		return fmt.Errorf("invalid recipient type: %s", message.RecipientType)
	}

	// Get Sender Info
	if message.SenderType == "user" {
		var user models.User
		if err = config.DB.Preload("Profile").First(&user, message.SenderID).Error; err == nil {
			senderName = user.Profile.FirstName
		} else {
			return fmt.Errorf("could not find sender user with ID %d: %w", message.SenderID, err)
		}
	} else if message.SenderType == "employer" {
		var employer models.Employer
		if err = config.DB.First(&employer, message.SenderID).Error; err == nil {
			senderName = employer.CompanyName
		} else {
			return fmt.Errorf("could not find sender employer with ID %d: %w", message.SenderID, err)
		}
	} else {
		return fmt.Errorf("invalid sender type: %s", message.SenderType)
	}

	// Prepare and send email
	data := struct {
		RecipientName string
		SenderName    string
	}{
		RecipientName: recipientName,
		SenderName:    senderName,
	}

	body, err := utils.ParseTemplate("new_chat_message.html", data)
	if err != nil {
		return fmt.Errorf("failed to render email template: %w", err)
	}

	return services.GlobalEmailService.SendEmail(recipientEmail, "You Have a New Message on Summa", body)
}
