package router

import (
	"Manufacturing-Supplier-Portal/app/echo-server/controller/equipments_controller"
	"Manufacturing-Supplier-Portal/app/echo-server/controller/payments_controller"
	"Manufacturing-Supplier-Portal/app/echo-server/controller/rentals_controller"
	"Manufacturing-Supplier-Portal/app/echo-server/controller/users_controller"
	"Manufacturing-Supplier-Portal/app/echo-server/controller/webhook_controller"
	"Manufacturing-Supplier-Portal/app/echo-server/middleware"
	"net/http"

	"github.com/AMFarhan21/fres"
	"github.com/labstack/echo/v4"
)

func Router(
	e *echo.Echo,
	jwtSecret string,
	usersController *users_controller.UsersController,
	equipmentsController *equipments_controller.EquipmentsController,
	rentalsController *rentals_controller.RentalsController,
	webHookController *webhook_controller.WebhookController,
	paymentsController *payments_controller.PaymentsController,
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
		return c.JSON(http.StatusOK, fres.Response.StatusOK("Manufacture-Supplier-Portal"))
	})

	auth := e.Group("api/auth")
	auth.POST("/register", usersController.RegisterUser)
	auth.POST("/login", usersController.LoginUser)

	users := e.Group("api/users", middlewares)
	users.GET("/me", usersController.GetUserLogin, userAccess)
	users.GET("/list", usersController.GetAllUsers, adminAccess)
	users.POST("/topup", usersController.TopUpDeposit, userAccess)

	equipments := e.Group("api/equipments", middlewares)
	equipments.GET("", equipmentsController.GetAllEquipments, userAccess)
	equipments.GET("/:id", equipmentsController.GetEquipmentById, userAccess)
	equipments.POST("", equipmentsController.CreateEquipment, adminAccess)
	equipments.PUT("/:id", equipmentsController.UpdateEquipment, adminAccess)
	equipments.DELETE("/:id", equipmentsController.DeleteEquipment, adminAccess)

	payments := e.Group("/api/payments", middlewares)
	payments.GET("/:id", paymentsController.GetPaymentsById, userAccess)

	e.POST("/webhook/handler", webHookController.HandleWebhook)

	rentals := e.Group("api/rentals", middlewares)
	rentals.POST("", rentalsController.CreateRental, userAccess)
}
