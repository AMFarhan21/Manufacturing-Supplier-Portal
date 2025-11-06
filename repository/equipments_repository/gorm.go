package equipments_repository

import (
	"Manufacturing-Supplier-Portal/service/equipments_service"
	"context"
	"errors"

	"gorm.io/gorm"
)

type EquipmentsGormRepository struct {
	*gorm.DB
}

func NewEquipmentsGormRepository(db *gorm.DB) *EquipmentsGormRepository {
	return &EquipmentsGormRepository{
		db.Table("equipments"),
	}
}

func (r *EquipmentsGormRepository) Create(data equipments_service.Equipments) (equipments_service.Equipments, error) {
	ctx := context.Background()
	err := r.DB.WithContext(ctx).Create(&data).Error
	if err != nil {
		return equipments_service.Equipments{}, err
	}

	return data, nil
}

func (r *EquipmentsGormRepository) GetAll() ([]equipments_service.Equipments, error) {
	ctx := context.Background()
	var equipments []equipments_service.Equipments
	err := r.DB.WithContext(ctx).Find(&equipments).Error
	if err != nil {
		return nil, err
	}

	return equipments, nil
}

func (r *EquipmentsGormRepository) GetById(id int) (equipments_service.Equipments, error) {
	ctx := context.Background()
	var equipment equipments_service.Equipments
	err := r.DB.WithContext(ctx).Where("id=?", id).First(&equipment).Error
	if err != nil {
		return equipments_service.Equipments{}, err
	}

	return equipment, nil
}

func (r EquipmentsGormRepository) Update(id int, data equipments_service.Equipments) (equipments_service.Equipments, error) {
	ctx := context.Background()
	row := r.DB.WithContext(ctx).Where("id=?", id).Updates(data)
	if row.RowsAffected == 0 {
		return equipments_service.Equipments{}, errors.New("equipment id not found")
	}

	err := row.Error
	if err != nil {
		return equipments_service.Equipments{}, err
	}

	data.Id = id

	return data, nil
}

func (r *EquipmentsGormRepository) Delete(id int) error {
	ctx := context.Background()
	row := r.DB.WithContext(ctx).Where("id=?", id).Delete(equipments_service.Equipments{})
	if row.RowsAffected == 0 {
		return errors.New("equipment id not found")
	}

	if err := row.Error; err != nil {
		return err
	}

	return nil
}
