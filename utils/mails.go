package utils

import (
	"fmt"
	"log"
	"net/smtp"
	"strings"
)

// EmailConfig holds email sender configurations
var EmailConfig = struct {
	SMTPHost    string
	SMTPPort    string
	SenderEmail string
	SenderPass  string
}{
	SMTPHost:    "sandbox.smtp.mailtrap.io", // Replace with your SMTP host
	SMTPPort:    "2525",              // Replace with your SMTP port
	SenderEmail: "e9d995191147fd", // Replace with your sender email
	SenderPass:  "55e4867f39369c",    // Replace with your sender password
}

// SendEmail sends an email with a dynamic subject and message
func SendEmail(to string, subject string, message []string) {
	// Convert message array into a formatted string
	body := strings.Join(message, "\n")

	// SMTP authentication
	auth := smtp.PlainAuth("", EmailConfig.SenderEmail, EmailConfig.SenderPass, EmailConfig.SMTPHost)

	// Email message format
	emailMessage := fmt.Sprintf("From: %s\nTo: %s\nSubject: %s\nMIME-version: 1.0;\nContent-Type: text/plain; charset=\"UTF-8\";\n\n%s",
		EmailConfig.SenderEmail, to, subject, body)

	// Send email
	addr := fmt.Sprintf("%s:%s", EmailConfig.SMTPHost, EmailConfig.SMTPPort)
	err := smtp.SendMail(addr, auth, EmailConfig.SenderEmail, []string{to}, []byte(emailMessage))
	if err != nil {
		log.Printf("Failed to send email: %v", err)
	} else {
		log.Printf("Email sent successfully to %s", to)
	}
}
