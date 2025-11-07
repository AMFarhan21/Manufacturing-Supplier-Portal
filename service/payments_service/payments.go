package payments_service

import "time"

type (
	Payments struct {
		Id            int       `json:"id"`
		UserId        string    `json:"user_id"`
		RentalId      int       `json:"rental_id"`
		Amount        float64   `json:"amount"`
		PaymentMethod string    `json:"payment_method"`
		Status        string    `json:"status"`
		CreatedAt     time.Time `json:"created_at"`
	}

	BookingsReport struct {
		Id           int      `json:"id"`
		Name         string   `json:"name"`
		TotalIncome  *float64 `json:"total_income"`
		TotalBooking *int     `json:"total_booking"`
	}
)
