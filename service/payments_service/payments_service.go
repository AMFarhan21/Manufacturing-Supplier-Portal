package payments_service

type PaymentsService struct {
	repo PaymentsRepo
}

type Service interface {
	Create(data Payments) (Payments, error)
	GetById(id int, userId string) (Payments, error)
	UpdateStatusAndMethod(id int, status, method string) error
	BookingReport() ([]BookingsReport, error)
	GetAllPayments(userId string) ([]Payments, error)
}

func NewPaymentsService(repo PaymentsRepo) Service {
	return &PaymentsService{
		repo: repo,
	}
}

func (s PaymentsService) Create(data Payments) (Payments, error) {
	return s.repo.Create(data)
}
func (s PaymentsService) GetById(id int, userId string) (Payments, error) {
	return s.repo.GetById(id, userId)
}
func (s PaymentsService) UpdateStatusAndMethod(id int, status, method string) error {
	return s.repo.UpdateStatusAndMethod(id, status, method)
}
func (s PaymentsService) BookingReport() ([]BookingsReport, error) {
	return s.repo.BookingReport()
}

func (s PaymentsService) GetAllPayments(userId string) ([]Payments, error) {
	return s.repo.GetAll(userId)
}
