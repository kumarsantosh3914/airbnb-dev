package dto

type SendOTPRequestDTO struct {
	Email   string `json:"email" validate:"required,email"`
	Purpose string `json:"purpose" validate:"required,oneof=email_verification password_reset"`
}

type VerifyOTPRequestDTO struct {
	Email   string `json:"email" validate:"required,email"`
	Code    string `json:"code" validate:"required,len=6,numeric"`
	Purpose string `json:"purpose" validate:"required,oneof=email_verification password_reset"`
}
