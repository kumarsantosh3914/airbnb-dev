package models

import "time"

type User struct {
	Id              int64
	Username        string
	Email           string
	Password        string
	IsEmailVerified bool
	CreatedAt       string
	UpdatedAt       string
}

type Otp struct {
	Id        int64
	Email     string
	Code      string
	Purpose   string
	ExpiresAt time.Time
	IsUsed    bool
	CreatedAt string
	UpdatedAt string
}
