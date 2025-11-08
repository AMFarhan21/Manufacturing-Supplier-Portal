package rentals_service

import (
	"Manufacturing-Supplier-Portal/model"
	"time"
)

type RentalsRepo interface {
	Create(data model.Rentals) (model.Rentals, error)
	GetRentalById(id int) (model.RentalEquipmentUser, error)
	UpdateStatusAndDateRepo(id int, status string, startDate, endDate time.Time) error
	SimulateAutomaticUpdateRentalStatus() error
}
