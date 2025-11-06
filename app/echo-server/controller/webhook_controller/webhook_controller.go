package webhook_controller

import (
	"Manufacturing-Supplier-Portal/service/payments_service"
	"Manufacturing-Supplier-Portal/service/rentals_service"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/AMFarhan21/fres"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type (
	WebhookController struct {
		paymentService payments_service.Service
		rentalService  rentals_service.Service
		validate       *validator.Validate
	}

	WebhookRequest struct {
		ID                     string    `json:"id"`
		ExternalID             string    `json:"external_id"`
		UserID                 string    `json:"user_id"`
		IsHigh                 bool      `json:"is_high"`
		PaymentMethod          string    `json:"payment_method"`
		Status                 string    `json:"status"`
		MerchantName           string    `json:"merchant_name"`
		Amount                 int       `json:"amount"`
		PaidAmount             int       `json:"paid_amount"`
		BankCode               string    `json:"bank_code"`
		PaidAt                 time.Time `json:"paid_at"`
		PayerEmail             string    `json:"payer_email"`
		Description            string    `json:"description"`
		AdjustedReceivedAmount int       `json:"adjusted_received_amount"`
		FeesPaidAmount         int       `json:"fees_paid_amount"`
		Updated                time.Time `json:"updated"`
		Created                time.Time `json:"created"`
		Currency               string    `json:"currency"`
		PaymentChannel         string    `json:"payment_channel"`
		PaymentDestination     string    `json:"payment_destination"`
	}
)

func NewWebhookController(paymentService payments_service.Service, rentalService rentals_service.Service) *WebhookController {
	return &WebhookController{
		paymentService: paymentService,
		rentalService:  rentalService,
		validate:       validator.New(),
	}
}

func (ctrl WebhookController) HandleWebhook(c echo.Context) error {
	var request WebhookRequest

	if err := c.Bind(&request); err != nil {
		log.Println("Failed to bind webhook request:", err)
		return c.JSON(http.StatusBadRequest, fres.Response.StatusBadRequest("Invalid request"))
	}

	log.Print("Received webhook from Xendit:", request)

	paymentIDandUserID := strings.Split(request.ExternalID, "-")

	paymentId, _ := strconv.Atoi(paymentIDandUserID[0])
	userId := paymentIDandUserID[1]

	if request.Status == "PAID" {
		err := ctrl.paymentService.UpdateStatusAndMethod(paymentId, request.Status, request.PaymentMethod)
		if err != nil {
			log.Println("Failed to update payment status:", err.Error())
			return c.JSON(http.StatusInternalServerError, fres.Response.StatusInternalServerError(http.StatusInternalServerError))
		}

		err = ctrl.rentalService.UpdateStatusAndDate(paymentId, userId, "ACTIVE")
		if err != nil {
			log.Println("Failed to update payment status:", err.Error())
			return c.JSON(http.StatusInternalServerError, fres.Response.StatusInternalServerError(http.StatusInternalServerError))
		}
	}

	if request.Status == "EXPIRED" {
		err := ctrl.paymentService.UpdateStatusAndMethod(paymentId, request.Status, request.PaymentMethod)
		if err != nil {
			log.Println("Failed to update payment status:", err.Error())
			return c.JSON(http.StatusInternalServerError, fres.Response.StatusInternalServerError(http.StatusInternalServerError))
		}

		err = ctrl.rentalService.UpdateStatusAndDate(paymentId, userId, "CANCELLED")
		if err != nil {
			log.Println("Failed to update payment status:", err.Error())
			return c.JSON(http.StatusInternalServerError, fres.Response.StatusInternalServerError(http.StatusInternalServerError))
		}
	}

	log.Print(request)
	return c.JSON(http.StatusOK, fres.Response.StatusOK(http.StatusOK))
}
