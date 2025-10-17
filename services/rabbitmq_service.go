package services

import (
	"context"
	"encoding/json"
	"time"

	"github.com/Alvinferdeveloper/summa-backend/config"
	"github.com/rabbitmq/amqp091-go"
)

type RabbitMQService struct {
	Conn    *amqp091.Connection
	Channel *amqp091.Channel
	Queue   amqp091.Queue
}

var GlobalRabbitMQService *RabbitMQService

func InitRabbitMQService() error {
	rabbitMQConfig := config.LoadRabbitMQConfig()

	conn, err := amqp091.Dial(rabbitMQConfig.URL)
	if err != nil {
		return err
	}

	ch, err := conn.Channel()
	if err != nil {
		return err
	}

	q, err := ch.QueueDeclare(
		"email_queue", // name
		true,          // durable
		false,         // delete when unused
		false,         // exclusive
		false,         // no-wait
		nil,           // arguments
	)
	if err != nil {
		return err
	}

	GlobalRabbitMQService = &RabbitMQService{
		Conn:    conn,
		Channel: ch,
		Queue:   q,
	}

	return nil
}

type EmailTask struct {
	To      string `json:"to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

func (s *RabbitMQService) PublishEmailTask(task *EmailTask) error {
	body, err := json.Marshal(task)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return s.Channel.PublishWithContext(ctx,
		"",           // exchange
		s.Queue.Name, // routing key
		false,        // mandatory
		false,        // immediate
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
}
