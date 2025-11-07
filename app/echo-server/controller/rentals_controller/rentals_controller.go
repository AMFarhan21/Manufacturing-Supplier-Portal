package rentals_controller

import (
	"Manufacturing-Supplier-Portal/service/rentals_service"
	"log"
	"net/http"
	"strings"

	"github.com/AMFarhan21/fres"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type (
	RentalsController struct {
		service  rentals_service.Service
		validate *validator.Validate
	}

	RentalsInput struct {
		EquipmentId  int    `json:"equipment_id" validate:"required"`
		RentalPeriod string `json:"rental_period" validate:"required,oneof=day week month year"`
	}
)

func NewRentalsController(service rentals_service.Service) *RentalsController {
	return &RentalsController{
		service:  service,
		validate: validator.New(),
	}
}

func (ctrl RentalsController) CreateRental(c echo.Context) error {
	userId := c.Get("id").(string)

	var request RentalsInput

	if err := c.Bind(&request); err != nil {
		log.Print("Error on CreateRental request body:", err.Error())
		return c.JSON(http.StatusBadRequest, fres.Response.StatusBadRequest(err.Error()))
	}

	if err := ctrl.validate.Struct(request); err != nil {
		log.Print("Error on CreateRental validation:", err.Error())
		return c.JSON(http.StatusBadRequest, fres.Response.StatusBadRequest(err.Error()))
	}

	rental, err := ctrl.service.CreateRental(rentals_service.Rentals{
		UserId:       userId,
		EquipmentId:  request.EquipmentId,
		RentalPeriod: request.RentalPeriod,
	})
	if err != nil {
		if strings.Contains(err.Error(), "you dont have enough money to rent this equipment") {
			log.Print("Error on CreateRental request body:", err.Error())
			return c.JSON(http.StatusBadRequest, fres.DefaultErrorResponse{
				Success: false,
				Message: err.Error(),
				Error:   http.StatusBadRequest,
			})
		}

		if strings.Contains(err.Error(), "this equipment is not available") {
			log.Print("Error on CreateRental request body:", err.Error())
			return c.JSON(http.StatusBadRequest, fres.DefaultErrorResponse{
				Success: false,
				Message: err.Error(),
				Error:   http.StatusBadRequest,
			})
		}
		log.Print("Error on create rental server:", err.Error())
		return c.JSON(http.StatusInternalServerError, fres.Response.StatusInternalServerError(http.StatusInternalServerError))
	}

	log.Print("Successfully create a rental")
	return c.JSON(http.StatusCreated, fres.Response.StatusCreated(rental))
}
