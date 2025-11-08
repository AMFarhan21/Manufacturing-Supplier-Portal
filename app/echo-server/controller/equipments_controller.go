package controller

import (
	"Manufacturing-Supplier-Portal/model"
	"Manufacturing-Supplier-Portal/service/equipments_service"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/AMFarhan21/fres"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type (
	EquipmentsController struct {
		service  equipments_service.Service
		validate *validator.Validate
	}

	EquipmentsInput struct {
		Name          string  `json:"name" validate:"required"`
		CategoryId    int     `json:"category_id" validate:"required"`
		Description   string  `json:"description" validate:"required"`
		PricePerDay   float64 `json:"price_per_day" validate:"required"`
		PricePerWeek  float64 `json:"price_per_week" validate:"required"`
		PricePerMonth float64 `json:"price_per_month" validate:"required"`
		PricePerYear  float64 `json:"price_per_year" validate:"required"`
		Available     bool    `json:"available"`
	}
)

func NewEquipmentsController(service equipments_service.Service) *EquipmentsController {
	return &EquipmentsController{
		service:  service,
		validate: validator.New(),
	}
}

func (ctrl EquipmentsController) CreateEquipment(c echo.Context) error {
	var request EquipmentsInput

	if err := c.Bind(&request); err != nil {
		log.Print("Error on CreateEquipment request body:", err.Error())
		return c.JSON(http.StatusBadRequest, fres.Response.StatusBadRequest(err.Error()))
	}

	if err := ctrl.validate.Struct(request); err != nil {
		log.Print("Error on CreateEquipment validation:", err.Error())
		return c.JSON(http.StatusBadRequest, fres.Response.StatusBadRequest(err.Error()))
	}

	equipments, err := ctrl.service.CreateEquipment(model.Equipments{
		Name:          request.Name,
		CategoryId:    request.CategoryId,
		Description:   request.Description,
		PricePerDay:   request.PricePerDay,
		PricePerWeek:  request.PricePerWeek,
		PricePerMonth: request.PricePerMonth,
		PricePerYear:  request.PricePerYear,
		Available:     &request.Available,
	})

	if err != nil {
		log.Print("Error on create equipment server:", err.Error())
		return c.JSON(http.StatusInternalServerError, fres.Response.StatusInternalServerError(http.StatusInternalServerError))
	}

	log.Print("Successfully create a equipment")
	return c.JSON(http.StatusCreated, fres.Response.StatusCreated(equipments))
}

func (ctrl EquipmentsController) GetAllEquipments(c echo.Context) error {
	equipments, err := ctrl.service.GetAllEquipments()

	if err != nil {
		log.Print("Error on GetAllEquipments server:", err.Error())
		return c.JSON(http.StatusInternalServerError, fres.Response.StatusInternalServerError(http.StatusInternalServerError))
	}

	log.Print("Successfully get all equipments")
	return c.JSON(http.StatusOK, fres.Response.StatusOK(equipments))
}

func (ctrl EquipmentsController) GetEquipmentById(c echo.Context) error {
	param := c.Param("id")
	id, _ := strconv.Atoi(param)

	equipment, err := ctrl.service.GetEquipmentById(id)

	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			log.Print("Error on GetEquipmentById request", err.Error())
			return c.JSON(http.StatusNotFound, fres.Response.StatusNotFound(err.Error()))
		}
		log.Print("Error on GetEquipmentById server:", err.Error())
		return c.JSON(http.StatusInternalServerError, fres.Response.StatusInternalServerError(http.StatusInternalServerError))
	}

	log.Print("Successfully get equipment")
	return c.JSON(http.StatusOK, fres.Response.StatusOK(equipment))
}

func (ctrl EquipmentsController) UpdateEquipment(c echo.Context) error {
	param := c.Param("id")
	id, _ := strconv.Atoi(param)

	var request EquipmentsInput

	if err := c.Bind(&request); err != nil {
		log.Print("Error on UpdateEquipment request body:", err.Error())
		return c.JSON(http.StatusBadRequest, fres.Response.StatusBadRequest(err.Error()))
	}

	if err := ctrl.validate.Struct(request); err != nil {
		log.Print("Error on UpdateEquipment validation:", err.Error())
		return c.JSON(http.StatusBadRequest, fres.Response.StatusBadRequest(err.Error()))
	}

	equipments, err := ctrl.service.UpdateEquipment(id, model.Equipments{
		Name:          request.Name,
		CategoryId:    request.CategoryId,
		Description:   request.Description,
		PricePerDay:   request.PricePerDay,
		PricePerWeek:  request.PricePerWeek,
		PricePerMonth: request.PricePerMonth,
		PricePerYear:  request.PricePerYear,
		Available:     &request.Available,
	})

	if err != nil {
		if strings.Contains(err.Error(), "equipment id not found") {
			log.Print("Error on UpdateEquipment request", err.Error())
			return c.JSON(http.StatusNotFound, fres.Response.StatusNotFound(err.Error()))
		}

		log.Print("Error on UpdateEquipment server:", err.Error())
		return c.JSON(http.StatusInternalServerError, fres.Response.StatusInternalServerError(http.StatusInternalServerError))
	}

	log.Print("Successfully update a equipment")
	return c.JSON(http.StatusOK, fres.Response.StatusOK(equipments))
}

func (ctrl EquipmentsController) DeleteEquipment(c echo.Context) error {
	param := c.Param("id")
	id, _ := strconv.Atoi(param)

	err := ctrl.service.DeleteEquipment(id)

	if err != nil {
		if strings.Contains(err.Error(), "equipment id not found") {
			log.Print("Error on UpdateEquipment request", err.Error())
			return c.JSON(http.StatusNotFound, fres.Response.StatusNotFound(err.Error()))
		}

		log.Print("Error on DeleteEquipment server:", err.Error())
		return c.JSON(http.StatusInternalServerError, fres.Response.StatusInternalServerError(http.StatusInternalServerError))
	}

	log.Print("Successfully delete equipment")
	return c.JSON(http.StatusOK, fres.SuccessResponse{
		Success: true,
		Message: "Successfully delete equipment",
		Data:    id,
	})
}
