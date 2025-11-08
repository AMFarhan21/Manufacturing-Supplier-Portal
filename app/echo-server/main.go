package main

import (
	"Manufacturing-Supplier-Portal/app/echo-server/controller"
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

	_ "Manufacturing-Supplier-Portal/docs"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8000
// @BasePath /api

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description type "Bearer" followed by a space and your JWT token

func main() {
	db := database.GetDatabaseConnection()

	secret := os.Getenv("SECRET")

	e := echo.New()
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.Use(middleware.Recover())

	xenditApi := os.Getenv("XENDIT_API")
	xenditUrl := os.Getenv("XENDIT_URL")
	xenditWebhookUrl := os.Getenv("XENDIT_WEBHOOK_URL")
	successRedirectUrl := os.Getenv("REDIRECT_URL")
	failureRedirectUrl := os.Getenv("REDIRECT_URL")

	xenditRepository := xendit.NewXenditRepository(xenditApi, xenditUrl, xenditWebhookUrl, successRedirectUrl, failureRedirectUrl)

	usersRepository := users_repository.NewUsersGormRepository(db)
	usersService := users_service.NewUsersService(usersRepository, xenditRepository, secret)
	usersController := controller.NewUsersController(usersService)

	equipmentsRepository := equipments_repository.NewEquipmentsGormRepository(db)
	equipmentsService := equipments_service.NewEquipmentsService(equipmentsRepository)
	equipmentsController := controller.NewEquipmentsController(equipmentsService)

	paymentsRepository := payments_repository.NewPaymentsGormRepository(db)
	paymentService := payments_service.NewPaymentsService(paymentsRepository)
	paymentsController := controller.NewPaymentsController(paymentService)

	rentalHistoriesRepository := rental_histories_repository.NewRentalHistoriesGormRepository(db)

	rentalsRepository := rentals_repository.NewRentalsGormRepository(db)
	rentalsService := rentals_service.NewRentalsService(rentalsRepository, equipmentsRepository, xenditRepository, paymentsRepository, rentalHistoriesRepository, usersRepository)
	rentalsController := controller.NewRentalsController(rentalsService)

	webHookController := controller.NewWebhookController(paymentService, rentalsService, usersService)
	router.Router(e, secret, usersController, equipmentsController, rentalsController, webHookController, paymentsController)

	fmt.Println("Successfully connected to the server!")
	e.Logger.Fatal(e.Start(":8000"))
}
