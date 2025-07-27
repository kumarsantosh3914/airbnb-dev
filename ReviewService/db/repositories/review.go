package db

import (
	"ReviewService/models"
	"database/sql"
	"fmt"
)

type ReviewRepository interface {
	CreateReview(userid int64, hotelid int64, bookingid int64, comment string, rating int64) (*models.Review, error)
	GetReviewByID() error
	UpdateReview() error
	DeleteReview() error
}

type ReviewRepositoryImpl struct {
	db *sql.DB
}

func NewReviewRepository(_db *sql.DB) ReviewRepository {
	return &ReviewRepositoryImpl{
		db: _db,
	}
}

func (r *ReviewRepositoryImpl) CreateReview(userid int64, hotelid int64, bookingid int64, comment string, rating int64) (*models.Review, error) {
	query := "INSERT INTO review (user_id, hotel_id, booking_id, comment, rating) VALUES (?, ?, ?, ?, ?)"

	result, err := r.db.Exec(query, userid, hotelid, bookingid, comment, rating)
	if err != nil {
		fmt.Println("Error creating review: ", err)
		return nil, err
	}

	lastInsertID, rowErr := result.LastInsertId()
	if rowErr != nil {
		fmt.Println("Error getting last insert ID: ", rowErr)
		return nil, rowErr
	}

	review := &models.Review{
		Id:        lastInsertID,
		UserId:    userid,
		HotelId:   hotelid,
		BookingId: bookingid,
		Comment:   comment,
		Rating:    rating,
	}

	fmt.Println("Review created successfully: ", review)

	return review, nil
}

func (r *ReviewRepositoryImpl) GetReviewByID() error {
	// Implementation for retrieving a review by ID from the database
	return nil
}

func (r *ReviewRepositoryImpl) UpdateReview() error {
	// Implementation for updating a review in the database
	return nil
}

func (r *ReviewRepositoryImpl) DeleteReview() error {
	// Implementation for deleting a review from the database
	return nil
}
