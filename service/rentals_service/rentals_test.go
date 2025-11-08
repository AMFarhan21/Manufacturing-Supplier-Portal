package rentals_service_test

// import (
// 	"Manufacturing-Supplier-Portal/model"
// 	mock_equipments_service "Manufacturing-Supplier-Portal/service/equipments_service/mock"
// 	mock_payments_service "Manufacturing-Supplier-Portal/service/payments_service/mock"
// 	mock_rental_histories_service "Manufacturing-Supplier-Portal/service/rental_histories_service/mock"
// 	"Manufacturing-Supplier-Portal/service/rentals_service"
// 	mock_rentals_service "Manufacturing-Supplier-Portal/service/rentals_service/mock"
// 	mock_users_service "Manufacturing-Supplier-Portal/service/users_service/mock"
// 	mock_xendit_service "Manufacturing-Supplier-Portal/service/xendit_service/mock"
// 	"errors"
// 	"testing"
// 	"time"

// 	"github.com/golang/mock/gomock"
// 	"github.com/stretchr/testify/assert"
// )

// func TestCreateRental(t *testing.T) {
// 	available := true
// 	now := time.Now()
// 	tests := []struct {
// 		name                string
// 		inputRental         model.Rentals
// 		mockRentals         func(m *mock_rentals_service.MockRentalsRepo)
// 		mockEquipments      func(m *mock_equipments_service.MockEquipmentsRepo)
// 		mockUsers           func(m *mock_users_service.MockUsersRepo)
// 		mockRentalHistories func(m *mock_rental_histories_service.MockRentalHistoriesRepo)
// 		mockXendit          func(m *mock_xendit_service.MockXenditRepo)
// 		mockPayments        func(m *mock_payments_service.MockPaymentsRepo)
// 		wantErr             bool
// 	}{
// 		{
// 			name:        "Error on get by id",
// 			inputRental: model.Rentals{EquipmentId: 1},
// 			mockRentals: func(m *mock_rentals_service.MockRentalsRepo) {},
// 			mockEquipments: func(m *mock_equipments_service.MockEquipmentsRepo) {
// 				m.EXPECT().GetById(1).Return(model.Equipments{}, errors.New("record not found"))
// 			},
// 			mockUsers:           func(m *mock_users_service.MockUsersRepo) {},
// 			mockRentalHistories: func(m *mock_rental_histories_service.MockRentalHistoriesRepo) {},
// 			mockXendit:          func(m *mock_xendit_service.MockXenditRepo) {},
// 			mockPayments:        func(m *mock_payments_service.MockPaymentsRepo) {},
// 			wantErr:             true,
// 		},
// 		{
// 			name: "Error on find user by id",
// 			inputRental: model.Rentals{
// 				EquipmentId: 1,
// 				UserId:      "410dkjfboi"},
// 			mockRentals: func(m *mock_rentals_service.MockRentalsRepo) {},
// 			mockEquipments: func(m *mock_equipments_service.MockEquipmentsRepo) {
// 				m.EXPECT().GetById(1).Return(model.Equipments{
// 					Id:            1,
// 					Name:          "Farhan",
// 					CategoryId:    2,
// 					Description:   "farhan",
// 					PricePerDay:   2000,
// 					PricePerWeek:  2000,
// 					PricePerMonth: 2000,
// 					PricePerYear:  2000,
// 					Available:     &available,
// 				}, nil)
// 			},
// 			mockUsers: func(m *mock_users_service.MockUsersRepo) {
// 				m.EXPECT().FindById("410dkjfboi").Return(model.UsersResponse{}, errors.New("record not found"))
// 			},
// 			mockRentalHistories: func(m *mock_rental_histories_service.MockRentalHistoriesRepo) {},
// 			mockXendit:          func(m *mock_xendit_service.MockXenditRepo) {},
// 			mockPayments:        func(m *mock_payments_service.MockPaymentsRepo) {},
// 			wantErr:             true,
// 		},
// 		{
// 			name: "Error on find user by id",
// 			inputRental: model.Rentals{
// 				EquipmentId: 1,
// 				UserId:      "410dkjfboi"},
// 			mockRentals: func(m *mock_rentals_service.MockRentalsRepo) {
// 				m.EXPECT().Create(gomock.Any()).Return(model.Rentals{}, errors.New("There is no data"))
// 			},
// 			mockEquipments: func(m *mock_equipments_service.MockEquipmentsRepo) {
// 				m.EXPECT().GetById(1).Return(model.Equipments{
// 					Id:            1,
// 					Name:          "Farhan",
// 					CategoryId:    2,
// 					Description:   "farhan",
// 					PricePerDay:   2000,
// 					PricePerWeek:  2000,
// 					PricePerMonth: 2000,
// 					PricePerYear:  2000,
// 					Available:     &available,
// 				}, nil)
// 			},
// 			mockUsers: func(m *mock_users_service.MockUsersRepo) {
// 				m.EXPECT().FindById("410dkjfboi").Return(model.UsersResponse{
// 					Id:            "hello world",
// 					Username:      "hello world",
// 					Email:         "hello world",
// 					DepositAmount: 2000000,
// 					Role:          "hello world",
// 				}, nil)
// 			},
// 			mockRentalHistories: func(m *mock_rental_histories_service.MockRentalHistoriesRepo) {},
// 			mockXendit:          func(m *mock_xendit_service.MockXenditRepo) {},
// 			mockPayments:        func(m *mock_payments_service.MockPaymentsRepo) {},
// 			wantErr:             true,
// 		},
// 		{
// 			name: "Error on find user by id",
// 			inputRental: model.Rentals{
// 				EquipmentId: 1,
// 				UserId:      "410dkjfboi"},
// 			mockRentals: func(m *mock_rentals_service.MockRentalsRepo) {
// 				m.EXPECT().Create(gomock.Any()).Return(model.Rentals{
// 					Id:           1,
// 					UserId:       "213cgzjhk",
// 					EquipmentId:  1,
// 					RentalPeriod: "dddd",
// 					StartDate:    &now,
// 					EndDate:      &now,
// 					Price:        12548,
// 					Status:       "hehe",
// 					CreatedAt:    time.Now(),
// 				}, nil)
// 			},
// 			mockEquipments: func(m *mock_equipments_service.MockEquipmentsRepo) {
// 				m.EXPECT().GetById(1).Return(model.Equipments{
// 					Id:            1,
// 					Name:          "Farhan",
// 					CategoryId:    2,
// 					Description:   "farhan",
// 					PricePerDay:   2000,
// 					PricePerWeek:  2000,
// 					PricePerMonth: 2000,
// 					PricePerYear:  2000,
// 					Available:     &available,
// 				}, nil)
// 			},
// 			mockUsers: func(m *mock_users_service.MockUsersRepo) {
// 				m.EXPECT().FindById("410dkjfboi").Return(model.UsersResponse{
// 					Id:            "hello world",
// 					Username:      "hello world",
// 					Email:         "hello world",
// 					DepositAmount: 2000000,
// 					Role:          "hello world",
// 				}, nil)
// 			},
// 			mockRentalHistories: func(m *mock_rental_histories_service.MockRentalHistoriesRepo) {
// 				m.EXPECT().CreateRentalHistory(gomock.Any()).Return(model.RentalHistories{}, errors.New("There is no data"))
// 			},
// 			mockXendit:   func(m *mock_xendit_service.MockXenditRepo) {},
// 			mockPayments: func(m *mock_payments_service.MockPaymentsRepo) {},
// 			wantErr:      true,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			ctrl := gomock.NewController(t)
// 			defer ctrl.Finish()
// 			mock_rentalsRepo := mock_rentals_service.NewMockRentalsRepo(ctrl)
// 			mock_equipmentsRepo := mock_equipments_service.NewMockEquipmentsRepo(ctrl)
// 			mock_usersRepo := mock_users_service.NewMockUsersRepo(ctrl)
// 			mock_rentalHistoriesRepo := mock_rental_histories_service.NewMockRentalHistoriesRepo(ctrl)
// 			mock_xenditRepo := mock_xendit_service.NewMockXenditRepo(ctrl)
// 			mock_paymentsRepo := mock_payments_service.NewMockPaymentsRepo(ctrl)

// 			tt.mockRentals(mock_rentalsRepo)
// 			tt.mockEquipments(mock_equipmentsRepo)
// 			tt.mockUsers(mock_usersRepo)
// 			tt.mockRentalHistories(mock_rentalHistoriesRepo)
// 			tt.mockXendit(mock_xenditRepo)
// 			tt.mockPayments(mock_paymentsRepo)

// 			rentalsService := rentals_service.NewRentalsService(
// 				mock_rentalsRepo,
// 				mock_equipmentsRepo,
// 				mock_xenditRepo,
// 				mock_paymentsRepo,
// 				mock_rentalHistoriesRepo,
// 				mock_usersRepo,
// 			)

// 			rentalWithInvoice, err := rentalsService.CreateRental(tt.inputRental)
// 			if tt.wantErr {
// 				assert.Equal(t, model.RentalsWithInvoiceUrl{}, rentalWithInvoice)
// 				assert.NotNil(t, err)
// 			} else {
// 				assert.Nil(t, err)
// 			}

// 		})
// 	}
// }
