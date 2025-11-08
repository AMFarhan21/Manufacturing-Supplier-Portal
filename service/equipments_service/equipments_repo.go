package equipments_service

import "Manufacturing-Supplier-Portal/model"

type EquipmentsRepo interface {
	Create(data model.Equipments) (model.Equipments, error)
	GetAll() ([]model.Equipments, error)
	GetById(id int) (model.Equipments, error)
	Update(id int, data model.Equipments) (model.Equipments, error)
	Delete(id int) error
	UpdateStatus(id int, status bool) error
}
