package services

import db "ReviewService/db/repositories"

type ReviewService interface {
	CreateReview() error
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

func (r *ReviewServiceImpl) CreateReview() error {
	return nil
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
