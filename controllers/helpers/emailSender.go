package helpers

import (
	"fmt"
	"log"
	"net/smtp"
	"strings"
)

func SendEmail() {
	fmt.Println("Preparing to send email...")

	// Load from env (better for security)
	from := "itgkkrrr@gmail.com"      // your Gmail address
	password := "zngk cuig ntpw qzxz" // App password from Gmail



	// Required details
	to := []string{
		"kuldeep8410mtr@gmail.com",
		"22bcs036@smvdu.ac.in",
		"gk022135@gmail.com",
	}

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	subject := "I am trying to send email using Golang"
	body := "This is a testing email sent from Golang backend developed by gkkrrr.\nThis belongs to KD Corporations."



	// Build message
	message := []byte("From: " + from + "\r\n" +
		"To: " + strings.Join(to, ",") + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" +
		body + "\r\n")




	// Authentication
	auth := smtp.PlainAuth("", from, password, smtpHost)



	// Send email
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		log.Println("Failed to send email:", err)
		return
	}

	fmt.Println("Email Sent Successfully!")
}

//kill port 3001
//fuser -k 3001/tcp

