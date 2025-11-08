package users_service

import "Manufacturing-Supplier-Portal/model"

type UsersRepo interface {
	Register(data model.Users) (model.Users, error)
	FindByEmail(email string) (model.Users, error)
	FindById(id string) (model.UsersResponse, error)
	GetAll() ([]model.UsersResponse, error)
	UpdateDepositAmount(userId string, amount float64) (float64, error)
}
