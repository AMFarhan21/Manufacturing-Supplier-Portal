package rentals_service

type RentalsRepo interface {
	Create(data Rentals) (Rentals, error)
	GetRentalById(id int) (RentalEquipmentUser, error)
	UpdateStatusAndDateRepo(id int, status, startDate, endDate string) error
}
