package users_service

import (
	"Manufacturing-Supplier-Portal/model"
	"Manufacturing-Supplier-Portal/service/mailjet_service"
	"Manufacturing-Supplier-Portal/service/xendit_service"
	"errors"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type (
	UsersService struct {
		usersRepo   UsersRepo
		xenditRepo  xendit_service.XenditRepo
		mailjetRepo mailjet_service.MailjetRepo
		jwtSecret   string
	}
	Service interface {
		RegisterUser(data model.Users) (string, error)
		Login(email, password string) (string, error)
		FindUserById(id string) (model.UsersResponse, error)
		GetAll() ([]model.UsersResponse, error)
		TopUp(userId string, amount float64) (float64, error)
		GetTopUpInvoiceURL(userId string, amount float64) (string, error)
		VerifiedEmail(token string) (model.Users, error)
	}

	MapClaims struct {
		Data model.Users `json:"data"`
		*jwt.RegisteredClaims
	}
)

func NewUsersService(usersRepo UsersRepo, xenditRepo xendit_service.XenditRepo, mailjetRepo mailjet_service.MailjetRepo, secret string) Service {
	return &UsersService{
		usersRepo:   usersRepo,
		xenditRepo:  xenditRepo,
		mailjetRepo: mailjetRepo,
		jwtSecret:   secret,
	}
}

func (s UsersService) RegisterUser(data model.Users) (string, error) {
	user, _ := s.usersRepo.FindByEmail(data.Email)
	if user.Email != "" {
		return "", errors.New("email already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	data.Id = uuid.NewString()
	data.Password = string(hashedPassword)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, MapClaims{
		Data: data,
		RegisteredClaims: &jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 5)),
		},
	})

	signedToken, err := token.SignedString([]byte(s.jwtSecret))
	if err != nil {
		return "", err
	}

	err = s.mailjetRepo.SendMailjetMessage("andifarhanhakzah@gmail.com", "farhan", data.Email, data.Username, signedToken)
	if err != nil {
		return "", err
	}

	return "Check your email and validate", nil
}

func (s UsersService) VerifiedEmail(token string) (model.Users, error) {
	claims := &MapClaims{}

	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.jwtSecret), nil
	})
	if err != nil {
		return model.Users{}, err
	}

	expAt, _ := claims.GetExpirationTime()
	if time.Now().After(expAt.Time) {
		return model.Users{}, errors.New("expired url")
	}

	log.Print(claims.Data)

	return s.usersRepo.Register(claims.Data)
}

func (s UsersService) Login(email, password string) (string, error) {
	user, err := s.usersRepo.FindByEmail(email)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	type MapClaims struct {
		Id   string `json:"id"`
		Role string `json:"role"`
		*jwt.RegisteredClaims
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, MapClaims{
		Id:   user.Id,
		Role: user.Role,
		RegisteredClaims: &jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	})

	signedToken, err := token.SignedString([]byte(s.jwtSecret))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (s UsersService) FindUserById(id string) (model.UsersResponse, error) {
	user, err := s.usersRepo.FindById(id)
	if err != nil {
		return model.UsersResponse{}, err
	}

	return user, nil
}

func (s UsersService) GetAll() ([]model.UsersResponse, error) {
	return s.usersRepo.GetAll()
}

func (s UsersService) GetTopUpInvoiceURL(userId string, amount float64) (string, error) {
	user, err := s.usersRepo.FindById(userId)
	if err != nil {
		return "", err
	}

	invoiceURL, err := s.xenditRepo.XenditInvoiceUrl(userId, "TOPUP", user.Username, user.Email, "TOPUP", "TOPUP", 9999999, amount)
	if err != nil {
		return "", err
	}

	if invoiceURL == "" {
		return "", errors.New("invoice URL is empty")
	}

	return invoiceURL, nil
}

func (s UsersService) TopUp(userId string, amount float64) (float64, error) {
	user, err := s.usersRepo.FindById(userId)
	if err != nil {
		return 0, err
	}

	amount = user.DepositAmount + amount

	return s.usersRepo.UpdateDepositAmount(userId, amount)
}
