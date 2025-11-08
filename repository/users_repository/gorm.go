package users_repository

import (
	"Manufacturing-Supplier-Portal/model"
	"context"
	"fmt"

	"gorm.io/gorm"
)

type UsersGormRepository struct {
	*gorm.DB
}

func NewUsersGormRepository(db *gorm.DB) *UsersGormRepository {
	return &UsersGormRepository{
		db.Table("users"),
	}
}

func (r *UsersGormRepository) Register(data model.Users) (model.Users, error) {
	ctx := context.Background()
	err := r.DB.WithContext(ctx).Create(&data).Error
	if err != nil {
		return model.Users{}, err
	}
	return data, nil
}

func (r *UsersGormRepository) FindByEmail(email string) (model.Users, error) {
	ctx := context.Background()
	var user model.Users
	err := r.DB.WithContext(ctx).Where("email=?", email).First(&user).Error
	if err != nil {
		return model.Users{}, err
	}
	return user, nil
}

func (r *UsersGormRepository) FindById(id string) (model.UsersResponse, error) {
	ctx := context.Background()
	var user model.UsersResponse
	err := r.DB.WithContext(ctx).Where("id=?", id).First(&user).Error
	if err != nil {
		return model.UsersResponse{}, err
	}

	return user, nil
}

func (r *UsersGormRepository) GetAll() ([]model.UsersResponse, error) {
	ctx := context.Background()
	var users []model.UsersResponse
	err := r.DB.WithContext(ctx).Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UsersGormRepository) UpdateDepositAmount(userId string, amount float64) (float64, error) {
	ctx := context.Background()
	row := r.DB.WithContext(ctx).Where("id=?", userId).Update("deposit_amount", amount)

	if err := row.Error; err != nil {
		return 0, err
	}

	if row.RowsAffected == 0 {
		return 0, fmt.Errorf("cannot find user with the id %v", userId)
	}

	return amount, nil
}
