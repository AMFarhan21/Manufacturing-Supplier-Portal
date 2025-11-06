package payments_service

type PaymentsRepo interface {
	Create(data Payments) (Payments, error)
	GetById(id int, userId string) (Payments, error)
	UpdateStatus(id int, userId, status string) error
}
