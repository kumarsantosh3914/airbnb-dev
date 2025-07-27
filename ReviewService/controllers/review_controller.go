package controllers

import "ReviewService/services"

type ReviewController struct {
	ReviewService services.ReviewService
}

func NewReviewController(_reviewService services.ReviewService) *ReviewController {
	return &ReviewController{
		ReviewService: _reviewService,
	}
}
