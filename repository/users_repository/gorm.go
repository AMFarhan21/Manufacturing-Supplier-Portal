package users_repository

import (
	"Manufacturing-Supplier-Portal/service/users_service"
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

func (r *UsersGormRepository) Register(data users_service.Users) (users_service.Users, error) {
	ctx := context.Background()
	err := r.DB.WithContext(ctx).Create(&data).Error
	if err != nil {
		return users_service.Users{}, err
	}
	return data, nil
}

func (r *UsersGormRepository) FindByEmail(email string) (users_service.Users, error) {
	ctx := context.Background()
	var user users_service.Users
	err := r.DB.WithContext(ctx).Where("email=?", email).First(&user).Error
	if err != nil {
		return users_service.Users{}, err
	}
	return user, nil
}

func (r *UsersGormRepository) FindById(id string) (users_service.UsersResponse, error) {
	ctx := context.Background()
	var user users_service.UsersResponse
	err := r.DB.WithContext(ctx).Where("id=?", id).First(&user).Error
	if err != nil {
		return users_service.UsersResponse{}, err
	}

	return user, nil
}

func (r *UsersGormRepository) GetAll() ([]users_service.UsersResponse, error) {
	ctx := context.Background()
	var users []users_service.UsersResponse
	err := r.DB.WithContext(ctx).Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UsersGormRepository) TopUp(userId string, amount float64) (float64, error) {
	ctx := context.Background()
	row := r.DB.WithContext(ctx).Where("id=?", userId).Update("deposit_amount", amount)

	if row.RowsAffected == 0 {
		return 0, fmt.Errorf("cannot find user with the id %v", userId)
	}

	if err := row.Error; err != nil {
		return 0, err
	}

	return amount, nil
}
