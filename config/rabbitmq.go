package config

import "os"

type RabbitMQConfig struct {
	URL string
}

func LoadRabbitMQConfig() *RabbitMQConfig {
	return &RabbitMQConfig{
		URL: os.Getenv("RABBITMQ_URL"),
	}
}
