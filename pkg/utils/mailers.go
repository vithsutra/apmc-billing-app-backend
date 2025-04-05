package utils

import (
	"bytes"
	"html/template"
	"log"

	"gopkg.in/gomail.v2"
)

func SendResetTokenMail(toEmail, otp string) error {
	type TemplateData struct {
		OTP           string
		ExpireMinutes int
	}

	tmpl, err := template.New("otpEmail").Parse(otpEmailTemplate)
	if err != nil {
		log.Println("Error parsing template:", err)
		return err
	}

	var body bytes.Buffer
	err = tmpl.Execute(&body, TemplateData{
		OTP:           otp,
		ExpireMinutes: 1,
	})
	if err != nil {
		log.Println("Error executing template:", err)
		return err
	}

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", "noreply.vithsutra@gmail.com")
	mailer.SetHeader("To", toEmail)
	mailer.SetHeader("Subject", "Your OTP for Password Reset")
	mailer.SetBody("text/html", body.String())

	dialer := gomail.NewDialer("smtp.gmail.com", 587, "noreply.vithsutra@gmail.com", "vlcoctlouuzmwqqv")

	if err := dialer.DialAndSend(mailer); err != nil {
		log.Println("Failed to send OTP email:", err)
		return err
	}

	log.Println("OTP email sent to", toEmail)
	return nil
}
