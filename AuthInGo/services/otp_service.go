package services

import (
	"AuthInGo/clients"
	"AuthInGo/dto"
	"AuthInGo/utils"
	db "AuthInGo/db/repositories"
	"fmt"
)

type OTPService interface {
	SendOTP(payload *dto.SendOTPRequestDTO) error
	VerifyOTP(payload *dto.VerifyOTPRequestDTO) error
	ResendOTP(payload *dto.SendOTPRequestDTO) error
}

type OTPServiceImpl struct {
	otpRepository      db.OTPRepository
	notificationClient clients.NotificationClient
}

func NewOTPService(_otpRepository db.OTPRepository, _notificationClient clients.NotificationClient) OTPService {
	return &OTPServiceImpl{
		otpRepository:      _otpRepository,
		notificationClient: _notificationClient,
	}
}

func (o *OTPServiceImpl) SendOTP(payload *dto.SendOTPRequestDTO) error {
	fmt.Println("Sending OTP in OTPService")

	// Invalidate any existing OTPs for this email and purpose
	err := o.otpRepository.InvalidateOTPs(payload.Email, payload.Purpose)
	if err != nil {
		fmt.Println("Error invalidating existing OTPs:", err)
		return fmt.Errorf("failed to invalidate existing OTPs: %v", err)
	}

	// Generate new OTP
	otpCode, err := utils.GenerateOTP()
	if err != nil {
		fmt.Println("Error generating OTP:", err)
		return fmt.Errorf("failed to generate OTP: %v", err)
	}

	// Create OTP record using utility function
	otp := utils.CreateOTPModel(payload.Email, otpCode, payload.Purpose)

	createdOTP, err := o.otpRepository.Create(otp)
	if err != nil {
		fmt.Println("Error creating OTP:", err)
		return fmt.Errorf("failed to create OTP: %v", err)
	}

	// Send OTP via notification client
	err = o.notificationClient.SendOTPEmail(payload.Email, otpCode, payload.Purpose)
	if err != nil {
		fmt.Println("Error sending OTP email:", err)
		return fmt.Errorf("failed to send OTP email: %v", err)
	}

	fmt.Printf("OTP sent successfully to %s for purpose %s with ID %d\n", 
		createdOTP.Email, createdOTP.Purpose, createdOTP.Id)
	return nil
}

func (o *OTPServiceImpl) VerifyOTP(payload *dto.VerifyOTPRequestDTO) error {
	fmt.Println("Verifying OTP in OTPService")

	// Get valid OTP
	otp, err := o.otpRepository.GetValidOTP(payload.Email, payload.Code, payload.Purpose)
	if err != nil {
		fmt.Println("Error getting valid OTP:", err)
		return fmt.Errorf("failed to get valid OTP: %v", err)
	}

	if otp == nil {
		fmt.Println("Invalid or expired OTP")
		return fmt.Errorf("invalid or expired OTP")
	}

	// Mark OTP as used
	err = o.otpRepository.MarkAsUsed(otp.Id)
	if err != nil {
		fmt.Println("Error marking OTP as used:", err)
		return fmt.Errorf("failed to mark OTP as used: %v", err)
	}

	fmt.Printf("OTP verified successfully for %s with purpose %s\n", payload.Email, payload.Purpose)
	return nil
}

func (o *OTPServiceImpl) ResendOTP(payload *dto.SendOTPRequestDTO) error {
	fmt.Println("Resending OTP in OTPService")
	
	// Resend is essentially the same as sending a new OTP
	return o.SendOTP(payload)
}


