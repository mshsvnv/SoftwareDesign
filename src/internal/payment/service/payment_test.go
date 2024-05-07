package service

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	"src/internal/payment/dto"
	"src/internal/payment/model"
	"src/internal/payment/repository/mocks"
)

type PaymentServiceTestSuite struct {
	suite.Suite

	mockRepo *mocks.IPaymentRepository
	service  IPaymentService
}

func (suite *PaymentServiceTestSuite) SetupTest() {
	suite.mockRepo = mocks.NewIPaymentRepository(suite.T())
	suite.service = NewPaymentService(suite.mockRepo)
}

func TestPaymentServiceTestSuite(t *testing.T) {
	suite.Run(t, new(PaymentServiceTestSuite))
}

// GetPaymentByOrderID
func (suite *PaymentServiceTestSuite) TestGetPaymentByOrderIDFail() {

	orderID := "123"

	suite.mockRepo.On("GetPaymentByOrderID", mock.Anything, orderID).
		Return(nil, errors.New("error")).Times(1)

	suite.mockRepo.On("Create", mock.Anything, mock.Anything).
		Return(errors.New("error")).Times(1)

	payment, err := suite.service.GetPaymentByOrderID(context.Background(), orderID)

	suite.Nil(payment)
	suite.NotNil(err)
}

func (suite *PaymentServiceTestSuite) TestGetPaymentByOrderIDNewPayment() {

	orderID := "123"

	suite.mockRepo.On("GetPaymentByOrderID", mock.Anything, orderID).
		Return(nil, errors.New("error")).Times(1)

	suite.mockRepo.On("Create", mock.Anything, mock.Anything).
		Return(nil).Times(1)

	payment, err := suite.service.GetPaymentByOrderID(context.Background(), orderID)

	suite.Nil(err)
	suite.NotNil(payment)
}

func (suite *PaymentServiceTestSuite) TestGetPaymentByIDSuccess() {

	orderID := "123"

	suite.mockRepo.On("GetPaymentByOrderID", mock.Anything, orderID).
		Return(&model.Payment{
			OrderID: orderID,
			Status:  model.PaymentStatusInProgress,
		}, nil).Times(1)

	payment, err := suite.service.GetPaymentByOrderID(context.Background(), orderID)

	suite.NotNil(payment)
	suite.Nil(err)
}

// GetMyPayments
func (suite *PaymentServiceTestSuite) TestGetMyPaymentsFail() {

	req := &dto.ListPaymentReq{
		UserID: "123",
		Status: string(model.PaymentStatusNew),
	}

	suite.mockRepo.On("GetMyPayments", mock.Anything, req).
		Return(nil, errors.New("error")).Times(1)

	payment, err := suite.service.GetMyPayments(context.Background(), req)

	suite.Nil(payment)
	suite.NotNil(err)
}

func (suite *PaymentServiceTestSuite) TestGetMyPaymentsSucces() {

	orderID1 := "123"
	orderID2 := "124"

	req := &dto.ListPaymentReq{
		UserID: "123",
		Status: string(model.PaymentStatusNew),
	}

	suite.mockRepo.On("GetMyPayments", mock.Anything, req).
		Return([]*model.Payment{
			{
				OrderID: orderID1,
				Status:  model.PaymentStatusNew,
			},
			{
				OrderID: orderID2,
				Status:  model.PaymentStatusNew,
			},
		}, nil).Times(1)

	payment, err := suite.service.GetMyPayments(context.Background(), req)

	suite.NotNil(payment)
	suite.Nil(err)
}
