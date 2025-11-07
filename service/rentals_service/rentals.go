package rentals_service

import "time"

type (
	Rentals struct {
		Id           int        `json:"id"`
		UserId       string     `json:"user_id"`
		EquipmentId  int        `json:"equipment_id"`
		RentalPeriod string     `json:"rental_period"`
		StartDate    *time.Time `json:"start_date"`
		EndDate      *time.Time `json:"end_date"`
		Price        float64    `json:"price"`
		Status       string     `json:"status"`
		CreatedAt    time.Time  `json:"created_at"`
	}

	RentalsWithInvoiceUrl struct {
		Id           int        `json:"id"`
		UserId       string     `json:"user_id"`
		EquipmentId  int        `json:"equipment_id"`
		RentalPeriod string     `json:"rental_period"`
		StartDate    *time.Time `json:"start_date"`
		EndDate      *time.Time `json:"end_date"`
		Price        float64    `json:"price"`
		Status       string     `json:"status"`
		CreatedAt    time.Time  `json:"created_at"`
		InvoiceUrl   string     `json:"invoice_url"`
	}

	RentalEquipmentUser struct {
		RentalId      int        `json:"rental_id"`
		UserId        string     `json:"user_id"`
		EquipmentId   int        `json:"equipment_id"`
		RentalPeriod  string     `json:"rental_period"`
		StartDate     *time.Time `json:"start_date"`
		EndDate       *time.Time `json:"end_date"`
		Price         float64    `json:"price"`
		Status        string     `json:"status"`
		CreatedAt     time.Time  `json:"created_at"`
		Description   string     `json:"description"`
		Username      string     `json:"username"`
		Email         string     `json:"email"`
		EquipmentName string     `json:"equipment_name"`
		Category      string     `json:"category"`
	}
)

// userId string, description string, username string, email string, name string, category string, rentalId int, amount float64
