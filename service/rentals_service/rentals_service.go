package rentals_service

import (
	"Manufacturing-Supplier-Portal/service/equipments_service"
	"Manufacturing-Supplier-Portal/service/payments_service"
	"Manufacturing-Supplier-Portal/service/rental_histories_service"
	"Manufacturing-Supplier-Portal/service/users_service"
	"Manufacturing-Supplier-Portal/service/xendit_service"
	"errors"
	"log"
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
	CreateRental(data Rentals) (RentalsWithInvoiceUrl, error)
	createPayment(rental Rentals) (payments_service.Payments, error)
	rentalWithInvoiceUrl(rentalEquipmentUser RentalEquipmentUser, rental Rentals, paymentId int) (RentalsWithInvoiceUrl, error)
	UpdateStatusAndDate(paymentId int, userId, status string) error
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

func (s RentalsService) CreateRental(data Rentals) (RentalsWithInvoiceUrl, error) {
	now := time.Now()
	data.CreatedAt = now

	equipment, err := s.equipmentRepo.GetById(data.EquipmentId)
	if err != nil {
		return RentalsWithInvoiceUrl{}, err
	}

	if !*equipment.Available {
		return RentalsWithInvoiceUrl{}, errors.New("this equipment is not available")
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

	user, err := s.userRepo.FindById(data.UserId)
	if err != nil {
		return RentalsWithInvoiceUrl{}, err
	}
	if user.DepositAmount < price {
		return RentalsWithInvoiceUrl{}, errors.New("you dont have enough money to rent this equipment")
	}

	data.Price = price
	data.Status = "PENDING"

	rental, err := s.rentalRepo.Create(data)
	if err != nil {
		return RentalsWithInvoiceUrl{}, err
	}
	_, err = s.rentalHistoriesRepo.CreateRentalHistory(rental_histories_service.RentalHistories{
		RentalId:  rental.Id,
		UserId:    rental.UserId,
		Status:    rental.Status,
		CreatedAt: rental.CreatedAt,
	})
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
	payment, err := s.paymentsRepo.Create(dataPayments)
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
	var startDate string
	var endDate string

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

		now := time.Now().Add(time.Hour * 24)
		startDate = now.Format("2006-01-02")
		endDate = now.Add(time.Hour * period).Format("2006-01-02")
		err := s.equipmentRepo.UpdateStatus(rental.EquipmentId, false)
		if err != nil {
			return err
		}

		_, err = s.rentalHistoriesRepo.CreateRentalHistory(rental_histories_service.RentalHistories{
			RentalId:  rental.RentalId,
			UserId:    userId,
			Status:    status,
			CreatedAt: rental.CreatedAt,
		})
		if err != nil {
			return err
		}

		// search for users
		user, err := s.userRepo.FindById(userId)
		if err != nil {
			return err
		}

		// substract user deposit
		amount := user.DepositAmount - payment.Amount

		// update user deposit amount
		_, err = s.userRepo.UpdateDepositAmount(userId, amount)
		if err != nil {
			return err
		}

	case "CANCELLED":
		_, err = s.rentalHistoriesRepo.CreateRentalHistory(rental_histories_service.RentalHistories{
			RentalId:  rental.RentalId,
			UserId:    userId,
			Status:    status,
			CreatedAt: rental.CreatedAt,
		})
		if err != nil {
			return err
		}
	}

	log.Printf("------------------------------------------%s-----------------------------------", "Rentals Service")
	log.Println("STATUS:", status)
	log.Println("PAYMENT_ID:", paymentId)
	log.Println("USER_ID:", userId)
	log.Println("START_DATE:", startDate)
	log.Println("END_DATE:", endDate)
	log.Println("RENTAL_ID:", payment.RentalId)
	log.Println("RentalId:", rental.RentalId)
	log.Printf("------------------------------------------%s-----------------------------------", "Rentals Service")
	return s.rentalRepo.UpdateStatusAndDateRepo(payment.RentalId, status, startDate, endDate)
}
