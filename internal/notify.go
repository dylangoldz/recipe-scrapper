package internal

import (
	"fmt"
	"net/smtp"
	"os"
)

// EmailNotifier represents an email notification service using Google SMTP
type EmailNotifier struct {
	from     string
	password string
	smtpHost string
	smtpPort string
}

// NewEmailNotifier creates a new EmailNotifier instance
func NewEmailNotifier() (*EmailNotifier, error) {
	from := os.Getenv("EMAIL_FROM")
	password := os.Getenv("EMAIL_PASSWORD")
	if from == "" || password == "" {
		return nil, fmt.Errorf("EMAIL_FROM and EMAIL_PASSWORD environment variables must be set")
	}

	return &EmailNotifier{
		from:     from,
		password: password,
		smtpHost: "smtp.gmail.com",
		smtpPort: "587",
	}, nil
}

// SendEmail sends an email notification
func (e *EmailNotifier) SendEmail(to, subject, body string) error {
	auth := smtp.PlainAuth("", e.from, e.password, e.smtpHost)

	msg := []byte(fmt.Sprintf("To: %s\r\n"+
		"Subject: %s\r\n"+
		"\r\n"+
		"%s\r\n", to, subject, body))

	err := smtp.SendMail(e.smtpHost+":"+e.smtpPort, auth, e.from, []string{to}, msg)
	if err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}

	return nil
}
