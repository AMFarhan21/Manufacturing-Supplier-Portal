package main

import (
	"Manufacturing-Supplier-Portal/app/echo-server/controller/equipments_controller"
	"Manufacturing-Supplier-Portal/app/echo-server/controller/payments_controller"
	"Manufacturing-Supplier-Portal/app/echo-server/controller/rentals_controller"
	"Manufacturing-Supplier-Portal/app/echo-server/controller/users_controller"
	"Manufacturing-Supplier-Portal/app/echo-server/controller/webhook_controller"
	"Manufacturing-Supplier-Portal/app/echo-server/router"
	"Manufacturing-Supplier-Portal/repository/equipments_repository"
	"Manufacturing-Supplier-Portal/repository/payments_repository"
	"Manufacturing-Supplier-Portal/repository/rental_histories_repository"
	"Manufacturing-Supplier-Portal/repository/rentals_repository"
	"Manufacturing-Supplier-Portal/repository/users_repository"
	"Manufacturing-Supplier-Portal/repository/xendit"
	"Manufacturing-Supplier-Portal/service/equipments_service"
	"Manufacturing-Supplier-Portal/service/payments_service"
	"Manufacturing-Supplier-Portal/service/rentals_service"
	"Manufacturing-Supplier-Portal/service/users_service"
	"Manufacturing-Supplier-Portal/utils/database"
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	db := database.GetDatabaseConnection()

	secret := os.Getenv("SECRET")

	e := echo.New()

	xenditApi := os.Getenv("XENDIT_API")
	xenditUrl := os.Getenv("XENDIT_URL")
	xenditWebhookUrl := os.Getenv("XENDIT_WEBHOOK_URL")
	successRedirectUrl := os.Getenv("REDIRECT_URL")
	failureRedirectUrl := os.Getenv("REDIRECT_URL")

	xenditRepository := xendit.NewXenditRepository(xenditApi, xenditUrl, xenditWebhookUrl, successRedirectUrl, failureRedirectUrl)

	usersRepository := users_repository.NewUsersGormRepository(db)
	usersService := users_service.NewUsersService(usersRepository, secret)
	usersController := users_controller.NewUsersController(usersService)

	equipmentsRepository := equipments_repository.NewEquipmentsGormRepository(db)
	equipmentsService := equipments_service.NewEquipmentsService(equipmentsRepository)
	equipmentsController := equipments_controller.NewEquipmentsController(equipmentsService)

	paymentsRepository := payments_repository.NewPaymentsGormRepository(db)
	paymentService := payments_service.NewPaymentsService(paymentsRepository)
	paymentsController := payments_controller.NewPaymentsController(paymentService)

	rentalHistoriesRepository := rental_histories_repository.NewRentalHistoriesGormRepository(db)

	rentalsRepository := rentals_repository.NewRentalsGormRepository(db)
	rentalsService := rentals_service.NewRentalsService(rentalsRepository, equipmentsRepository, xenditRepository, paymentsRepository, rentalHistoriesRepository)
	rentalsController := rentals_controller.NewRentalsController(rentalsService)

	webHookController := webhook_controller.NewWebhookController(paymentService, rentalsService)
	router.Router(e, secret, usersController, equipmentsController, rentalsController, webHookController, paymentsController)

	fmt.Println("Successfully connected to the server!")
	e.Logger.Fatal(e.Start(":8000"))
}
