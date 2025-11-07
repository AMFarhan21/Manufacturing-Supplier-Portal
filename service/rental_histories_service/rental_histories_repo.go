package rental_histories_service

type RentalHistoriesRepo interface {
	CreateRentalHistory(data RentalHistories) (RentalHistories, error)
	GetAll(userId string) ([]RentalHistories, error)
}
