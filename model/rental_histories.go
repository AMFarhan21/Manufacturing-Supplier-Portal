package model

import "time"

type RentalHistories struct {
	Id        int       `json:"id"`
	RentalId  int       `json:"rental_id"`
	UserId    string    `json:"user_id"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}
