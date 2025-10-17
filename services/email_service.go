package services

import (
	"github.com/Alvinferdeveloper/summa-backend/config"
	"gopkg.in/gomail.v2"
)

type EmailService struct {
	Mailer *gomail.Dialer
}

var GlobalEmailService *EmailService

func InitEmailService() error {
	mailConfig, err := config.LoadMailConfig()
	if err != nil {
		return err
	}

	dialer := gomail.NewDialer(mailConfig.SMTPHost, mailConfig.SMTPPort, mailConfig.SMTPUser, mailConfig.SMTPPass)

	GlobalEmailService = &EmailService{
		Mailer: dialer,
	}

	return nil
}

func (s *EmailService) SendEmail(to, subject, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", s.Mailer.Username)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	return s.Mailer.DialAndSend(m)
}
