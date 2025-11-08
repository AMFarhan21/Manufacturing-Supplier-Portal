package controller

import (
	"Manufacturing-Supplier-Portal/service/payments_service"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/AMFarhan21/fres"
	"github.com/labstack/echo/v4"
)

type PaymentsController struct {
	service payments_service.Service
}

func NewPaymentsController(service payments_service.Service) *PaymentsController {
	return &PaymentsController{
		service: service,
	}
}

func (ctrl PaymentsController) GetPaymentsById(c echo.Context) error {
	param := c.Param("id")
	id, _ := strconv.Atoi(param)
	userId := c.Get("id").(string)
	payment, err := ctrl.service.GetById(id, userId)
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			log.Print("Error on GetPaymentsById request:", err.Error())
			return c.JSON(http.StatusNotFound, fres.Response.StatusNotFound("Cant find payment with that id and userId"))
		}
		log.Print("Error on GetPaymentsById server:", err.Error())
		return c.JSON(http.StatusInternalServerError, fres.Response.StatusInternalServerError(http.StatusInternalServerError))
	}

	log.Print("Successfully get payment by id")
	return c.JSON(http.StatusOK, fres.Response.StatusOK(payment))
}

func (ctrl PaymentsController) BookingReport(c echo.Context) error {
	bookingReport, err := ctrl.service.BookingReport()
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			log.Print("Error on BookingReport request", err.Error())
			return c.JSON(http.StatusNotFound, fres.Response.StatusNotFound(err.Error()))
		}
		log.Print("Error on BookingReport server:", err.Error())
		return c.JSON(http.StatusInternalServerError, fres.Response.StatusInternalServerError(http.StatusInternalServerError))
	}

	log.Print("Successfully get booking report")
	return c.JSON(http.StatusOK, fres.Response.StatusOK(bookingReport))
}

func (ctrl PaymentsController) GetAllPaymentsByUserId(c echo.Context) error {
	userId := c.Get("id").(string)
	payments, err := ctrl.service.GetAllPayments(userId)

	if err != nil {
		log.Print("Error on GetAllPaymentsByUserId server:", err.Error())
		return c.JSON(http.StatusInternalServerError, fres.Response.StatusInternalServerError(http.StatusInternalServerError))
	}

	log.Print("Successfully get all payments")
	return c.JSON(http.StatusOK, fres.Response.StatusOK(payments))
}
