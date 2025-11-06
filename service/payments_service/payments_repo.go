package payments_service

type PaymentsRepo interface {
	Create(data Payments) (Payments, error)
	GetById(id int, userId string) (Payments, error)
	UpdateStatusAndMethod(id int, status, method string) error
}
