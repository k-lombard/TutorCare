package models

import (
	"fmt"
	"math/rand"
	"net/smtp"
	"os"
	"strconv"
	"strings"
	"time"
)

type mail struct {
	Sender  string
	Pass    string
	To      []string
	Subject string
	Body    string
}

func SendEmailVerificationCode(to []string) int {
	code := generateVerificationCode()

	subject := "TutorCare Email Verification Code"
	body := "<p>Your validation code is <b>" + strconv.Itoa(code) + "</b></p><br><a href='http://localhost:4200/verify'>Verify Email Here</a>"
	pass := os.Getenv("EMAIL_PASS")

	request := mail{
		Sender:  "tutorcaregatech@gmail.com",
		Pass:    pass,
		To:      to,
		Subject: subject,
		Body:    body,
	}

	sendEmail(request)
	return code
}

func sendEmail(email mail) {

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	message := buildMessage(email)

	// Authentication.
	auth := smtp.PlainAuth("", email.Sender, email.Pass, smtpHost)

	// Send email
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, email.Sender, email.To, []byte(message))

	if err != nil {
		fmt.Println("Email not sent")
		panic(err)
	}
	fmt.Println("Email sent successfully")
}

func buildMessage(mail mail) string {
	msg := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\r\n"
	msg += fmt.Sprintf("From: %s\r\n", mail.Sender)
	msg += fmt.Sprintf("To: %s\r\n", strings.Join(mail.To, ";"))
	msg += fmt.Sprintf("Subject: %s\r\n", mail.Subject)
	msg += fmt.Sprintf("\r\n%s\r\n", mail.Body)

	return msg
}

func generateVerificationCode() int {
	rand.Seed(time.Now().UnixNano())
	return 1000 + rand.Intn(999999-1000)
}
