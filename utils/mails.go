package utils

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/smtp"
	"time"
	"task/config"
)

// EmailConfig holds email sender configurations
var EmailConfig = struct {
	SMTPHost    string
	SMTPPort    string
	SenderEmail string
	SenderPass  string
}{
	SMTPHost:    config.GetEnv("SMTP_HOST", "sandbox.smtp.mailtrap.io"),  // Add a default value
	SMTPPort:    config.GetEnv("SMTP_PORT", "2525"),  // Add a default value
	SenderEmail: config.GetEnv("SMTP_USER", "e9d995191147fd"), // Add a default value
	SenderPass:  config.GetEnv("SMTP_PASS", "55e4867f39369c"),  // Add a default value
	
}

// EmailData holds dynamic email content
type EmailData struct {
	Subject string
	Message []string
	AppName string
	Year    int
}

// SendEmail sends an email with a dynamic subject and message
func SendEmail(to string, subject string, message []string) {
	// Convert message array into a formatted string
	tmpl, err := template.ParseFiles("template/email.html")
	if err != nil {
		log.Printf("Failed to load template: %v", err)
		return
	}

	// Create dynamic email data
	emailData := EmailData{
		Subject: subject,
		Message: message,
		AppName: "Task Manager", // You can fetch this dynamically if needed
		Year:    time.Now().Year(),
	}

	// Render the template with the provided data
	var body bytes.Buffer
	if err := tmpl.Execute(&body, emailData); err != nil {
		log.Printf("Failed to render template: %v", err)
		return
	}

	// SMTP authentication
	auth := smtp.PlainAuth("", EmailConfig.SenderEmail, EmailConfig.SenderPass, EmailConfig.SMTPHost)

	// Email message format (HTML)
	emailMessage := fmt.Sprintf("From: %s\nTo: %s\nSubject: %s\nMIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n%s",
		EmailConfig.SenderEmail, to, subject, body.String())

	// Send email
	addr := fmt.Sprintf("%s:%s", EmailConfig.SMTPHost, EmailConfig.SMTPPort)
	err = smtp.SendMail(addr, auth, EmailConfig.SenderEmail, []string{to}, []byte(emailMessage))
	if err != nil {
		log.Printf("Failed to send email: %v", err)
	} else {
		log.Printf("Email sent successfully to %s", to)
	}
}
