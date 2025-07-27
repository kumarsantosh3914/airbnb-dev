package models

type Review struct {
	Id        int64
	UserId    int64
	HotelId   int64
	BookingId int64
	Comment   string
	Rating    int64
	CreatedAt string
	UpdatedAt string
	IsSynced  bool
}
