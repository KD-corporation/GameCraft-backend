package helpers

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
	"strings"
)

func SendEmail(to []string, otp string) bool {
	fmt.Println("Preparing to send email...")

	// Load from env (better for security)
	from := os.Getenv("SENDER_EMAIL")
	password := os.Getenv("SENDER_EMAIL_PASSWORD")

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	subject := "I am trying to send email using Golang"
	// body := "This is a testing email sent from Golang backend developed by gkkrrr.\nThis belongs to KD Corporations."

	body := `
		<!DOCTYPE html>
		<html>
		<head>
			<meta charset="UTF-8" />
			<title>Your OTP Code</title>
			<style>
			body {
				font-family: Arial, sans-serif;
				background-color: #f4f4f4;
				padding: 30px;
			}
			.container {
				max-width: 600px;
				margin: auto;
				background: #ffffff;
				border-radius: 8px;
				padding: 25px;
				box-shadow: 0 3px 8px rgba(0,0,0,0.1);
				text-align: center;
			}
			h2 {
				color: #0073e6;
				margin-bottom: 10px;
			}
			p {
				font-size: 15px;
				color: #555;
			}
			.otp-box {
				font-size: 26px;
				font-weight: bold;
				letter-spacing: 4px;
				color: #333;
				background: #f1f1f1;
				padding: 15px 25px;
				display: inline-block;
				border-radius: 6px;
				margin: 20px 0;
			}
			.footer {
				margin-top: 20px;
				font-size: 12px;
				color: #999;
			}
			</style>
		</head>
		<body>
			<div class="container">
			<h2>🔐 Your OTP Code</h2>
			<p>Please use the OTP below to complete your verification process:</p>
			<div class="otp-box">` + otp + `</div>
			<p>This OTP will expire in <b>5 minutes</b>. Do not share it with anyone.</p>
			<div class="footer">
				&copy; 2025 KD Corporations. All rights reserved.
			</div>
			</div>
		</body>
		</html>
		`


	// Build message
	message := []byte("From: " + from + "\r\n" +
		"To: " + strings.Join(to, ",") + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"MIME-Version: 1.0\r\n" +
		"Content-Type: text/html; charset=\"UTF-8\"\r\n" +
		"\r\n" +
		body + "\r\n")




	// Authentication
	auth := smtp.PlainAuth("", from, password, smtpHost)



	// Send email
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		log.Println("Failed to send email:", err)
		return false
	}



	fmt.Println("Email Sent Successfully!")
	return true
}

//kill port 3001
//fuser -k 3001/tcp

