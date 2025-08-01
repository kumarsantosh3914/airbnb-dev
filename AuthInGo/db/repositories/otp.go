package db

import (
	"AuthInGo/models"
	"database/sql"
	"fmt"
	"time"
)

type OTPRepository interface {
	Create(otp *models.Otp) (*models.Otp, error)
	GetValidOTP(email string, code string, purpose string) (*models.Otp, error)
	MarkAsUsed(id int64) error
	InvalidateOTPs(email string, purpose string) error
	DeleteExpiredOTPs() error
}

type OTPRepositoryImpl struct {
	db *sql.DB
}

func NewOTPRepository(_db *sql.DB) OTPRepository {
	return &OTPRepositoryImpl{
		db: _db,
	}
}

func (o *OTPRepositoryImpl) Create(otp *models.Otp) (*models.Otp, error) {
	query := "INSERT INTO otps (email, code, purpose, expires_at, is_used) VALUES (?, ?, ?, ?, ?)"

	result, err := o.db.Exec(query, otp.Email, otp.Code, otp.Purpose, otp.ExpiresAt, otp.IsUsed)
	if err != nil {
		fmt.Println("Error creating OTP:", err)
		return nil, err
	}

	lastInsertID, rowErr := result.LastInsertId()
	if rowErr != nil {
		fmt.Println("Error getting last insert ID:", rowErr)
		return nil, rowErr
	}

	createdOTP := &models.Otp{
		Id:        lastInsertID,
		Email:     otp.Email,
		Code:      otp.Code,
		Purpose:   otp.Purpose,
		ExpiresAt: otp.ExpiresAt,
		IsUsed:    otp.IsUsed,
	}

	fmt.Println("OTP created successfully:", createdOTP.Email, "for purpose:", createdOTP.Purpose)
	return createdOTP, nil
}

func (r *OTPRepositoryImpl) GetValidOTP(email, code, purpose string) (*models.Otp, error) {
	query := `
        SELECT id, email, code, purpose, expires_at, is_used, created_at, updated_at
        FROM otps 
        WHERE email = ? AND code = ? AND purpose = ? AND is_used = false AND expires_at > NOW()
        ORDER BY created_at DESC
        LIMIT 1
    `

	var otp models.Otp
	err := r.db.QueryRow(query, email, code, purpose).Scan(
		&otp.Id,
		&otp.Email,
		&otp.Code,
		&otp.Purpose,
		&otp.ExpiresAt,
		&otp.IsUsed,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No valid OTP found
		}
		return nil, fmt.Errorf("failed to get OTP: %v", err)
	}

	return &otp, nil
}

func (r *OTPRepositoryImpl) MarkAsUsed(id int64) error {
	query := `UPDATE otps SET is_used = true, updated_at = ? WHERE id = ?`

	_, err := r.db.Exec(query, time.Now(), id)
	if err != nil {
		return fmt.Errorf("failed to mark OTP as used: %v", err)
	}

	return nil
}

func (r *OTPRepositoryImpl) InvalidateOTPs(email, purpose string) error {
	query := `UPDATE otps SET is_used = true, updated_at = ? WHERE email = ? AND purpose = ? AND is_used = false`

	_, err := r.db.Exec(query, time.Now(), email, purpose)
	if err != nil {
		return fmt.Errorf("failed to invalidate OTPs: %v", err)
	}

	return nil
}

func (r *OTPRepositoryImpl) DeleteExpiredOTPs() error {
	query := `DELETE FROM otps WHERE expires_at < NOW() OR is_used = true`

	_, err := r.db.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to delete expired OTPs: %v", err)
	}

	return nil
}
