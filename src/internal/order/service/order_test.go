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

	mockRepo       *mocks.IOrderRepository
	mockRacketRepo *mocks.IRacketRepository
	service        IOrderService
}

func (suite *OrderServiceTestSuite) SetupTest() {
	suite.mockRepo = mocks.NewIOrderRepository(suite.T())
	suite.mockRacketRepo = mocks.NewIRacketRepository(suite.T())
	suite.service = NewOrderService(suite.mockRepo, suite.mockRacketRepo)
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
				UserID:     "1",
				Status:     model.OrderStatusNew,
				TotalPrice: 248.,
			}, nil).Times(1)

	order, err := suite.service.GetOrderByID(context.Background(), orderID)

	suite.NotNil(order)
	suite.Equal("1", order.UserID)
	suite.Equal(model.OrderStatusNew, order.Status)
	suite.Equal(248., order.TotalPrice)
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
					UserID:     "userID",
					TotalPrice: 111.2,
					Status:     model.OrderStatusNew,
				},
			}, nil).Times(1)

	orders, err := suite.service.GetMyOrders(context.Background(), req)

	suite.NotNil(orders)
	suite.Equal(1, len(orders))
	suite.Equal("userID", orders[0].UserID)
	suite.Equal(111.2, orders[0].TotalPrice)
	suite.Equal(model.OrderStatusNew, orders[0].Status)
	suite.Nil(err)
}

// PlaceOrder
func (suite *OrderServiceTestSuite) TestPlaceOrderGetRacketByIDFail() {

	req := &dto.PlaceOrderReq{
		UserID: "userID",
		Lines: []dto.PlaceOrderLineReq{
			{
				RacketID: "productID",
				Quantity: 2,
			},
		},
	}

	suite.mockRacketRepo.On("GetRacketByID", mock.Anything, "productID").
		Return(nil, errors.New("error")).Times(1)

	order, err := suite.service.PlaceOrder(context.Background(), req)
	suite.Nil(order)
	suite.NotNil(err)
}

func (suite *OrderServiceTestSuite) TestPlaceOrderCreateFail() {
	req := &dto.PlaceOrderReq{
		UserID: "userID",
		Lines: []dto.PlaceOrderLineReq{
			{
				RacketID: "productID",
				Quantity: 2,
			},
		},
	}

	suite.mockRacketRepo.On("GetRacketByID", mock.Anything, "productID").
		Return(&model.Racket{
			Brand: "brand",
			Price: 1.1,
		}, nil).Times(1)

	suite.mockRepo.On("CreateOrder", mock.Anything, "userID", mock.Anything).
		Return(nil, errors.New("error")).Times(1)

	order, err := suite.service.PlaceOrder(context.Background(), req)

	suite.Nil(order)
	suite.NotNil(err)
}

func (suite *OrderServiceTestSuite) TestPlaceOrderSuccess() {

	req := &dto.PlaceOrderReq{
		UserID: "userID",
		Lines: []dto.PlaceOrderLineReq{
			{
				RacketID: "12",
				Quantity: 2,
			},
		},
	}

	suite.mockRacketRepo.On("GetRacketByID", mock.Anything, "12").
		Return(&model.Racket{
			Brand: "brand",
			Price: 10.1,
		}, nil).Times(1)

	suite.mockRepo.On("CreateOrder", mock.Anything, "userID", mock.Anything).
		Return(&model.Order{
			UserID: "userID",
			Lines: []*model.OrderLine{
				{
					RacketID: "12",
					Quantity: 2,
				},
			},
		}, nil).Times(1)

	order, err := suite.service.PlaceOrder(context.Background(), req)

	suite.NotNil(order)
	suite.Equal(req.UserID, order.UserID)
	suite.Equal(1, len(order.Lines))
	suite.Equal(req.Lines[0].RacketID, order.Lines[0].RacketID)
	suite.Equal(req.Lines[0].Quantity, order.Lines[0].Quantity)
	suite.Nil(err)
}

// Cancel Order
func (suite *OrderServiceTestSuite) TestCancelOrderFail() {
	
	userID := "userID"
	orderID := "orderID"

	suite.mockRepo.On("GetOrderByID", mock.Anything, orderID, false).
		Return(&model.Order{
			UserID:     userID,
			TotalPrice: 111.1,
			Status:     model.OrderStatusNew,
		}, nil).Times(1)

	suite.mockRepo.On("UpdateOrder", mock.Anything, &model.Order{
		UserID:     userID,
		TotalPrice: 111.1,
		Status:     model.OrderStatusCancelled,
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
			UserID:     "userID1",
			TotalPrice: 111.1,
			Status:     model.OrderStatusNew,
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
			UserID:     userID,
			TotalPrice: 111.1,
			Status:     model.OrderStatusCancelled,
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
			UserID:     userID,
			TotalPrice: 111.1,
			Status:     model.OrderStatusNew,
		}, nil).Times(1)

	suite.mockRepo.On("UpdateOrder", mock.Anything, &model.Order{
		UserID:     userID,
		TotalPrice: 111.1,
		Status:     model.OrderStatusCancelled,
	}).Return(nil).Times(1)

	order, err := suite.service.CancelOrder(context.Background(), orderID, userID)
	suite.NotNil(order)
	suite.Equal(userID, order.UserID)
	suite.Equal(111.1, order.TotalPrice)
	suite.Equal(model.OrderStatusCancelled, order.Status)
	suite.Nil(err)
}