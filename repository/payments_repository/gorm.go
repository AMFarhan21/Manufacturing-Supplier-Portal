package payments_repository

import (
	"Manufacturing-Supplier-Portal/service/payments_service"
	"context"

	"gorm.io/gorm"
)

type (
	PaymentsGormRepository struct {
		*gorm.DB
	}

	Updates struct {
		Status        string `json:"status"`
		PaymentMethod string `json:"payment_method"`
	}
)

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

func (r *PaymentsGormRepository) UpdateStatusAndMethod(id int, status, method string) error {
	ctx := context.Background()
	err := r.DB.WithContext(ctx).Where("id=?", id).Updates(Updates{
		Status:        status,
		PaymentMethod: method,
	}).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *PaymentsGormRepository) BookingReport() (payments_service.BookingsReport, error) {
	ctx := context.Background()
	var bookingReport payments_service.BookingsReport
	err := r.DB.WithContext(ctx).
		Select("equipments.id, equipments.name, SUM(payments.amount) as total_income, COUNT(payments.id) as total_booking").
		Joins("LEFT JOIN rentals ON rentals.id = payments.rental_id").
		Joins("LEFT JOIN equipments ON equipments.id = rentals.equipment_id").
		Where("payments.status = ?", "PAID").
		Group("equipments.id, equipments.name").
		Order("equipments.id").
		Scan(&bookingReport).Error
	if err != nil {
		return payments_service.BookingsReport{}, err
	}

	return bookingReport, nil
}
