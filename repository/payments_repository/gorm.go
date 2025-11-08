package payments_repository

import (
	"Manufacturing-Supplier-Portal/model"
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

func (r *PaymentsGormRepository) Create(data model.Payments) (model.Payments, error) {
	ctx := context.Background()
	err := r.DB.WithContext(ctx).Create(&data).Error
	if err != nil {
		return model.Payments{}, err
	}

	return data, nil
}

func (r *PaymentsGormRepository) GetAll(userId string) ([]model.Payments, error) {
	ctx := context.Background()
	var payments []model.Payments
	err := r.DB.WithContext(ctx).Where("user_id=?", userId).Find(&payments).Error
	if err != nil {
		return nil, err
	}

	return payments, nil
}

func (r *PaymentsGormRepository) GetById(id int, userId string) (model.Payments, error) {
	ctx := context.Background()
	var payment model.Payments
	err := r.DB.WithContext(ctx).Where("id=?", id).Where("user_id=?", userId).First(&payment).Error
	if err != nil {
		return model.Payments{}, err
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

func (r *PaymentsGormRepository) BookingReport() ([]model.BookingsReport, error) {
	ctx := context.Background()
	var bookingReport []model.BookingsReport
	err := r.DB.WithContext(ctx).
		Select("equipments.id, equipments.name, SUM(payments.amount) as total_income, COUNT(payments.id) as total_booking").
		Joins("LEFT JOIN rentals ON rentals.id = payments.rental_id").
		Joins("LEFT JOIN equipments ON equipments.id = rentals.equipment_id").
		Where("payments.status = ?", "PAID").
		Group("equipments.id, equipments.name").
		Order("equipments.id").
		Scan(&bookingReport).Error
	if err != nil {
		return nil, err
	}

	return bookingReport, nil
}
