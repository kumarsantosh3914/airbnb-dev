package services

import (
	db "ReviewService/db/repositories"
	"ReviewService/dto"
	"ReviewService/models"
)

type ReviewService interface {
	CreateReview(payload *dto.ReviewDTO) (*models.Review, error)
	GetByID(id int64) (*models.Review, error)
	GetAll() ([]*models.Review, error)
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
