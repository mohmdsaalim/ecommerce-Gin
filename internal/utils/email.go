package utils
import (
	"fmt"
	"log"

	"gopkg.in/gomail.v2"
)

type EmailConfig struct {
	Host     string
	Port     int
	Email    string
	Password string
}

func SendOTPEmail(config EmailConfig, to string, otp string) error {

	m := gomail.NewMessage()

	m.SetHeader("From", config.Email)
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Your OTP Verification Code")

	body := fmt.Sprintf(`
		<h2>Email Verification</h2>
		<p>Your OTP code is:</p>
		<h1>%s</h1>
		<p>This OTP is valid for 5 minutes.</p>
	`, otp)

	m.SetBody("text/html", body)

	d := gomail.NewDialer(
		config.Host,
		config.Port,
		config.Email,
		config.Password,
	)

	if err := d.DialAndSend(m); err != nil {
		log.Println("Failed to send email:", err)
		return err
	}

	return nil
}