package controller

import (
	"Manufacturing-Supplier-Portal/model"
	"Manufacturing-Supplier-Portal/service/users_service"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/AMFarhan21/fres"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type (
	UsersController struct {
		service  users_service.Service
		validate *validator.Validate
	}

	RegisterInput struct {
		Username      string  `json:"username" validate:"required,min=5"`
		Email         string  `json:"email" validate:"required,email"`
		Password      string  `json:"password" validate:"required,min=5"`
		DepositAmount float64 `json:"deposit_amount"`
		Role          string  `json:"role" validate:"oneof=user admin"`
	}

	LoginInput struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=5"`
	}

	TopUpInput struct {
		DepositAmount float64 `json:"deposit_amount" validate:"required,min=10000"`
	}
)

func NewUsersController(service users_service.Service) *UsersController {
	return &UsersController{
		service:  service,
		validate: validator.New(),
	}
}

// ShowAccount godoc
// @Summary      Register user
// @Description  Register a new user account
// @Tags         Authentication
// @Accept       json
// @Produce      json
// @Param		request body RegisterInput true "Registration details"
// @Success      201  {object}  users_service.UsersResponse "User registered successfully"
// @Failure      400  {object}  map[string]interface{} "Invalid input or validation error"
// @Router       /auth/register [post]
func (ctrl UsersController) RegisterUser(c echo.Context) error {
	var request RegisterInput

	if request.Role == "" {
		request.Role = "user"
	}

	if err := c.Bind(&request); err != nil {
		log.Print("Error on RegisterUser request body:", err.Error())
		return c.JSON(http.StatusBadRequest, fres.Response.StatusBadRequest(err.Error()))
	}

	if err := ctrl.validate.Struct(request); err != nil {
		log.Print("Error on RegisterUser validation:", err.Error())
		return c.JSON(http.StatusBadRequest, fres.Response.StatusBadRequest(err.Error()))
	}

	user, err := ctrl.service.RegisterUser(model.Users{
		Username:      request.Username,
		Email:         request.Email,
		Password:      request.Password,
		DepositAmount: request.DepositAmount,
		Role:          request.Role,
	})
	if err != nil {
		if strings.Contains(err.Error(), "email already exists") {
			log.Print("Error on RegisterUser service:", err.Error())
			return c.JSON(http.StatusConflict, fres.Response.StatusConflict(err.Error()))
		}

		log.Print("Error on RegisterUser service:", err.Error())
		return c.JSON(http.StatusInternalServerError, fres.Response.StatusInternalServerError(http.StatusInternalServerError))
	}

	log.Print("Successfully register a user")
	return c.JSON(http.StatusCreated, fres.Response.StatusCreated(model.UsersResponse{
		Id:            user.Id,
		Username:      user.Username,
		Email:         user.Email,
		DepositAmount: user.DepositAmount,
		Role:          user.Role,
	}))

}

// @Summary      Login
// @Description  Authenticate user and return JWT token
// @Tags         Authentication
// @Accept       json
// @Produce      json
// @Param        request body LoginInput true "Login credentials"
// @Success      200    {object}  map[string]interface{}
// @Failure      400    {object}  map[string]interface{}
// @Failure      401    {object}  map[string]interface{}
// @Router       /auth/login [post]
func (ctrl UsersController) LoginUser(c echo.Context) error {
	var request LoginInput

	if err := c.Bind(&request); err != nil {
		log.Print("Error on LoginUser request body:", err.Error())
		return c.JSON(http.StatusBadRequest, fres.Response.StatusBadRequest(err.Error()))
	}

	if err := ctrl.validate.Struct(request); err != nil {
		log.Print("Error on LoginUser validation:", err.Error())
		return c.JSON(http.StatusBadRequest, fres.Response.StatusBadRequest(err.Error()))
	}

	signedToken, err := ctrl.service.Login(request.Email, request.Password)
	if err != nil {
		if strings.Contains(err.Error(), "invalid email or password") {
			log.Print("Error on LoginUser service:", err.Error())
			return c.JSON(http.StatusUnauthorized, fres.Response.StatusUnauthorized(err.Error()))
		}

		log.Print("Error on LoginUser service:", err.Error())
		return c.JSON(http.StatusInternalServerError, fres.Response.StatusInternalServerError(http.StatusInternalServerError))
	}

	log.Print("Successfully login")
	return c.JSON(http.StatusOK, fres.Response.StatusOK(signedToken))

}

// @Summary      Get user by jwttoken id
// @Description  User can see the specific information in their user account
// @Tags         Users
// @Accept       json
// @Produce      json
// @Security	BearerAuth
// @Success      200    {object}  users_service.Users "User retrieved successfully"
// @Failure      400    {object}  map[string]interface{} "Invalid userId"
// @Failure      401    {object}  map[string]interface{} "Unauthorized"
// @Failure 	404 	{object}  map[string]interface{} "Forbidden"
// @Failure      500    {object}  map[string]interface{} "User not found"
// @Router       /users/me [get]
func (ctrl UsersController) GetUserLogin(c echo.Context) error {
	id := c.Get("id").(string)

	user, err := ctrl.service.FindUserById(id)
	if err != nil {
		log.Print("Error on GetUserLogin:", err.Error())
		return c.JSON(http.StatusInternalServerError, fres.Response.StatusInternalServerError(http.StatusInternalServerError))
	}

	log.Print("Successfully get login user")
	return c.JSON(http.StatusOK, fres.Response.StatusOK(model.UsersResponse{
		Id:            user.Id,
		Username:      user.Username,
		Email:         user.Email,
		DepositAmount: user.DepositAmount,
		Role:          user.Role,
	}))
}

// @Summary      Get all users
// @Description  Get list of users
// @Tags         Admin - Users
// @Accept       json
// @Produce      json
// @Security	BearerAuth
// @Success      200    {object}  users_service.Users "Successfully retrieved all users"
// @Failure      401    {object}  map[string]interface{} "Unauthorized"
// @Failure      403    {object}  map[string]interface{} "Forbidden"
// @Failure      500    {object}  map[string]interface{} "Users not found"
// @Router       /users/list [get]
func (ctrl UsersController) GetAllUsers(c echo.Context) error {
	users, err := ctrl.service.GetAll()
	if err != nil {
		log.Print("Error on GetUserLogin:", err.Error())
		return c.JSON(http.StatusInternalServerError, fres.Response.StatusInternalServerError(http.StatusInternalServerError))
	}

	log.Print("Successfully get login user")
	return c.JSON(http.StatusOK, fres.Response.StatusOK(users))
}

// @Summary      Top Up
// @Description  Users can top up their deposit amount
// @Tags         Users
// @Accept       json
// @Produce      json
// @Security	BearerAuth
// @Success      200    {object}  map[string]interface{} "Successfully top up your deposit"
// @Param		request body TopUpInput true "Input deposit amount"
// @Failure      400    {object}  map[string]interface{} "Invalid request or validation input"
// @Failure      401    {object}  map[string]interface{} "Unauthorized"
// @Failure      500    {object}  map[string]interface{} "Top up failed"
// @Router       /users/topup [post]
func (ctrl UsersController) TopUpDeposit(c echo.Context) error {
	id := c.Get("id").(string)

	var request TopUpInput

	if err := c.Bind(&request); err != nil {
		log.Print("Error on TopUpDeposit request body:", err.Error())
		return c.JSON(http.StatusBadRequest, fres.Response.StatusBadRequest(err.Error()))
	}

	if err := ctrl.validate.Struct(request); err != nil {
		log.Print("Error on TopUpDeposit validation:", err.Error())
		return c.JSON(http.StatusBadRequest, fres.Response.StatusBadRequest(err.Error()))
	}

	InvoiceURL, err := ctrl.service.TopUp(id, request.DepositAmount)
	if err != nil {
		if strings.Contains(err.Error(), "cannot find user with the id") {
			log.Print("Error on TopUpDeposit service:", err.Error())
			return c.JSON(http.StatusNotFound, fres.Response.StatusNotFound(err.Error()))
		}

		log.Print("Error on TopUpDeposit service:", err.Error())
		return c.JSON(http.StatusInternalServerError, fres.Response.StatusInternalServerError(http.StatusInternalServerError))
	}

	log.Print("Successfully top up your deposit")
	return c.JSON(http.StatusOK, fres.Response.StatusOK(fmt.Sprintf("Successfully top up. Your current deposit is Rp.%.2f", InvoiceURL)))
}
