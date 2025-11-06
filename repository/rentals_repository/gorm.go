package rentals_repository

import (
	"Manufacturing-Supplier-Portal/service/rentals_service"
	"context"

	"gorm.io/gorm"
)

type (
	RentalsGormRepository struct {
		*gorm.DB
	}

	Status struct {
		Status    string `json:"status"`
		StartDate string `json:"start_date"`
		EndDate   string `json:"end_date"`
	}
)

func NewRentalsGormRepository(db *gorm.DB) *RentalsGormRepository {
	return &RentalsGormRepository{
		db.Table("rentals"),
	}
}

func (r RentalsGormRepository) Create(data rentals_service.Rentals) (rentals_service.Rentals, error) {
	ctx := context.Background()
	err := r.DB.WithContext(ctx).Create(&data).Error
	if err != nil {
		return rentals_service.Rentals{}, err
	}

	return data, nil
}

func (r RentalsGormRepository) GetRentalById(id int) (rentals_service.RentalEquipmentUser, error) {
	ctx := context.Background()
	var rentalEquipmentUser rentals_service.RentalEquipmentUser
	err := r.DB.WithContext(ctx).
		Select(`rentals.id as rental_id, user_id, equipment_id, rental_period, start_date, end_date, created_at, equipments.description as description, users.username as username, users.email as email, equipments.name as equipment_name, categories.name as category`).
		Joins(`JOIN equipments ON equipments.id = rentals.equipment_id`).
		Joins(`JOIN users ON users.id = rentals.user_id`).
		Joins(`JOIN categories ON categories.id = equipments.category_id`).
		Where("rentals.id=?", id).
		Scan(&rentalEquipmentUser).Error
	if err != nil {
		return rentals_service.RentalEquipmentUser{}, err
	}

	return rentalEquipmentUser, nil
}

func (r RentalsGormRepository) UpdateStatusAndDate(id int, status, startDate, endDate string) error {
	ctx := context.Background()
	err := r.DB.WithContext(ctx).Where("id=?", id).Updates(Status{
		Status:    status,
		StartDate: startDate,
		EndDate:   endDate,
	}).Error

	if err != nil {
		return err
	}

	return nil
}
