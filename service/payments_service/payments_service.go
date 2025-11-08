package payments_service

import "Manufacturing-Supplier-Portal/model"

type PaymentsService struct {
	repo PaymentsRepo
}

type Service interface {
	Create(data model.Payments) (model.Payments, error)
	GetById(id int, userId string) (model.Payments, error)
	UpdateStatusAndMethod(id int, status, method string) error
	BookingReport() ([]model.BookingsReport, error)
	GetAllPayments(userId string) ([]model.Payments, error)
}

func NewPaymentsService(repo PaymentsRepo) Service {
	return &PaymentsService{
		repo: repo,
	}
}

func (s PaymentsService) Create(data model.Payments) (model.Payments, error) {
	return s.repo.Create(data)
}
func (s PaymentsService) GetById(id int, userId string) (model.Payments, error) {
	return s.repo.GetById(id, userId)
}
func (s PaymentsService) UpdateStatusAndMethod(id int, status, method string) error {
	return s.repo.UpdateStatusAndMethod(id, status, method)
}
func (s PaymentsService) BookingReport() ([]model.BookingsReport, error) {
	return s.repo.BookingReport()
}

func (s PaymentsService) GetAllPayments(userId string) ([]model.Payments, error) {
	return s.repo.GetAll(userId)
}
