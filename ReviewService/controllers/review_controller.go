package controllers

import (
	"ReviewService/dto"
	"ReviewService/services"
	"ReviewService/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type ReviewController struct {
	ReviewService services.ReviewService
}

func NewReviewController(_reviewService services.ReviewService) *ReviewController {
	return &ReviewController{
		ReviewService: _reviewService,
	}
}

func (rc *ReviewController) CreateReview(w http.ResponseWriter, r *http.Request) {
	payload := r.Context().Value("payload").(dto.ReviewDTO)

	fmt.Println("Payload received: ", payload)

	review, err := rc.ReviewService.CreateReview(&payload)
	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError, "Falied to create review", err)
		return
	}

	utils.WriteJsonSuccessResponse(w, http.StatusCreated, "Review created successfully", review)

	fmt.Println("Review created successfully: ", review)
}

func (rc *ReviewController) GetAll(w http.ResponseWriter, r *http.Request) {
	reviews, err := rc.ReviewService.GetAll()
	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError, "Failed to fetch reviews", err)
		return
	}

	utils.WriteJsonSuccessResponse(w, http.StatusOK, "Reviews fetched successfully", reviews)
}

func (rc *ReviewController) GetByID(w http.ResponseWriter, r *http.Request) {
	reviewIdStr := chi.URLParam(r, "id")
	if reviewIdStr == "" {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Review ID is required", fmt.Errorf("Missing review ID"))
		return
	}

	reviewId, err := strconv.ParseInt(reviewIdStr, 10, 64)
	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Invalid review ID", err)
		return
	}

	review, err := rc.ReviewService.GetByID(reviewId)
	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError, "Failed to fetch review", err)
		return
	}
	if review == nil {
		utils.WriteJsonErrorResponse(w, http.StatusNotFound, "Review not found", nil)
		return
	}

	utils.WriteJsonSuccessResponse(w, http.StatusOK, "Review fetched successfully", review)
}
