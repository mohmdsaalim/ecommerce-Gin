package utils

import (
	"crypto/rand"
)

// GenerateOTP creates a numeric OTP of given length
func GenerateOTP(length int) (string, error) {
	const digits = "0123456789"
	otp := make([]byte, length)
	_, err := rand.Read(otp)
	if err != nil {
		return "", err
	}

	for i := 0; i < length; i++ {
		otp[i] = digits[otp[i]%10]
	}

	return string(otp), nil
}