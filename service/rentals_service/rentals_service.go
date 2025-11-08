package rentals_service

import (
	"Manufacturing-Supplier-Portal/model"
	"Manufacturing-Supplier-Portal/service/equipments_service"
	"Manufacturing-Supplier-Portal/service/payments_service"
	"Manufacturing-Supplier-Portal/service/rental_histories_service"
	"Manufacturing-Supplier-Portal/service/users_service"
	"Manufacturing-Supplier-Portal/service/xendit_service"
	"errors"
	"fmt"
	"time"
)

type RentalsService struct {
	rentalRepo          RentalsRepo
	equipmentRepo       equipments_service.EquipmentsRepo
	xenditRepo          xendit_service.XenditRepo
	paymentsRepo        payments_service.PaymentsRepo
	rentalHistoriesRepo rental_histories_service.RentalHistoriesRepo
	userRepo            users_service.UsersRepo
}

type Service interface {
	CreateRental(data model.Rentals, paymentMethod string) (model.RentalsResponse, error)
	createPayment(rental model.Rentals, status string) (model.Payments, error)
	rentalWithInvoiceUrl(rentalEquipmentUser model.RentalEquipmentUser, rental model.Rentals, paymentId int) (model.RentalsResponse, error)
	UpdateStatusAndDate(paymentId int, userId, status string) error
	GetAllRentalHistoriesByUserId(userId string) ([]model.RentalHistories, error)
	SimulateAutomaticUpdateRentalStatus() error
}

func NewRentalsService(
	rentalRepo RentalsRepo,
	equipmentRepo equipments_service.EquipmentsRepo,
	xenditRepo xendit_service.XenditRepo,
	paymentsRepo payments_service.PaymentsRepo,
	rentalHistoriesRepo rental_histories_service.RentalHistoriesRepo,
	userRepo users_service.UsersRepo,
) Service {
	return &RentalsService{
		rentalRepo:          rentalRepo,
		equipmentRepo:       equipmentRepo,
		xenditRepo:          xenditRepo,
		paymentsRepo:        paymentsRepo,
		rentalHistoriesRepo: rentalHistoriesRepo,
		userRepo:            userRepo,
	}
}

func (s RentalsService) CreateRental(data model.Rentals, paymentMethod string) (model.RentalsResponse, error) {

	equipment, err := s.equipmentRepo.GetById(data.EquipmentId)
	if err != nil {
		return model.RentalsResponse{}, err
	}
	if !*equipment.Available {
		return model.RentalsResponse{}, errors.New("this equipment is not available")
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

	now := time.Now()
	data.CreatedAt = now
	if paymentMethod == "wallet" {
		user, err := s.userRepo.FindById(data.UserId)
		if err != nil {
			return model.RentalsResponse{}, err
		}
		if user.DepositAmount < price {
			return model.RentalsResponse{}, errors.New("you dont have enough money to rent this equipment")
		}

		data.Price = price
		data.Status = "BOOKED"
		rental, err := s.rentalRepo.Create(data)
		if err != nil {
			return model.RentalsResponse{}, err
		}

		amount := user.DepositAmount - rental.Price
		_, err = s.userRepo.UpdateDepositAmount(rental.UserId, amount)
		if err != nil {
			return model.RentalsResponse{}, err
		}

		_, err = s.rentalHistoriesRepo.CreateRentalHistory(model.RentalHistories{
			RentalId:  rental.Id,
			UserId:    rental.UserId,
			Status:    rental.Status,
			CreatedAt: rental.CreatedAt,
		})
		if err != nil {
			return model.RentalsResponse{}, err
		}

		_, err = s.createPayment(rental, "PAID")
		if err != nil {
			return model.RentalsResponse{}, err
		}

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
		start := now.Add(time.Minute * 1)
		end := now.Add(time.Hour * period)

		err = s.rentalRepo.UpdateStatusAndDateRepo(rental.Id, data.Status, start, end)
		if err != nil {
			return model.RentalsResponse{}, err
		}

		return model.RentalsResponse{
			Id:           rental.Id,
			UserId:       rental.UserId,
			EquipmentId:  rental.EquipmentId,
			RentalPeriod: rental.RentalPeriod,
			StartDate:    &start,
			EndDate:      &end,
			Price:        rental.Price,
			Status:       rental.Status,
			CreatedAt:    rental.CreatedAt,
		}, nil
	}
	data.Price = price
	data.Status = "PENDING"

	rental, err := s.rentalRepo.Create(data)
	if err != nil {
		return model.RentalsResponse{}, err
	}
	_, err = s.rentalHistoriesRepo.CreateRentalHistory(model.RentalHistories{
		RentalId:  rental.Id,
		UserId:    rental.UserId,
		Status:    rental.Status,
		CreatedAt: rental.CreatedAt,
	})
	if err != nil {
		return model.RentalsResponse{}, err
	}

	payment, err := s.createPayment(rental, "PENDING")
	if err != nil {
		return model.RentalsResponse{}, err
	}

	rentalEquipmentUser, err := s.rentalRepo.GetRentalById(rental.Id)
	if err != nil {
		return model.RentalsResponse{}, err
	}

	rentalWithInvoice, err := s.rentalWithInvoiceUrl(rentalEquipmentUser, rental, payment.Id)
	if err != nil {
		return model.RentalsResponse{}, err
	}

	return rentalWithInvoice, nil
}

func (s RentalsService) createPayment(rental model.Rentals, status string) (model.Payments, error) {
	var dataPayments model.Payments
	dataPayments.UserId = rental.UserId
	dataPayments.RentalId = rental.Id
	dataPayments.Amount = rental.Price
	dataPayments.Status = status
	dataPayments.CreatedAt = time.Now()
	payment, err := s.paymentsRepo.Create(dataPayments)
	if err != nil {
		return model.Payments{}, err
	}

	return payment, nil
}

func (s RentalsService) rentalWithInvoiceUrl(rentalEquipmentUser model.RentalEquipmentUser, rental model.Rentals, paymentId int) (model.RentalsResponse, error) {
	invoiceUrl, err := s.xenditRepo.XenditInvoiceUrl(rentalEquipmentUser.UserId, "PAYMENT", rentalEquipmentUser.Username, rentalEquipmentUser.Email, rentalEquipmentUser.EquipmentName, rentalEquipmentUser.Category, paymentId, rental.Price)
	if err != nil || invoiceUrl == "" {
		return model.RentalsResponse{}, fmt.Errorf("failed to create invoice URL, try again:%v", err)
	}
	var rentalWithInvoiceUrl model.RentalsResponse
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

func (s RentalsService) UpdateStatusAndDate(paymentId int, userId, status string) error {
	payment, err := s.paymentsRepo.GetById(paymentId, userId)
	if err != nil {
		return err
	}

	rental, err := s.rentalRepo.GetRentalById(payment.RentalId)
	if err != nil {
		return err
	}

	var period time.Duration
	var startDate time.Time
	var endDate time.Time

	switch status {
	case "BOOKED":
		switch rental.RentalPeriod {
		case "day":
			period = 24
		case "week":
			period = 24 * 7
		case "month":
			period = 24 * 30
		case "year":
			period = 24 * 365
		}

		now := time.Now().Add(time.Minute * 1)
		// startDate = now.Format("2006-01-02")
		// endDate = now.Add(time.Hour * period).Format("2006-01-02")
		startDate = now
		endDate = now.Add(time.Hour * period)
		err := s.equipmentRepo.UpdateStatus(rental.EquipmentId, false)
		if err != nil {
			return err
		}

		_, err = s.rentalHistoriesRepo.CreateRentalHistory(model.RentalHistories{
			RentalId:  rental.RentalId,
			UserId:    userId,
			Status:    status,
			CreatedAt: rental.CreatedAt,
		})
		if err != nil {
			return err
		}

	case "CANCELLED":
		_, err = s.rentalHistoriesRepo.CreateRentalHistory(model.RentalHistories{
			RentalId:  rental.RentalId,
			UserId:    userId,
			Status:    status,
			CreatedAt: rental.CreatedAt,
		})
		if err != nil {
			return err
		}
	}

	return s.rentalRepo.UpdateStatusAndDateRepo(payment.RentalId, status, startDate, endDate)
}

func (s RentalsService) GetAllRentalHistoriesByUserId(userId string) ([]model.RentalHistories, error) {
	return s.rentalHistoriesRepo.GetAll(userId)
}

func (s RentalsService) SimulateAutomaticUpdateRentalStatus() error {
	return s.rentalRepo.SimulateAutomaticUpdateRentalStatus()
}
