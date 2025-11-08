package payments_service

import "Manufacturing-Supplier-Portal/model"

type PaymentsRepo interface {
	Create(data model.Payments) (model.Payments, error)
	GetById(id int, userId string) (model.Payments, error)
	UpdateStatusAndMethod(id int, status, method string) error
	BookingReport() ([]model.BookingsReport, error)
	GetAll(userId string) ([]model.Payments, error)
}
