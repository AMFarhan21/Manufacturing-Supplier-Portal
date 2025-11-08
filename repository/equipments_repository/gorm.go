package equipments_repository

import (
	"Manufacturing-Supplier-Portal/model"
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

func (r *EquipmentsGormRepository) Create(data model.Equipments) (model.Equipments, error) {
	ctx := context.Background()
	err := r.DB.WithContext(ctx).Create(&data).Error
	if err != nil {
		return model.Equipments{}, err
	}

	return data, nil
}

func (r *EquipmentsGormRepository) GetAll() ([]model.Equipments, error) {
	ctx := context.Background()
	var equipments []model.Equipments
	err := r.DB.WithContext(ctx).Find(&equipments).Error
	if err != nil {
		return nil, err
	}

	return equipments, nil
}

func (r *EquipmentsGormRepository) GetById(id int) (model.Equipments, error) {
	ctx := context.Background()
	var equipment model.Equipments
	err := r.DB.WithContext(ctx).Where("id=?", id).First(&equipment).Error
	if err != nil {
		return model.Equipments{}, err
	}

	return equipment, nil
}

func (r EquipmentsGormRepository) Update(id int, data model.Equipments) (model.Equipments, error) {
	ctx := context.Background()
	row := r.DB.WithContext(ctx).Where("id=?", id).Updates(data)
	if row.RowsAffected == 0 {
		return model.Equipments{}, errors.New("equipment id not found")
	}

	err := row.Error
	if err != nil {
		return model.Equipments{}, err
	}

	data.Id = id

	return data, nil
}

func (r *EquipmentsGormRepository) Delete(id int) error {
	ctx := context.Background()
	row := r.DB.WithContext(ctx).Where("id=?", id).Delete(model.Equipments{})
	if row.RowsAffected == 0 {
		return errors.New("equipment id not found")
	}

	if err := row.Error; err != nil {
		return err
	}

	return nil
}

func (r *EquipmentsGormRepository) UpdateStatus(id int, available bool) error {
	ctx := context.Background()
	row := r.DB.WithContext(ctx).Where("id=?", id).Update("available", available)
	if row.RowsAffected == 0 {
		return errors.New("equipment id not found")
	}

	if err := row.Error; err != nil {
		return err
	}

	return nil
}
