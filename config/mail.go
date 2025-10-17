package config

import (
	"os"
	"strconv"
)

type MailConfig struct {
	SMTPHost string
	SMTPPort int
	SMTPUser string
	SMTPPass string
}

func LoadMailConfig() (*MailConfig, error) {
	port, err := strconv.Atoi(os.Getenv("SMTP_PORT"))
	if err != nil {
		return nil, err
	}

	return &MailConfig{
		SMTPHost: os.Getenv("SMTP_HOST"),
		SMTPPort: port,
		SMTPUser: os.Getenv("SMTP_USER"),
		SMTPPass: os.Getenv("GOOGLE_APP_PASSWORD"),
	}, nil
}
