package payments_service

import "time"

type Payments struct {
	Id            int       `json:"id"`
	UserId        string    `json:"user_id"`
	RentalId      int       `json:"rental_id"`
	Amount        float64   `json:"amount"`
	PaymentMethod string    `json:"payment_method"`
	Status        string    `json:"status"`
	CreatedAt     time.Time `json:"created_at"`
}
