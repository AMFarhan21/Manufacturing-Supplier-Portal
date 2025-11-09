package router

import (
	"Manufacturing-Supplier-Portal/app/echo-server/controller"
	"Manufacturing-Supplier-Portal/app/echo-server/middleware"
	"log"
	"net/http"

	"github.com/AMFarhan21/fres"
	"github.com/labstack/echo/v4"
)

func Router(
	e *echo.Echo,
	jwtSecret string,
	usersController *controller.UsersController,
	equipmentsController *controller.EquipmentsController,
	rentalsController *controller.RentalsController,
	webHookController *controller.WebhookController,
	paymentsController *controller.PaymentsController,
) {
	middlewares := middleware.JWTMiddleware(jwtSecret)
	adminAccess := middleware.ACLMiddleware(map[string]bool{
		"admin": true,
	})
	userAccess := middleware.ACLMiddleware(map[string]bool{
		"admin": true,
		"user":  true,
	})

	e.GET("/", func(c echo.Context) error {
		log.Print("HELLO WORLD")
		return c.JSON(http.StatusOK, fres.Response.StatusOK("Manufacture-Supplier-Portal"))
	})

	auth := e.Group("api/auth")
	auth.POST("/register", usersController.RegisterUser)
	auth.POST("/login", usersController.LoginUser)

	users := e.Group("api/users", middlewares)
	users.GET("/me", usersController.GetUserLogin, userAccess)
	users.GET("/list", usersController.GetAllUsers, adminAccess)
	users.POST("/topup", usersController.TopUpDeposit, userAccess)
	users.GET("/ValidateEmailAddress", usersController.VerifiedEmail)

	equipments := e.Group("api/equipments", middlewares)
	equipments.GET("", equipmentsController.GetAllEquipments, userAccess)
	equipments.GET("/:id", equipmentsController.GetEquipmentById, userAccess)
	equipments.POST("", equipmentsController.CreateEquipment, adminAccess)
	equipments.PUT("/:id", equipmentsController.UpdateEquipment, adminAccess)
	equipments.DELETE("/:id", equipmentsController.DeleteEquipment, adminAccess)

	payments := e.Group("/api/payments", middlewares)
	payments.GET("", paymentsController.GetAllPaymentsByUserId, userAccess)
	payments.GET("/:id", paymentsController.GetPaymentsById, userAccess)
	payments.GET("/bookingreport", paymentsController.BookingReport, adminAccess)

	e.POST("/webhook/handler", webHookController.HandleWebhook)

	rentals := e.Group("api/rentals", middlewares)
	rentals.POST("", rentalsController.CreateRental, userAccess)
	rentals.GET("/history", rentalsController.GetAllRentalHistoriesByUserId, userAccess)
	rentals.GET("/refresh", rentalsController.SimulateAutomaticUpdateRentalStatus)
}
