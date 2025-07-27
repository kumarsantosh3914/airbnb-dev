package dto

type ReviewDTO struct {
	Id        int64  `json:"id"`
	UserId    int64  `json:"user_id"`
	HotelId   int64  `json:"hotel_id"`
	BookingId int64  `json:"booking_id"`
	Comment   string `json:"comment"`
	Rating    int64  `json:"rating"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	IsSynced  bool   `json:"is_synced"`
}
