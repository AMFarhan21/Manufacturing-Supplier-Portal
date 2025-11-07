package payments_service_test

import (
	"Manufacturing-Supplier-Portal/service/payments_service"
	mock_payments_service "Manufacturing-Supplier-Portal/service/payments_service/mock"
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	tests := []struct {
		name          string
		inputPayments payments_service.Payments
		mockPayments  func(m *mock_payments_service.MockPaymentsRepo)
		wantErr       bool
	}{
		{
			name:          "Error on create",
			inputPayments: payments_service.Payments{},
			mockPayments: func(m *mock_payments_service.MockPaymentsRepo) {
				m.EXPECT().Create(gomock.Any()).Return(payments_service.Payments{}, errors.New("There is no data"))
			},
			wantErr: true,
		},
		{
			name:          "Success on create",
			inputPayments: payments_service.Payments{},
			mockPayments: func(m *mock_payments_service.MockPaymentsRepo) {
				m.EXPECT().Create(gomock.Any()).Return(payments_service.Payments{
					Id:            1,
					UserId:        "kdjshfsakjdsgh",
					RentalId:      1,
					Amount:        199959595,
					PaymentMethod: "BANK",
					Status:        "PAID",
					CreatedAt:     time.Now(),
				}, nil)
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mock_paymentsRepo := mock_payments_service.NewMockPaymentsRepo(ctrl)

			tt.mockPayments(mock_paymentsRepo)

			paymentsService := payments_service.NewPaymentsService(
				mock_paymentsRepo,
			)

			create, err := paymentsService.Create(tt.inputPayments)
			if tt.wantErr {
				assert.Equal(t, payments_service.Payments{}, create)
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}

		})
	}
}

func TestGetById(t *testing.T) {
	tests := []struct {
		id           int
		userId       string
		name         string
		mockPayments func(m *mock_payments_service.MockPaymentsRepo)
		wantErr      bool
	}{
		{
			id:     1,
			userId: "asdhfdfsadojf",
			name:   "Error on get by id",
			mockPayments: func(m *mock_payments_service.MockPaymentsRepo) {
				m.EXPECT().GetById(1, "asdhfdfsadojf").Return(payments_service.Payments{}, errors.New("There is no data"))
			},
			wantErr: true,
		},
		{
			id:     1,
			userId: "asdhfdfsadojf",
			name:   "Success on get by id",
			mockPayments: func(m *mock_payments_service.MockPaymentsRepo) {
				m.EXPECT().GetById(1, "asdhfdfsadojf").Return(payments_service.Payments{
					Id:            1,
					UserId:        "kdjshfsakjdsgh",
					RentalId:      1,
					Amount:        199959595,
					PaymentMethod: "BANK",
					Status:        "PAID",
					CreatedAt:     time.Now(),
				}, nil)
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mock_paymentsRepo := mock_payments_service.NewMockPaymentsRepo(ctrl)

			tt.mockPayments(mock_paymentsRepo)

			paymentsService := payments_service.NewPaymentsService(
				mock_paymentsRepo,
			)

			getById, err := paymentsService.GetById(tt.id, tt.userId)
			if tt.wantErr {
				assert.Equal(t, payments_service.Payments{}, getById)
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}

		})
	}
}

func TestUpdateStatusAndMethod(t *testing.T) {
	tests := []struct {
		id           int
		status       string
		method       string
		name         string
		mockPayments func(m *mock_payments_service.MockPaymentsRepo)
		wantErr      bool
	}{
		{
			id:     1,
			status: "hello",
			method: "hai",
			name:   "Error on update status and method",
			mockPayments: func(m *mock_payments_service.MockPaymentsRepo) {
				m.EXPECT().UpdateStatusAndMethod(1, "hello", "hai").Return(errors.New("There is no data"))
			},
			wantErr: true,
		},
		{
			id:     1,
			status: "hello",
			method: "hai",
			name:   "Error on update status and method",
			mockPayments: func(m *mock_payments_service.MockPaymentsRepo) {
				m.EXPECT().UpdateStatusAndMethod(1, "hello", "hai").Return(nil)
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mock_paymentsRepo := mock_payments_service.NewMockPaymentsRepo(ctrl)

			tt.mockPayments(mock_paymentsRepo)

			paymentsService := payments_service.NewPaymentsService(
				mock_paymentsRepo,
			)

			err := paymentsService.UpdateStatusAndMethod(tt.id, tt.status, tt.method)
			if tt.wantErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}

		})
	}
}

func TestBookingReport(t *testing.T) {
	var amount float64 = 123456
	var totalBooking int = 1
	tests := []struct {
		id           int
		userId       string
		name         string
		mockPayments func(m *mock_payments_service.MockPaymentsRepo)
		wantErr      bool
	}{
		{
			name: "Error on BookingReport",
			mockPayments: func(m *mock_payments_service.MockPaymentsRepo) {
				m.EXPECT().BookingReport().Return([]payments_service.BookingsReport{}, errors.New("There is no data"))
			},
			wantErr: true,
		},
		{
			name: "Success on BookingReport",
			mockPayments: func(m *mock_payments_service.MockPaymentsRepo) {
				m.EXPECT().BookingReport().Return([]payments_service.BookingsReport{
					{
						Id:           1,
						Name:         "ayam",
						TotalIncome:  &amount,
						TotalBooking: &totalBooking,
					},
				}, nil)
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mock_paymentsRepo := mock_payments_service.NewMockPaymentsRepo(ctrl)

			tt.mockPayments(mock_paymentsRepo)

			paymentsService := payments_service.NewPaymentsService(
				mock_paymentsRepo,
			)

			bookingReport, err := paymentsService.BookingReport()
			if tt.wantErr {
				assert.Equal(t, []payments_service.BookingsReport{}, bookingReport)
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}

		})
	}
}
