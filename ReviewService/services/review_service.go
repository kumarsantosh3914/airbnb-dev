package services

import (
	"ReviewService/clients"
	env "ReviewService/config/env"
	db "ReviewService/db/repositories"
	"ReviewService/dto"
	"ReviewService/models"
	"fmt"
)

type ReviewService interface {
	CreateReview(payload *dto.ReviewDTO) (*models.Review, error)
	GetByID(id int64) (*models.Review, error)
	GetAll() ([]*models.Review, error)
	DeleteReview() error
}

type ReviewServiceImpl struct {
	reviewRepository db.ReviewRepository
	bookingClient    *clients.BookingClient
}

func NewReviewService(_reviewRepository db.ReviewRepository) ReviewService {
	bookingClient := clients.NewBookingClient(env.GetString("BOOKING_SERVICE_URL", "http://localhost:3000"), nil)
	return &ReviewServiceImpl{
		reviewRepository: _reviewRepository,
		bookingClient:    bookingClient,
	}
}

func (r *ReviewServiceImpl) CreateReview(payload *dto.ReviewDTO) (*models.Review, error) {
	// validate booking with BookingService
	isValid, err := r.bookingClient.ValidateBooking(payload.BookingId, payload.UserId)
	if err != nil {
		return nil, err
	}

	if !isValid {
		return nil, fmt.Errorf("invalid booking for user %d", payload.UserId)
	}

	review, err := r.reviewRepository.CreateReview(
		payload.UserId,
		payload.HotelId,
		payload.BookingId,
		payload.Comment,
		payload.Rating,
	)
	if err != nil {
		return nil, err
	}
	return review, nil
}

func (r *ReviewServiceImpl) GetByID(id int64) (*models.Review, error) {
	review, err := r.reviewRepository.GetByID(id)
	if err != nil {
		return nil, err
	}
	return review, nil
}

func (r *ReviewServiceImpl) GetAll() ([]*models.Review, error) {
	reviews, err := r.reviewRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return reviews, nil
}

func (r *ReviewServiceImpl) DeleteReview() error {
	return nil
}
