package service

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	"src/internal/order/dto"
	"src/internal/order/model"
	"src/internal/order/repository/mocks"
)

type OrderServiceTestSuite struct {
	suite.Suite

	mockRepo        *mocks.IOrderRepository
	mockRepoRacket  *mocks.IOrderRacketRepository
	mockRepoPayment *mocks.IPaymentRepository

	service IOrderService
}

func (suite *OrderServiceTestSuite) SetupTest() {
	suite.mockRepo = mocks.NewIOrderRepository(suite.T())
	suite.mockRepoRacket = mocks.NewIOrderRacketRepository(suite.T())
	suite.mockRepoPayment = mocks.NewIPaymentRepository(suite.T())

	suite.service = NewOrderService(
		suite.mockRepo,
		suite.mockRepoRacket,
		suite.mockRepoPayment,
	)
}

func TestOrderServiceTestSuite(t *testing.T) {
	suite.Run(t, new(OrderServiceTestSuite))
}

// GetOrderByID
func (suite *OrderServiceTestSuite) TestGetOrderByIDFail() {

	orderID := "123"

	suite.mockRepo.On("GetOrderByID", mock.Anything, orderID, true).
		Return(nil, errors.New("error")).Times(1)

	order, err := suite.service.GetOrderByID(context.Background(), orderID)

	suite.Nil(order)
	suite.NotNil(err)
}

func (suite *OrderServiceTestSuite) TestGetOrderByIDSuccess() {

	orderID := "123"

	suite.mockRepo.On("GetOrderByID", mock.Anything, orderID, true).
		Return(
			&model.Order{
				UserID: "1",
				Status: model.OrderStatusNew,
			}, nil).Times(1)

	order, err := suite.service.GetOrderByID(context.Background(), orderID)

	suite.NotNil(order)
	suite.Equal("1", order.UserID)
	suite.Equal(model.OrderStatusNew, order.Status)
	suite.Nil(err)
}

// GetMyOrders
func (suite *OrderServiceTestSuite) TestListOrdersFail() {

	req := &dto.ListOrderReq{
		Status: "new",
	}

	suite.mockRepo.On("GetMyOrders", mock.Anything, req).
		Return(nil, errors.New("error")).Times(1)

	orders, err := suite.service.GetMyOrders(context.Background(), req)

	suite.Nil(orders)
	suite.NotNil(err)
}

func (suite *OrderServiceTestSuite) TestListOrdersSuccess() {
	req := &dto.ListOrderReq{
		Status: "new",
	}

	suite.mockRepo.On("GetMyOrders", mock.Anything, req).
		Return(
			[]*model.Order{
				{
					UserID: "userID",
					Status: model.OrderStatusNew,
				},
			}, nil).Times(1)

	orders, err := suite.service.GetMyOrders(context.Background(), req)

	suite.NotNil(orders)
	suite.Equal(1, len(orders))
	suite.Equal("userID", orders[0].UserID)
	suite.Equal(model.OrderStatusNew, orders[0].Status)
	suite.Nil(err)
}

// PlaceOrder
func (suite *OrderServiceTestSuite) TestPlaceOrderGetRacketByIDFail() {

	req := &dto.PlaceOrderReq{
		UserID: "userID",
		Rackets: []dto.PlaceOrderRacketReq{
			{
				RacketID: "productID",
				Quantity: 2,
			},
		},
	}

	suite.mockRepoRacket.On("GetRacketByID", mock.Anything, "productID").
		Return(nil, errors.New("error")).Times(1)

	order, err := suite.service.PlaceOrder(context.Background(), req)

	suite.Nil(order)
	suite.NotNil(err)
}

func (suite *OrderServiceTestSuite) TestPlaceOrderCreateOrderFail() {

	req := &dto.PlaceOrderReq{
		UserID: "userID",
		Rackets: []dto.PlaceOrderRacketReq{
			{
				RacketID: "productID",
				Quantity: 2,
			},
		},
	}

	suite.mockRepoRacket.On("GetRacketByID", mock.Anything, "productID").
		Return(&model.Racket{
			Price: 1.1,
		}, nil).Times(1)

	suite.mockRepo.On("Create", mock.Anything, "userID", mock.Anything).
		Return(nil, errors.New("error")).Times(1)

	order, err := suite.service.PlaceOrder(context.Background(), req)

	suite.Nil(order)
	suite.NotNil(err)
}

func (suite *OrderServiceTestSuite) TestPlaceOrderCreatePaymentFail() {

	req := &dto.PlaceOrderReq{
		UserID: "userID",
		Rackets: []dto.PlaceOrderRacketReq{
			{
				RacketID: "12",
				Quantity: 2,
			},
		},
	}

	suite.mockRepoRacket.On("GetRacketByID", mock.Anything, "12").
		Return(&model.Racket{
			Price: 1.1,
		}, nil).Times(1)

	suite.mockRepo.On("Create", mock.Anything, "userID", mock.Anything).
		Return(&model.Order{
			ID:     "orderID",
			UserID: "userID",
			Rackets: []*model.OrderRacket{
				{
					RacketID: "12",
					Quantity: 2,
				},
			},
		}, nil).Times(1)

	payment := &model.Payment{
		OrderID:    "orderID",
		Status:     model.PaymentStatusInProgress,
		TotalPrice: 2.2,
	}

	suite.mockRepoPayment.On("Create", mock.Anything, payment).
		Return(errors.New("error")).Times(1)

	order, err := suite.service.PlaceOrder(context.Background(), req)

	suite.Nil(order)
	suite.NotNil(err)
}

func (suite *OrderServiceTestSuite) TestPlaceOrderSuccess() {

	req := &dto.PlaceOrderReq{
		UserID: "userID",
		Rackets: []dto.PlaceOrderRacketReq{
			{
				RacketID: "12",
				Quantity: 2,
			},
		},
	}

	suite.mockRepoRacket.On("GetRacketByID", mock.Anything, "12").
		Return(&model.Racket{
			Price: 10.1,
		}, nil).Times(1)

	suite.mockRepo.On("Create", mock.Anything, "userID", mock.Anything).
		Return(&model.Order{
			ID:     "orderID",
			UserID: "userID",
			Rackets: []*model.OrderRacket{
				{
					RacketID: "12",
					Quantity: 2,
				},
			},
		}, nil).Times(1)

	payment := &model.Payment{
		OrderID:    "orderID",
		Status:     model.PaymentStatusInProgress,
		TotalPrice: 20.2,
	}

	suite.mockRepoPayment.On("Create", mock.Anything, payment).
		Return(nil).Times(1)

	order, err := suite.service.PlaceOrder(context.Background(), req)

	suite.NotNil(order)
	suite.Equal(req.UserID, order.UserID)
	suite.Equal(1, len(order.Rackets))
	suite.Equal(req.Rackets[0].RacketID, order.Rackets[0].RacketID)
	suite.Equal(req.Rackets[0].Quantity, order.Rackets[0].Quantity)
	suite.Nil(err)
}

// Cancel Order
func (suite *OrderServiceTestSuite) TestCancelOrderFail() {

	userID := "userID"
	orderID := "orderID"

	suite.mockRepo.On("GetOrderByID", mock.Anything, orderID, false).
		Return(&model.Order{
			UserID: userID,
			Status: model.OrderStatusNew,
		}, nil).Times(1)

	suite.mockRepo.On("Update", mock.Anything, &model.Order{
		UserID: userID,
		Status: model.OrderStatusCancelled,
	}).Return(errors.New("error")).Times(1)

	order, err := suite.service.CancelOrder(context.Background(), orderID, userID)
	suite.Nil(order)
	suite.NotNil(err)
}

func (suite *OrderServiceTestSuite) TestCancelOrderDifferenceUserId() {
	userID := "userID"
	orderID := "orderID"

	suite.mockRepo.On("GetOrderByID", mock.Anything, orderID, false).
		Return(&model.Order{
			UserID: "userID1",
			Status: model.OrderStatusNew,
		}, nil).Times(1)

	order, err := suite.service.CancelOrder(context.Background(), orderID, userID)
	suite.Nil(order)
	suite.NotNil(err)
}

func (suite *OrderServiceTestSuite) TestCancelOrderInvalidStatus() {
	userID := "userID"
	orderID := "orderID"

	suite.mockRepo.On("GetOrderByID", mock.Anything, orderID, false).
		Return(&model.Order{
			UserID: userID,
			Status: model.OrderStatusCancelled,
		}, nil).Times(1)

	order, err := suite.service.CancelOrder(context.Background(), orderID, userID)
	suite.Nil(order)
	suite.NotNil(err)
}

func (suite *OrderServiceTestSuite) TestCancelOrderGetOrderByIDFail() {
	userID := "userID"
	orderID := "orderID"

	suite.mockRepo.On("GetOrderByID", mock.Anything, orderID, false).
		Return(nil, errors.New("error")).Times(1)

	order, err := suite.service.CancelOrder(context.Background(), orderID, userID)
	suite.Nil(order)
	suite.NotNil(err)
}

func (suite *OrderServiceTestSuite) TestCancelOrderSuccess() {
	userID := "userID"
	orderID := "orderID"

	suite.mockRepo.On("GetOrderByID", mock.Anything, orderID, false).
		Return(&model.Order{
			ID:     orderID,
			UserID: userID,
			Status: model.OrderStatusNew,
		}, nil).Times(1)

	suite.mockRepo.On("Update", mock.Anything, &model.Order{
		ID:     orderID,
		UserID: userID,
		Status: model.OrderStatusCancelled,
	}).Return(nil).Times(1)

	suite.mockRepoPayment.On("GetPaymentByID", mock.Anything, orderID).
		Return(&model.Payment{
			OrderID: orderID,
			Status:  model.PaymentStatusInProgress,
		}, nil).Times(1)

	suite.mockRepoPayment.On("Update", mock.Anything, &model.Payment{
		OrderID: orderID,
		Status:  model.PaymentStatusCancelled,
	}).Return(nil).Times(1)

	order, err := suite.service.CancelOrder(context.Background(), orderID, userID)
	suite.NotNil(order)
	suite.Equal(userID, order.UserID)
	suite.Equal(model.OrderStatusCancelled, order.Status)
	suite.Nil(err)
}
