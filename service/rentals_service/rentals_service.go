package rentals_service

import (
	"Manufacturing-Supplier-Portal/service/equipments_service"
	"Manufacturing-Supplier-Portal/service/payments_service"
	"Manufacturing-Supplier-Portal/service/xendit_service"
	"errors"
	"time"
)

type RentalsService struct {
	rentalRepo         RentalsRepo
	equipmentRepo      equipments_service.EquipmentsRepo
	xenditRepo         xendit_service.XenditRepo
	paymentsRepository payments_service.PaymentsRepo
}

type Service interface {
	CreateRental(data Rentals) (RentalsWithInvoiceUrl, error)
	createPayment(rental Rentals) (payments_service.Payments, error)
	rentalWithInvoiceUrl(rentalEquipmentUser RentalEquipmentUser, rental Rentals, paymentId int) (RentalsWithInvoiceUrl, error)
}

func NewRentalsService(rentalRepo RentalsRepo, equipmentRepo equipments_service.EquipmentsRepo, xenditRepo xendit_service.XenditRepo, paymentsRepository payments_service.PaymentsRepo) Service {
	return &RentalsService{
		rentalRepo:         rentalRepo,
		equipmentRepo:      equipmentRepo,
		xenditRepo:         xenditRepo,
		paymentsRepository: paymentsRepository,
	}
}

func (s RentalsService) CreateRental(data Rentals) (RentalsWithInvoiceUrl, error) {
	var period time.Duration

	switch data.RentalPeriod {
	case "day":
		period = 24
	case "week":
		period = 24 * 7
	case "month":
		period = 24 * 30
	case "year":
		period = 24 * 365
	}

	now := time.Now()

	data.StartDate = now.Format("2006-01-02")
	data.EndDate = now.Add(time.Hour * period).Format("2006-01-02")
	data.CreatedAt = now

	equipment, err := s.equipmentRepo.GetById(data.EquipmentId)
	if err != nil {
		return RentalsWithInvoiceUrl{}, err
	}

	var price float64
	switch data.RentalPeriod {
	case "day":
		price = equipment.PricePerDay
	case "week":
		price = equipment.PricePerWeek
	case "month":
		price = equipment.PricePerMonth
	case "year":
		price = equipment.PricePerYear
	}

	data.Price = price
	data.Status = "pending"

	rental, err := s.rentalRepo.Create(data)
	if err != nil {
		return RentalsWithInvoiceUrl{}, err
	}

	payment, err := s.createPayment(rental)
	if err != nil {
		return RentalsWithInvoiceUrl{}, err
	}

	rentalEquipmentUser, err := s.rentalRepo.GetRentalById(rental.Id)
	if err != nil {
		return RentalsWithInvoiceUrl{}, err
	}

	rentalWithInvoice, err := s.rentalWithInvoiceUrl(rentalEquipmentUser, rental, payment.Id)
	if err != nil {
		return RentalsWithInvoiceUrl{}, err
	}

	return rentalWithInvoice, nil
}

func (s RentalsService) createPayment(rental Rentals) (payments_service.Payments, error) {
	var dataPayments payments_service.Payments
	dataPayments.UserId = rental.UserId
	dataPayments.RentalId = rental.Id
	dataPayments.Amount = rental.Price
	dataPayments.Status = "PENDING"
	dataPayments.CreatedAt = time.Now()
	payment, err := s.paymentsRepository.Create(dataPayments)
	if err != nil {
		return payments_service.Payments{}, err
	}

	return payment, nil
}

func (s RentalsService) rentalWithInvoiceUrl(rentalEquipmentUser RentalEquipmentUser, rental Rentals, paymentId int) (RentalsWithInvoiceUrl, error) {
	invoiceUrl, err := s.xenditRepo.XenditInvoiceUrl(rentalEquipmentUser.UserId, rentalEquipmentUser.Description, rentalEquipmentUser.Username, rentalEquipmentUser.Email, rentalEquipmentUser.EquipmentName, rentalEquipmentUser.Category, paymentId, rental.Price)
	if err != nil || invoiceUrl == "" {
		return RentalsWithInvoiceUrl{}, errors.New("failed to create invoice URL, try again")
	}
	var rentalWithInvoiceUrl RentalsWithInvoiceUrl
	rentalWithInvoiceUrl.Id = rental.Id
	rentalWithInvoiceUrl.UserId = rental.UserId
	rentalWithInvoiceUrl.EquipmentId = rental.EquipmentId
	rentalWithInvoiceUrl.RentalPeriod = rental.RentalPeriod
	rentalWithInvoiceUrl.StartDate = rental.StartDate
	rentalWithInvoiceUrl.EndDate = rental.EndDate
	rentalWithInvoiceUrl.Price = rental.Price
	rentalWithInvoiceUrl.Status = rental.Status
	rentalWithInvoiceUrl.CreatedAt = rental.CreatedAt
	rentalWithInvoiceUrl.InvoiceUrl = invoiceUrl

	return rentalWithInvoiceUrl, nil
}
