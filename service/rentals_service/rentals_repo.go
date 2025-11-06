package rentals_service

type RentalsRepo interface {
	Create(data Rentals) (Rentals, error)
	GetRentalById(id int) (RentalEquipmentUser, error)
	UpdateStatusAndDate(id int, status, startDate, endDate string) error
}
