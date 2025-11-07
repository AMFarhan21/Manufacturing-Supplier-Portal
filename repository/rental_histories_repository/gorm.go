package rental_histories_repository

import (
	"Manufacturing-Supplier-Portal/service/rental_histories_service"
	"context"

	"gorm.io/gorm"
)

type (
	RentalHistoriesGormRepository struct {
		*gorm.DB
	}
)

func NewRentalHistoriesGormRepository(db *gorm.DB) *RentalHistoriesGormRepository {
	return &RentalHistoriesGormRepository{
		db.Table("rental_histories"),
	}
}

func (r *RentalHistoriesGormRepository) CreateRentalHistory(data rental_histories_service.RentalHistories) (rental_histories_service.RentalHistories, error) {
	ctx := context.Background()
	err := r.DB.WithContext(ctx).Create(&data).Error
	if err != nil {
		return rental_histories_service.RentalHistories{}, err
	}

	return data, nil
}

func (r *RentalHistoriesGormRepository) GetAll(userId string) ([]rental_histories_service.RentalHistories, error) {
	ctx := context.Background()
	var rentals_history []rental_histories_service.RentalHistories
	err := r.DB.WithContext(ctx).Where("user_id=?", userId).Find(&rentals_history).Error
	if err != nil {
		return nil, err
	}
	return rentals_history, nil
}
