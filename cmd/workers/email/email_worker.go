package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Alvinferdeveloper/summa-backend/services"
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
		services.GlobalRabbitMQService.EmailQueue.Name, // queue
		"",    // consumer
		true,  // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			var task services.EmailTask
			if err := json.Unmarshal(d.Body, &task); err != nil {
				log.Printf("Error unmarshalling email task: %v", err)
				continue
			}

			if err := services.GlobalEmailService.SendEmail(task.To, task.Subject, task.Body); err != nil {
				log.Printf("Failed to send email to %s: %v", task.To, err)
			}
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-ctx.Done()
	log.Println("Shutting down email worker...")
}
