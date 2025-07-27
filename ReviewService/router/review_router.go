package router

import (
	"ReviewService/controllers"

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
	// put all the review routers here
}
