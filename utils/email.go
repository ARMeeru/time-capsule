package utils

import (
	"crypto/tls"
	"strconv"

	"gopkg.in/gomail.v2"

	"github.com/ARMeeru/time-capsule/config"
)

func SendEmail(to string, subject string, body string) error {
	mailer := gomail.NewMessage()

	// Use a valid 'From' email address
	fromEmail := config.GetEnv("EMAIL_FROM")
	if fromEmail == "" {
		fromEmail = "no-reply@example.com" // Default value if not set
	}

	mailer.SetHeader("From", fromEmail)
	mailer.SetHeader("To", to)
	mailer.SetHeader("Subject", subject)
	mailer.SetBody("text/plain", body)

	port, _ := strconv.Atoi(config.GetEnv("EMAIL_PORT"))

	dialer := gomail.NewDialer(
		config.GetEnv("EMAIL_HOST"),
		port,
		config.GetEnv("EMAIL_USER"),
		config.GetEnv("EMAIL_PASSWORD"),
	)

	// Skip TLS verification (for development only)
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	err := dialer.DialAndSend(mailer)
	return err
}
