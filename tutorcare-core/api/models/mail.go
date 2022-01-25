package models

import (
	"fmt"
	"math/rand"
	"net/smtp"
	"strconv"
	"strings"
	"time"
)

type Mail struct {
	Sender  string
	Pass    string
	To      []string
	Subject string
	Body    string
}

func sendEmail(email Mail) {
	//Default Email Sender
	if email.Sender == "" {
		email.Sender = "tutorcaregatech@gmail.com"
		email.Pass = "tutorcareemailpassword"
	}

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	message := BuildMessage(email)

	// Authentication.
	auth := smtp.PlainAuth("", email.Sender, email.Pass, smtpHost)

	// Send email
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, email.Sender, email.To, []byte(message))

	if err != nil {
		panic(err)
	}
}

func BuildMessage(mail Mail) string {
	msg := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\r\n"
	msg += fmt.Sprintf("From: %s\r\n", mail.Sender)
	msg += fmt.Sprintf("To: %s\r\n", strings.Join(mail.To, ";"))
	msg += fmt.Sprintf("Subject: %s\r\n", mail.Subject)
	msg += fmt.Sprintf("\r\n%s\r\n", mail.Body)

	return msg
}

func emailVerificationCode(to []string) int {
	code := generateVerificationCode()

	subject := "TutorCare Email Verification Code"
	body := "<p>Your validation code is <b>" + strconv.Itoa(code) + "</b></p>"

	request := Mail{
		To:      to,
		Subject: subject,
		Body:    body,
	}

	sendEmail(request)
	return code
}

func generateVerificationCode() int {
	rand.Seed(time.Now().UnixNano())
	return 1000 + rand.Intn(999999-1000)
}
