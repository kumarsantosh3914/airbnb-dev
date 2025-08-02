package router

import (
	"AuthInGo/controllers"

	"github.com/go-chi/chi/v5"
)

type OTPRouter struct {
	otpController *controllers.OTPController
}

func NewOTPRouter(otpController *controllers.OTPController) *OTPRouter {
	return &OTPRouter{
		otpController: otpController,
	}
}

func (or *OTPRouter) Register(r chi.Router) {
	r.Route("/api/otp", func(r chi.Router) {
		r.Post("/send", or.otpController.SendOTP)
		r.Post("/verify", or.otpController.VerifyOTP)
		r.Post("/resend", or.otpController.ResendOTP)
	})
}
