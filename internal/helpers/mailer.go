package helpers

import (
	"net/smtp"
	"os"
)

func SendMail(to, subject, body string) error {
	addr := os.Getenv("EMAIL_ADDR")
	pass := os.Getenv("EMAIL_PASS")

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	msg := []byte("To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"MIME-version: 1.0;\r\n" +
		"Content-Type: text/html; charset=\"UTF-8\";\r\n\r\n" +
		body + "\r\n")

	auth := smtp.PlainAuth("", addr, pass, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, addr, []string{to}, msg)
	return err
}
