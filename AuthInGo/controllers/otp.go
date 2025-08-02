package controllers

import (
	"AuthInGo/dto"
	"AuthInGo/services"
	"encoding/json"
	"fmt"
	"net/http"
)

type OTPController struct {
	otpService services.OTPService
}

func NewOTPController(otpService services.OTPService) *OTPController {
	return &OTPController{
		otpService: otpService,
	}
}

// SendOTP handles OTP sending requests
func (c *OTPController) SendOTP(w http.ResponseWriter, r *http.Request) {
	var payload dto.SendOTPRequestDTO

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate required fields
	if payload.Email == "" {
		http.Error(w, "Email is required", http.StatusBadRequest)
		return
	}

	if payload.Purpose == "" {
		payload.Purpose = "email_verification" // Default purpose
	}

	err := c.otpService.SendOTP(&payload)
	if err != nil {
		fmt.Printf("Error sending OTP: %v\n", err)
		http.Error(w, fmt.Sprintf("Failed to send OTP: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"message": "OTP sent successfully",
		"email":   payload.Email,
		"purpose": payload.Purpose,
	})
}

// VerifyOTP handles OTP verification requests
func (c *OTPController) VerifyOTP(w http.ResponseWriter, r *http.Request) {
	var payload dto.VerifyOTPRequestDTO

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate required fields
	if payload.Email == "" || payload.Code == "" {
		http.Error(w, "Email and code are required", http.StatusBadRequest)
		return
	}

	if payload.Purpose == "" {
		payload.Purpose = "email_verification" // Default purpose
	}

	err := c.otpService.VerifyOTP(&payload)
	if err != nil {
		fmt.Printf("Error verifying OTP: %v\n", err)
		http.Error(w, fmt.Sprintf("Failed to verify OTP: %v", err), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"message": "OTP verified successfully",
		"email":   payload.Email,
		"purpose": payload.Purpose,
	})
}

// ResendOTP handles OTP resend requests
func (c *OTPController) ResendOTP(w http.ResponseWriter, r *http.Request) {
	var payload dto.SendOTPRequestDTO

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate required fields
	if payload.Email == "" {
		http.Error(w, "Email is required", http.StatusBadRequest)
		return
	}

	if payload.Purpose == "" {
		payload.Purpose = "email_verification" // Default purpose
	}

	err := c.otpService.ResendOTP(&payload)
	if err != nil {
		fmt.Printf("Error resending OTP: %v\n", err)
		http.Error(w, fmt.Sprintf("Failed to resend OTP: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"message": "OTP resent successfully",
		"email":   payload.Email,
		"purpose": payload.Purpose,
	})
}
