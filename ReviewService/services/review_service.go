package services

import (
	db "ReviewService/db/repositories"
	"ReviewService/dto"
	"ReviewService/models"
)

type ReviewService interface {
	CreateReview(payload *dto.ReviewDTO) (*models.Review, error)
	GetReviewByID() error
	UpdateReview() error
	DeleteReview() error
}

type ReviewServiceImpl struct {
	reviewRepository db.ReviewRepository
}

func NewReviewService(_reviewRepository db.ReviewRepository) ReviewService {
	return &ReviewServiceImpl{
		reviewRepository: _reviewRepository,
	}
}

func (r *ReviewServiceImpl) CreateReview(payload *dto.ReviewDTO) (*models.Review, error) {
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

func (r *ReviewServiceImpl) GetReviewByID() error {
	return nil
}

func (r *ReviewServiceImpl) UpdateReview() error {
	return nil
}

func (r *ReviewServiceImpl) DeleteReview() error {
	return nil
}
