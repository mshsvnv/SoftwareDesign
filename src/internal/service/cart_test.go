package service

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	"src/internal/dto"
	"src/internal/model"
	"src/internal/repository/mocks"
)

type CartServiceTestSuite struct {
	suite.Suite
	mockRepo       *mocks.ICartRepository
	mockRepoRacket *mocks.IRacketRepository
	service        ICartService
}

func (suite *CartServiceTestSuite) SetupTest() {
	suite.mockRepo = mocks.NewICartRepository(suite.T())
	suite.mockRepoRacket = mocks.NewIRacketRepository(suite.T())
	suite.service = NewCartService(nil, suite.mockRepo, suite.mockRepoRacket)
}

func TestCartServiceTestSuite(t *testing.T) {
	suite.Run(t, new(CartServiceTestSuite))
}

// AddRacket
func (suite *CartServiceTestSuite) TestAddRacketGetCartFail() {

	req := &dto.AddRacketCartReq{
		UserID:   1,
		RacketID: 10,
		Quantity: 12,
	}

	racket := &model.Racket{
		ID:       req.RacketID,
		Quantity: 100,
		Price:    10,
	}

	suite.mockRepo.On("GetCartByID", mock.Anything, req.UserID).
		Return(nil, errors.New("error")).Times(1)

	suite.mockRepoRacket.On("GetRacketByID", mock.Anything, req.RacketID).
		Return(racket, nil).Times(1)

	suite.mockRepoRacket.On("Update", mock.Anything, racket).
		Return(nil).Times(1)

	suite.mockRepo.On("Create", mock.Anything,
		&model.Cart{
			UserID:     req.UserID,
			TotalPrice: float32(req.Quantity) * float32(racket.Price),
			Quantity:   req.Quantity,
			Lines: []*model.CartLine{{
				RacketID: req.RacketID,
				Quantity: req.Quantity,
			}},
		}).
		Return(errors.New("error")).Times(1)

	cart, err := suite.service.AddRacket(context.Background(), req)

	suite.Nil(cart)
	suite.NotNil(err)
}

func (suite *CartServiceTestSuite) TestAddRacketCreateSuccess() {

	req := &dto.AddRacketCartReq{
		UserID:   1,
		RacketID: 10,
		Quantity: 12,
	}

	racket := &model.Racket{
		ID:       req.RacketID,
		Quantity: 100,
		Price:    10,
	}

	suite.mockRepo.On("GetCartByID", mock.Anything, req.UserID).
		Return(nil, errors.New("error")).Times(1)

	suite.mockRepoRacket.On("GetRacketByID", mock.Anything, req.RacketID).
		Return(racket, nil).Times(1)

	suite.mockRepoRacket.On("Update", mock.Anything, racket).
		Return(nil).Times(1)

	suite.mockRepo.On("Create", mock.Anything,
		&model.Cart{
			UserID:     req.UserID,
			TotalPrice: float32(req.Quantity) * float32(racket.Price),
			Quantity:   req.Quantity,
			Lines: []*model.CartLine{{
				RacketID: req.RacketID,
				Quantity: req.Quantity,
			}},
		}).
		Return(nil).Times(1)

	cart, err := suite.service.AddRacket(context.Background(), req)

	suite.NotNil(cart)
	suite.Nil(err)
}

func (suite *CartServiceTestSuite) TestAddRacketSuccess() {

	req := &dto.AddRacketCartReq{
		UserID:   1,
		RacketID: 10,
		Quantity: 12,
	}

	racket := &model.Racket{
		ID:       req.RacketID,
		Quantity: 100,
		Price:    10,
	}

	cart := &model.Cart{
		UserID:     req.UserID,
		TotalPrice: float32(req.Quantity) * float32(racket.Price),
		Quantity:   req.Quantity,
		Lines: []*model.CartLine{{
			RacketID: req.RacketID,
			Quantity: req.Quantity,
		}},
	}

	suite.mockRepo.On("GetCartByID", mock.Anything, req.UserID).
		Return(cart, nil).Times(1)

	cart, err := suite.service.AddRacket(context.Background(), req)

	suite.NotNil(cart)
	suite.Nil(err)
}

func (suite *CartServiceTestSuite) TestGetCartByUserIDFail() {

	userID := 0

	cart := &model.Cart{
		UserID: userID,
	}

	suite.mockRepo.On("GetCartByID", mock.Anything, userID).
		Return(nil, errors.New("errors")).Times(1)

	suite.mockRepo.On("Create", mock.Anything, cart).
		Return(errors.New("errors")).Times(1)

	cart, err := suite.service.GetCartByID(context.Background(), userID)

	suite.Nil(cart)
	suite.NotNil(err)
}

func (suite *CartServiceTestSuite) TestGetCartByUserIDCreateSuccess() {

	userID := 0

	cart := &model.Cart{
		UserID: userID,
	}

	suite.mockRepo.On("GetCartByID", mock.Anything, userID).
		Return(nil, errors.New("errors")).Times(1)

	suite.mockRepo.On("Create", mock.Anything, cart).
		Return(nil).Times(1)

	cart, err := suite.service.GetCartByID(context.Background(), userID)

	suite.Nil(err)
	suite.NotNil(cart)
}
func (suite *CartServiceTestSuite) TestGetCartByUserIDSuccess() {

	userID := 0

	cart := &model.Cart{
		UserID: userID,
	}

	suite.mockRepo.On("GetCartByID", mock.Anything, userID).
		Return(cart, nil).Times(1)

	cart, err := suite.service.GetCartByID(context.Background(), userID)

	suite.NotNil(cart)
	suite.Nil(err)
}
