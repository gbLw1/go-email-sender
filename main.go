package main

import (
	"fmt"
	"log"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
)

func sendMailSMTP() {
	auth := smtp.PlainAuth(
		"",
		os.Getenv("EMAIL"),
		os.Getenv("PASSWORD"),
		"smtp.gmail.com",
	)

	msg := "Subject: Hello there!\nThis is the email body."

	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		os.Getenv("EMAIL"),
		[]string{
			os.Getenv("EMAIL"),
		},
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
	sendMailSMTP()
}
