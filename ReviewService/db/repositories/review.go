package db

import (
	"ReviewService/models"
	"database/sql"
	"fmt"
)

type ReviewRepository interface {
	CreateReview(userid int64, hotelid int64, bookingid int64, comment string, rating int64) (*models.Review, error)
	GetByID(id int64) (*models.Review, error)
	GetAll() ([]*models.Review, error)
	SoftDelete() error
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

func (r *ReviewRepositoryImpl) GetByID(id int64) (*models.Review, error) {
	query := "SELECT id, user_id, hotel_id, booking_id, comment, rating, created_at, updated_at, is_synced FROM review WHERE id = ?"

	row := r.db.QueryRow(query, id)

	review := &models.Review{}

	err := row.Scan(
		&review.Id,
		&review.UserId,
		&review.HotelId,
		&review.BookingId,
		&review.Comment,
		&review.Rating,
		&review.CreatedAt,
		&review.UpdatedAt,
		&review.IsSynced,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No review found with the given ID")
			return nil, nil
		} else {
			fmt.Println("Error scanning review:", err)
			return nil, err
		}
	}

	fmt.Println("Review fetched successfully:", review)

	return review, nil
}

func (r *ReviewRepositoryImpl) GetAll() ([]*models.Review, error) {
	query := "SELECT id, user_id, hotel_id, booking_id, comment, rating, created_at, updated_at, is_synced FROM review"
	rows, err := r.db.Query(query)
	if err != nil {
		fmt.Println("Error querying reviews: ", err)
		return nil, err
	}
	defer rows.Close()

	var reviews []*models.Review
	for rows.Next() {
		review := new(models.Review)
		err := rows.Scan(
			&review.Id,
			&review.UserId,
			&review.HotelId,
			&review.BookingId,
			&review.Comment,
			&review.Rating,
			&review.CreatedAt,
			&review.UpdatedAt,
			&review.IsSynced,
		)
		if err != nil {
			fmt.Println("Error scanning review row: ", err)
			return nil, err
		}
		reviews = append(reviews, review)
	}

	if err = rows.Err(); err != nil {
		fmt.Println("Row iteration error: ", err)
		return nil, err
	}

	return reviews, nil
}

func (r *ReviewRepositoryImpl) SoftDelete() error {
	return nil
}
