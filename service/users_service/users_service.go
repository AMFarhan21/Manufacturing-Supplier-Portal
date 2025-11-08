package users_service

import (
	"Manufacturing-Supplier-Portal/model"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type (
	UsersService struct {
		repo      UsersRepo
		jwtSecret string
	}
	Service interface {
		RegisterUser(data model.Users) (model.Users, error)
		Login(email, password string) (string, error)
		FindUserById(id string) (model.UsersResponse, error)
		GetAll() ([]model.UsersResponse, error)
		TopUp(userId string, amount float64) (float64, error)
	}
)

func NewUsersService(repo UsersRepo, secret string) Service {
	return &UsersService{
		repo:      repo,
		jwtSecret: secret,
	}
}

func (s UsersService) RegisterUser(data model.Users) (model.Users, error) {
	user, _ := s.repo.FindByEmail(data.Email)
	if user.Email != "" {
		return model.Users{}, errors.New("email already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return model.Users{}, err
	}

	data.Id = uuid.NewString()
	data.Password = string(hashedPassword)

	return s.repo.Register(data)
}

func (s UsersService) Login(email, password string) (string, error) {
	user, err := s.repo.FindByEmail(email)
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
	user, err := s.repo.FindById(id)
	if err != nil {
		return model.UsersResponse{}, err
	}

	return user, nil
}

func (s UsersService) GetAll() ([]model.UsersResponse, error) {
	return s.repo.GetAll()
}

func (s UsersService) TopUp(userId string, amount float64) (float64, error) {
	user, err := s.repo.FindById(userId)
	if err != nil {
		return 0, err
	}

	amount = user.DepositAmount + amount

	return s.repo.UpdateDepositAmount(userId, amount)
}
