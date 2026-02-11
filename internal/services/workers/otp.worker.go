package workers

import (
	"log"

	"github.com/mohmdsaalim/ecommerce-Gin/config"
	"github.com/mohmdsaalim/ecommerce-Gin/internal/utils"
)

type OTPJob struct {
	Email string
	Code  string
}

var OTPChannel = make(chan OTPJob, 100)

func StartOTPWorker() {

	log.Println("📨 OTP Worker Started")

	for job := range OTPChannel {

		log.Println("Sending OTP to:", job.Email)

		emailConfig := utils.EmailConfig{
			Host:     config.AppConfig.SMTP.Host,
			Port:     config.AppConfig.SMTP.Port,
			Email:    config.AppConfig.SMTP.Email,
			Password: config.AppConfig.SMTP.Password,
		}

		err := utils.SendOTPEmail(emailConfig, job.Email, job.Code)
		if err != nil {
			log.Println("❌ OTP email failed:", err)
			continue
		}

		log.Println("✅ OTP email sent successfully to:", job.Email)
	}
}