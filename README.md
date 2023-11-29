# Go email sender

This project is a simple prove of concept of how to send emails using Go.

--- 

## Service providers

In this project I used three different service providers to send emails:

- [x] [net/smtp](https://pkg.go.dev/net/smtp)
- [Gomail](https://pkg.go.dev/gopkg.in/gomail.v2)
- [SendGrid](https://sendgrid.com/en-us)

## How to use

### Setup

First of all, you need to create a `.env` file in the root of the project 
with the following content:

```env
EMAIL=YOUR_EMAIL
PASSWORD=YOUR_APP_PASSWORD
```

### Run

To run the project, you need to execute the following command:

```bash
go run main.go
```

