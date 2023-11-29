package main

import (
	"bytes"
	"fmt"
	"log"
	"net/smtp"
	"os"
	"text/template"

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

func sendMailHTML(subject string, templatePath string, to []string) {
	var body bytes.Buffer

	t, err := template.ParseFiles(templatePath)
	t.Execute(&body, struct{ Message string }{Message: "Hello from gbL"})

	auth := smtp.PlainAuth(
		"",
		os.Getenv("EMAIL"),
		os.Getenv("PASSWORD"),
		"smtp.gmail.com",
	)

	headers := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";"

	msg := "Subject: " + subject + "\n" + headers + "\n\n" + body.String()

	err = smtp.SendMail(
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
	// sendMailSMTP("Test",
	// 	"This is a test email",
	// 	[]string{
	// 		os.Getenv("EMAIL"),
	// 	})

	sendMailHTML("golang SMTP Test",
		"./test.html",
		[]string{
			os.Getenv("EMAIL"),
		})
}
