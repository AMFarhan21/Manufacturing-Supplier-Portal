package payments_service

type PaymentsService struct {
	repo PaymentsRepo
}

type Service interface {
	Create(data Payments) (Payments, error)
	GetById(id int, userId string) (Payments, error)
	UpdateStatus(id int, userId, status string) error
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
func (s PaymentsService) UpdateStatus(id int, userId, status string) error {
	return s.repo.UpdateStatus(id, userId, status)
}
