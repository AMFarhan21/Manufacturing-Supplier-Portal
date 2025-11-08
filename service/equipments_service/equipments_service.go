package equipments_service

import "Manufacturing-Supplier-Portal/model"

type EquipmentsService struct {
	repo EquipmentsRepo
}

type Service interface {
	CreateEquipment(data model.Equipments) (model.Equipments, error)
	GetAllEquipments() ([]model.Equipments, error)
	GetEquipmentById(id int) (model.Equipments, error)
	UpdateEquipment(id int, data model.Equipments) (model.Equipments, error)
	DeleteEquipment(id int) error
}

func NewEquipmentsService(repo EquipmentsRepo) Service {
	return &EquipmentsService{
		repo: repo,
	}
}

func (s EquipmentsService) CreateEquipment(data model.Equipments) (model.Equipments, error) {
	return s.repo.Create(data)
}
func (s EquipmentsService) GetAllEquipments() ([]model.Equipments, error) {
	return s.repo.GetAll()
}
func (s EquipmentsService) GetEquipmentById(id int) (model.Equipments, error) {
	return s.repo.GetById(id)
}
func (s EquipmentsService) UpdateEquipment(id int, data model.Equipments) (model.Equipments, error) {
	return s.repo.Update(id, data)
}
func (s EquipmentsService) DeleteEquipment(id int) error {
	return s.repo.Delete(id)
}
