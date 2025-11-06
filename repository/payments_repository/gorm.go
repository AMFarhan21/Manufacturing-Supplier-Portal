package payments_repository

import (
	"Manufacturing-Supplier-Portal/service/payments_service"
	"context"

	"gorm.io/gorm"
)

type PaymentsGormRepository struct {
	*gorm.DB
}

func NewPaymentsGormRepository(db *gorm.DB) *PaymentsGormRepository {
	return &PaymentsGormRepository{
		db.Table("payments"),
	}
}

func (r *PaymentsGormRepository) Create(data payments_service.Payments) (payments_service.Payments, error) {
	ctx := context.Background()
	err := r.DB.WithContext(ctx).Create(&data).Error
	if err != nil {
		return payments_service.Payments{}, err
	}

	return data, nil
}

func (r *PaymentsGormRepository) GetById(id int, userId string) (payments_service.Payments, error) {
	ctx := context.Background()
	var payment payments_service.Payments
	err := r.DB.WithContext(ctx).Where("id=?", id).Where("user_id=?", userId).First(&payment).Error
	if err != nil {
		return payments_service.Payments{}, err
	}

	return payment, nil
}

func (r *PaymentsGormRepository) UpdateStatus(id int, userId, status string) error {
	ctx := context.Background()
	err := r.DB.WithContext(ctx).Where("id=?", id).Where("user_id=?", userId).Update("status", status).Error
	if err != nil {
		return err
	}

	return nil
}
