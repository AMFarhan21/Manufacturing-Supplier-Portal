package rental_histories_service

import "Manufacturing-Supplier-Portal/model"

type RentalHistoriesRepo interface {
	CreateRentalHistory(data model.RentalHistories) (model.RentalHistories, error)
	GetAll(userId string) ([]model.RentalHistories, error)
}
