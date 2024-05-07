package service

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	"src/internal/racket/dto"
	"src/internal/racket/model"
	"src/internal/racket/repository/mocks"

	cart "src/internal/cart/repository/mocks"
	order "src/internal/order/repository/mocks"
)

type RacketServiceTestSuite struct {
	suite.Suite

	mockRepo      *mocks.IRacketRepository
	mockRepoOrder *order.IOrderRepository
	mockRepoCart  *cart.ICartRepository

	service IRacketService
}

func (suite *RacketServiceTestSuite) SetupTest() {
	suite.mockRepo = mocks.NewIRacketRepository(suite.T())
	suite.mockRepoOrder = order.NewIOrderRepository(suite.T())
	suite.mockRepoCart = cart.NewICartRepository(suite.T())

	suite.service = NewRacketService(suite.mockRepo,
		suite.mockRepoCart,
		suite.mockRepoOrder)
}

func TestRacketServiceTestSuite(t *testing.T) {
	suite.Run(t, new(RacketServiceTestSuite))
}

// GetRacketByID
func (suite *RacketServiceTestSuite) TestGetRacketByIDFail() {

	racketID := "racketID"

	suite.mockRepo.On("GetRacketByID", mock.Anything, racketID).
		Return(nil, errors.New("error")).Times(1)

	racket, err := suite.service.GetRacketByID(context.Background(), racketID)

	suite.Nil(racket)
	suite.NotNil(err)
}

func (suite *RacketServiceTestSuite) TestGetRacketByIDSuccess() {

	racketID := "racketID"

	suite.mockRepo.On("GetRacketByID", mock.Anything, racketID).
		Return(
			&model.Racket{
				Brand: "racket",
				Price: 1.1,
			},
			nil,
		).Times(1)

	racket, err := suite.service.GetRacketByID(context.Background(), racketID)

	suite.NotNil(racket)
	suite.Equal("racket", racket.Brand)
	suite.Equal(1.1, racket.Price)
	suite.Nil(err)
}

// ListRackets
func (suite *RacketServiceTestSuite) TestListRacketsFail() {

	req := &dto.ListRacketReq{
		Brand: "racket",
	}

	suite.mockRepo.On("ListRackets", mock.Anything, req).
		Return(nil, errors.New("error")).Times(1)

	rackets, err := suite.service.ListRackets(context.Background(), req)

	suite.Nil(rackets)
	suite.NotNil(err)
}

func (suite *RacketServiceTestSuite) TestListRacketsSuccess() {

	req := &dto.ListRacketReq{
		Brand: "racket",
	}

	suite.mockRepo.On("ListRackets", mock.Anything, req).
		Return(
			[]*model.Racket{
				{
					Brand: "racket",
					Price: 1.1,
				},
			},
			nil,
		).Times(1)

	rackets, err := suite.service.ListRackets(context.Background(), req)

	suite.NotNil(rackets)
	suite.Equal(1, len(rackets))
	suite.Equal(1.1, rackets[0].Price)
	suite.Nil(err)
}

// Create
func (suite *RacketServiceTestSuite) TestCreateFail() {
	req := &dto.CreateRacketReq{
		Brand: "racket",
		Price: 1.1,
	}

	suite.mockRepo.On("Create", mock.Anything, &model.Racket{
		Brand: "racket",
		Price: 1.1,
	}).Return(errors.New("error")).Times(1)

	racket, err := suite.service.Create(context.Background(), req)

	suite.Nil(racket)
	suite.NotNil(err)
}

func (suite *RacketServiceTestSuite) TestCreateSuccess() {

	req := &dto.CreateRacketReq{
		Brand: "racket",
		Price: 1.1,
	}

	suite.mockRepo.On("Create", mock.Anything, &model.Racket{
		Brand: "racket",
		Price: 1.1,
	}).Return(nil).Times(1)

	racket, err := suite.service.Create(context.Background(), req)

	suite.NotNil(racket)
	suite.Equal(req.Brand, racket.Brand)
	suite.Equal(req.Price, racket.Price)
	suite.Nil(err)
}

// Update
func (suite *RacketServiceTestSuite) TestUpdateFail() {

	racketID := "racketID"

	req := &dto.UpdateRacketReq{
		Brand: "racket",
		Price: 1.1,
	}

	suite.mockRepo.On("GetRacketByID", mock.Anything, racketID).
		Return(&model.Racket{
			Brand: "racket",
			Price: 1.1,
		},
			nil).Times(1)

	suite.mockRepo.On("Update", mock.Anything, &model.Racket{
		Brand: "racket",
		Price: 1.1,
	}).Return(errors.New("error")).Times(1)

	racket, err := suite.service.Update(context.Background(), racketID, req)

	suite.Nil(racket)
	suite.NotNil(err)
}

func (suite *RacketServiceTestSuite) TestUpdateGetRacketByIDFail() {

	racketID := "racketID"

	req := &dto.UpdateRacketReq{
		Brand: "racket",
		Price: 1.1,
	}

	suite.mockRepo.On("GetRacketByID", mock.Anything, racketID).
		Return(nil, errors.New("error")).Times(1)

	racket, err := suite.service.Update(context.Background(), racketID, req)
	suite.Nil(racket)
	suite.NotNil(err)
}

func (suite *RacketServiceTestSuite) TestUpdateSuccess() {
	racketID := "racketID"
	req := &dto.UpdateRacketReq{
		Brand: "racket",
		Price: 1.1,
	}

	suite.mockRepo.On("GetRacketByID", mock.Anything, racketID).
		Return(&model.Racket{
			Brand: "racket",
			Price: 1.1,
		},
			nil).Times(1)

	suite.mockRepo.On("Update", mock.Anything, &model.Racket{
		Brand: "racket",
		Price: 1.1,
	}).Return(nil).Times(1)

	racket, err := suite.service.Update(context.Background(), racketID, req)

	suite.NotNil(racket)
	suite.Equal(req.Brand, racket.Brand)
	suite.Equal(req.Price, racket.Price)
	suite.Nil(err)
}

// Delete
func (suite *RacketServiceTestSuite) TestDeleteFailNoRacket() {

	id := "123"

	suite.mockRepo.On("GetRacketByID", mock.Anything, id).
		Return(nil, errors.New("racket not found")).Times(1)

	err := suite.service.Delete(context.Background(), id)

	suite.NotNil(err)
}

func (suite *RacketServiceTestSuite) TestDeleteCartFail() {

	id := "123"

	suite.mockRepo.On("GetRacketByID", mock.Anything, id).
		Return(&model.Racket{
			Brand: "racket",
			Price: 1.1,
		}, nil).Times(1)

	suite.mockRepoCart.On("DeleteByRacketID", mock.Anything, id).
		Return(errors.New("failed to delete associated carts")).Times(1)

	err := suite.service.Delete(context.Background(), id)

	suite.NotNil(err)
}

func (suite *RacketServiceTestSuite) TestDeleteOrderFail() {

	id := "123"

	suite.mockRepo.On("GetRacketByID", mock.Anything, id).
		Return(&model.Racket{
			Brand: "racket",
			Price: 1.1,
		}, nil).Times(1)

	suite.mockRepoCart.On("DeleteByRacketID", mock.Anything, id).
		Return(nil).Times(1)

	suite.mockRepoOrder.On("DeleteOrdersByRacketID", mock.Anything, id).
		Return(errors.New("failed to delete associated orders")).Times(1)

	err := suite.service.Delete(context.Background(), id)

	suite.NotNil(err)
}

func (suite *RacketServiceTestSuite) TestDeleteSuccess() {

	id := "123"

	suite.mockRepo.On("GetRacketByID", mock.Anything, id).
		Return(&model.Racket{
			Brand: "racket",
			Price: 1.1,
		}, nil).Times(1)

	suite.mockRepoCart.On("DeleteByRacketID", mock.Anything, id).
		Return(nil).Times(1)

	suite.mockRepoOrder.On("DeleteOrdersByRacketID", mock.Anything, id).
		Return(nil).Times(1)

	suite.mockRepo.On("Delete", mock.Anything, id).
		Return(nil).Times(1)

	err := suite.service.Delete(context.Background(), id)

	suite.Nil(err)
}
