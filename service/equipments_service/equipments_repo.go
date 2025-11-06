package equipments_service

type EquipmentsRepo interface {
	Create(data Equipments) (Equipments, error)
	GetAll() ([]Equipments, error)
	GetById(id int) (Equipments, error)
	Update(id int, data Equipments) (Equipments, error)
	Delete(id int) error
	UpdateStatus(id int, status bool) error
}
