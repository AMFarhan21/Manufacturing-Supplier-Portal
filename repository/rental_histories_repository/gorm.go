package rental_histories_repository

import (
	"Manufacturing-Supplier-Portal/model"
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

func (r *RentalHistoriesGormRepository) CreateRentalHistory(data model.RentalHistories) (model.RentalHistories, error) {
	ctx := context.Background()
	err := r.DB.WithContext(ctx).Create(&data).Error
	if err != nil {
		return model.RentalHistories{}, err
	}

	return data, nil
}

func (r *RentalHistoriesGormRepository) GetAll(userId string) ([]model.RentalHistories, error) {
	ctx := context.Background()
	var rentals_history []model.RentalHistories
	err := r.DB.WithContext(ctx).Where("user_id=?", userId).Find(&rentals_history).Error
	if err != nil {
		return nil, err
	}
	return rentals_history, nil
}
