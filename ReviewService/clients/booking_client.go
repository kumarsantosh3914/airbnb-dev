package clients

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type BookingClient struct {
	baseURL    string
	httpClient *http.Client
}

type BookingData struct {
	ID            int    `json:"id"`
	UserID        int    `json:"userId"`
	HotelID       int    `json:"hotelId"`
	BookingAmount int    `json:"bookingAmount"`
	Status        string `json:"status"`
	TotalGuests   int    `json:"totalGuests"`
	CratedAt      string `json:"createdAt"`
	UpdatedAt     string `json:"updatedAt"`
}

type BookingResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    BookingData `json:"data"`
}

func NewBookingClient(baseURL string, httpClient *http.Client) *BookingClient {
	return &BookingClient{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: 30 * time.Second, // Set a timeout for the HTTP client
		},
	}
}

func (bc *BookingClient) GetBookingByID(bookingID int64) (*BookingData, error) {
	url := fmt.Sprintf("%s/api/v1/bookings/%d", bc.baseURL, bookingID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := bc.httpClient.Do(req)

	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get booking: status code %d", resp.StatusCode)
	}

	var bookingResponse BookingResponse
	if err := json.NewDecoder(resp.Body).Decode(&bookingResponse); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	if bookingResponse.Status != "success" {
		return nil, fmt.Errorf("error from booking service: %s", bookingResponse.Message)
	}

	return &bookingResponse.Data, nil
}

func (bc *BookingClient) ValidateBooking(bookingID int64, userID int64) (bool, error) {
	booking, err := bc.GetBookingByID(bookingID)
	if err != nil {
		return false, err
	}

	// validate that the booking belongs to the user
	if booking.UserID != int(userID) {
		return false, fmt.Errorf("booking does not belong to user: expected %d, got %d", userID, booking.UserID)
	}

	if booking.Status != "CONFIRMED" {
		return false, fmt.Errorf("booking is not confirmed: %s", booking.Status)
	}

	return true, nil
}
