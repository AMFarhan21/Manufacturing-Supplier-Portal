package users_service

type UsersRepo interface {
	Register(data Users) (Users, error)
	FindByEmail(email string) (Users, error)
	FindById(id string) (UsersResponse, error)
	GetAll() ([]UsersResponse, error)
	TopUp(userId string, amount float64) (float64, error)
}
