package clients

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// NotificationClient handles communication with the notification service
type NotificationClient interface {
	SendOTPEmail(email, otpCode, purpose string) error
}

type NotificationClientImpl struct {
	baseURL string
}

// NotificationDto represents the structure for notification service API
type NotificationDto struct {
	To         string                 `json:"to"`
	Subject    string                 `json:"subject"`
	TemplateId string                 `json:"templateId"`
	Params     map[string]interface{} `json:"params"`
}

func NewNotificationClient(baseURL string) NotificationClient {
	return &NotificationClientImpl{
		baseURL: baseURL,
	}
}

func (n *NotificationClientImpl) SendOTPEmail(email, otpCode, purpose string) error {
	// Prepare notification payload
	subject := getEmailSubject(purpose)
	templateId := getTemplateId(purpose)
	
	notificationPayload := NotificationDto{
		To:         email,
		Subject:    subject,
		TemplateId: templateId,
		Params: map[string]interface{}{
			"email":   email,
			"otp":     otpCode,
			"purpose": purpose,
		},
	}

	// Convert to JSON
	jsonPayload, err := json.Marshal(notificationPayload)
	if err != nil {
		return fmt.Errorf("failed to marshal notification payload: %v", err)
	}

	// Make HTTP request to notification service
	resp, err := http.Post(
		fmt.Sprintf("%s/api/v1/notifications/send", n.baseURL),
		"application/json",
		bytes.NewBuffer(jsonPayload),
	)
	if err != nil {
		return fmt.Errorf("failed to send HTTP request to notification service: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("notification service returned status: %d", resp.StatusCode)
	}

	return nil
}

// getEmailSubject returns appropriate email subject based on purpose
func getEmailSubject(purpose string) string {
	switch purpose {
	case "email_verification":
		return "Email Verification - Your OTP Code"
	case "password_reset":
		return "Password Reset - Your OTP Code"
	default:
		return "Your OTP Code"
	}
}

// getTemplateId returns appropriate template ID based on purpose
func getTemplateId(purpose string) string {
	switch purpose {
	case "email_verification":
		return "email_verification_otp"
	case "password_reset":
		return "password_reset_otp"
	default:
		return "generic_otp"
	}
}
