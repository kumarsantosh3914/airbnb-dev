package utils

import (
	"AuthInGo/models"
	"crypto/rand"
	"time"
)

// GenerateOTP generates a cryptographically secure 6-digit OTP
func GenerateOTP() (string, error) {
	const digits = "0123456789"
	const otpLength = 6
	
	bytes := make([]byte, otpLength)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	
	for i := 0; i < otpLength; i++ {
		bytes[i] = digits[bytes[i]%byte(len(digits))]
	}
	
	return string(bytes), nil
}

// CreateOTPModel creates a new OTP model with the given parameters
func CreateOTPModel(email, code, purpose string) *models.Otp {
	return &models.Otp{
		Email:     email,
		Code:      code,
		Purpose:   purpose,
		ExpiresAt: time.Now().Add(10 * time.Minute), // 10 minutes expiration
		IsUsed:    false,
	}
}

// GetOTPExpirationTime returns the standard OTP expiration duration
func GetOTPExpirationTime() time.Duration {
	return 10 * time.Minute
}