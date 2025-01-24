package services

import (
	"fmt"

	"crazygames.io/config"
	"github.com/go-gomail/gomail"
)

type EmailService struct{}

func NewEmailService() *EmailService {
	return &EmailService{}
}

func (s *EmailService) SendEmail(to string, subject string, body string) error {
	mail := gomail.NewMessage()
	mail.SetHeader("From", config.SMTP.Username)
	mail.SetHeader("To", to)
	mail.SetHeader("Subject", subject)
	mail.SetBody("text/html", body)

	dialer := gomail.NewDialer(config.SMTP.Host, config.SMTP.Port, config.SMTP.Username, config.SMTP.Password)
	if err := dialer.DialAndSend(mail); err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}
	return nil
}
