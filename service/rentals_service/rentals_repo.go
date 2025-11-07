package rentals_service

import "time"

type RentalsRepo interface {
	Create(data Rentals) (Rentals, error)
	GetRentalById(id int) (RentalEquipmentUser, error)
	UpdateStatusAndDateRepo(id int, status string, startDate, endDate time.Time) error
	SimulateAutomaticUpdateRentalStatus() error
}
