package equipments_service

type EquipmentsService struct {
	repo EquipmentsRepo
}

type Service interface {
	CreateEquipment(data Equipments) (Equipments, error)
	GetAllEquipments() ([]Equipments, error)
	GetEquipmentById(id int) (Equipments, error)
	UpdateEquipment(id int, data Equipments) (Equipments, error)
	DeleteEquipment(id int) error
}

func NewEquipmentsService(repo EquipmentsRepo) Service {
	return &EquipmentsService{
		repo: repo,
	}
}

func (s EquipmentsService) CreateEquipment(data Equipments) (Equipments, error) {
	return s.repo.Create(data)
}
func (s EquipmentsService) GetAllEquipments() ([]Equipments, error) {
	return s.repo.GetAll()
}
func (s EquipmentsService) GetEquipmentById(id int) (Equipments, error) {
	return s.repo.GetById(id)
}
func (s EquipmentsService) UpdateEquipment(id int, data Equipments) (Equipments, error) {
	return s.repo.Update(id, data)
}
func (s EquipmentsService) DeleteEquipment(id int) error {
	return s.repo.Delete(id)
}
