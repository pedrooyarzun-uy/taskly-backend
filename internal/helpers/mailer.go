package helpers

import (
	"os"

	"gopkg.in/gomail.v2"
)

func SendMail(to, subject, body string) error {
	addr := os.Getenv("EMAIL_ADDR")
	pass := os.Getenv("EMAIL_PASS")

	m := gomail.NewMessage()
	m.SetHeader("From", addr)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	d := gomail.NewDialer("smtpout.secureserver.net", 587, addr, pass)
	d.SSL = false
	return d.DialAndSend(m)
}
