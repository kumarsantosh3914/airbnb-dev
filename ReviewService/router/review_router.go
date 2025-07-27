package router

import (
	"ReviewService/controllers"
	middlewares "ReviewService/middlewares"

	"github.com/go-chi/chi/v5"
)

type ReviewRouter struct {
	reviewController *controllers.ReviewController
}

func NewReviewRouter(_reviewController *controllers.ReviewController) Router {
	return &ReviewRouter{
		reviewController: _reviewController,
	}
}

func (rr *ReviewRouter) Register(r chi.Router) {
	r.Route("/api/v1/reviews", func(r chi.Router) {
		r.With(middlewares.ReviewCreateRequestValidator).Post("/", rr.reviewController.CreateReview)
	})
}
