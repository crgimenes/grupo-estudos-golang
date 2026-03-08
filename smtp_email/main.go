package main

import (
	"fmt"
	"net/smtp"
)

func BuildMessage(from, to, subject, body string) []byte {
	msg := "From: " + from + "\r\n" +
		"To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" + body + "\r\n"
	return []byte(msg)
}

func SMTPAuth(user, pass, host string) smtp.Auth {
	return smtp.PlainAuth("", user, pass, host)
}

func main() {
	msg := BuildMessage("noreply@example.com", "team@example.com", "status report", "all services are up")
	auth := SMTPAuth("api-user", "secret", "smtp.example.com")

	fmt.Printf("auth=%T\n", auth)
	fmt.Println(string(msg))
}
