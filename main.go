package main

import (
	"fmt"
	"log"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
)

func sendMailSMTP(subject string, message string, to []string) {
	auth := smtp.PlainAuth(
		"",
		os.Getenv("EMAIL"),
		os.Getenv("PASSWORD"),
		"smtp.gmail.com",
	)

	msg := "Subject: " + subject + "\n" + message

	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		os.Getenv("EMAIL"),
		to,
		[]byte(msg),
	)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Email sent successfully!")
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	sendMailSMTP("Test", "This is a test email", []string{
		os.Getenv("EMAIL"),
	})
}
