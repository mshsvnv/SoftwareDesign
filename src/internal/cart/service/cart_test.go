package service

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	"src/internal/cart/dto"
	"src/internal/cart/model"
	"src/internal/cart/repository/mocks"
)

type CartServiceTestSuite struct {
	suite.Suite

	mockRepo *mocks.ICartRepository
	service  ICartService
}

func (suite *CartServiceTestSuite) SetupTest() {
	suite.mockRepo = mocks.NewICartRepository(suite.T())
	suite.service = NewCartService(suite.mockRepo)
}

func TestCartServiceTestSuite(t *testing.T) {
	suite.Run(t, new(CartServiceTestSuite))
}

// GetCartByUserID
func (suite *CartServiceTestSuite) TestGetCartByUserIDFail() {

	userID := "userID"

	suite.mockRepo.On("GetCartByUserID", mock.Anything, userID).
		Return(nil, errors.New("error")).Times(1)

	suite.mockRepo.On("Create", mock.Anything, &model.Cart{
		UserID: "userID",
	}).Return(nil).Times(1)

	cart, err := suite.service.GetCartByUserID(context.Background(), userID)

	suite.NotNil(cart)
	suite.Equal("userID", cart.UserID)
	suite.Equal(0, len(cart.Lines))
	suite.Nil(err)
}

func (suite *CartServiceTestSuite) TestGetCartByUserIDCreateFail() {

	userID := "userID"

	suite.mockRepo.On("GetCartByUserID", mock.Anything, userID).
		Return(nil, errors.New("error")).Times(1)

	suite.mockRepo.On("Create", mock.Anything, &model.Cart{
		UserID: "userID",
	}).Return(errors.New("error")).Times(1)

	cart, err := suite.service.GetCartByUserID(context.Background(), userID)
	suite.Nil(cart)
	suite.NotNil(err)
}

func (suite *CartServiceTestSuite) TestGetCartByUserIDSuccess() {
	userID := "userID"

	suite.mockRepo.On("GetCartByUserID", mock.Anything, userID).
		Return(
			&model.Cart{
				ID:     "cartId1",
				UserID: "userID",
				Lines: []*model.CartLine{
					{
						RacketID: "RacketID1",
						Quantity: 4,
					},
					{
						RacketID: "RacketID2",
						Quantity: 3,
					},
				},
			},
			nil,
		).Times(1)

	cart, err := suite.service.GetCartByUserID(context.Background(), userID)
	suite.NotNil(cart)
	suite.Equal("userID", cart.UserID)
	suite.Equal(2, len(cart.Lines))
	suite.Nil(err)
}

// AddRacket
func (suite *CartServiceTestSuite) TestAddRacketCartNotFound() {

	req := &dto.AddRacketReq{
		UserID: "userID",
		Line: &dto.CartLineReq{
			RacketID: "RacketID2",
			Quantity: 3,
		},
	}

	suite.mockRepo.On("GetCartByUserID", mock.Anything, "userID").
		Return(nil, errors.New("error")).Times(1)

	suite.mockRepo.On("Create", mock.Anything, &model.Cart{
		UserID: "userID",
		Lines: []*model.CartLine{
			{
				RacketID: "RacketID2",
				Quantity: 3,
			},
		},
	}).Return(nil).Times(1)

	cart, err := suite.service.AddRacket(context.Background(), req)

	suite.NotNil(cart)
	suite.Equal("userID", cart.UserID)
	suite.Equal(1, len(cart.Lines))
	suite.Nil(err)
}
func (suite *CartServiceTestSuite) TestAddRacketCartNotFoundCreateFail() {

	req := &dto.AddRacketReq{
		UserID: "userID",
		Line: &dto.CartLineReq{
			RacketID: "RacketID2",
			Quantity: 3,
		},
	}

	suite.mockRepo.On("GetCartByUserID", mock.Anything, "userID").
		Return(nil, errors.New("error")).Times(1)

	suite.mockRepo.On("Create", mock.Anything, &model.Cart{
		UserID: "userID",
		Lines: []*model.CartLine{
			{
				RacketID: "RacketID2",
				Quantity: 3,
			},
		},
	}).Return(errors.New("error")).Times(1)

	cart, err := suite.service.AddRacket(context.Background(), req)

	suite.Nil(cart)
	suite.NotNil(err)
}

func (suite *CartServiceTestSuite) TestAddRacketAlreadyExistInCart() {

	req := &dto.AddRacketReq{
		UserID: "userID",
		Line: &dto.CartLineReq{
			RacketID: "RacketID2",
			Quantity: 3,
		},
	}

	suite.mockRepo.On("GetCartByUserID", mock.Anything, "userID").
		Return(
			&model.Cart{
				ID:     "cartId1",
				UserID: "userID",
				Lines: []*model.CartLine{
					{
						RacketID: "RacketID1",
						Quantity: 4,
					},
					{
						RacketID: "RacketID2",
						Quantity: 3,
					},
				},
			},
			nil,
		).Times(1)

	cart, err := suite.service.AddRacket(context.Background(), req)
	suite.NotNil(cart)
	suite.Equal("userID", cart.UserID)
	suite.Equal(2, len(cart.Lines))
	suite.Nil(err)
}

func (suite *CartServiceTestSuite) TestAddRacketUpdateFail() {

	req := &dto.AddRacketReq{
		UserID: "userID",
		Line: &dto.CartLineReq{
			RacketID: "RacketID2",
			Quantity: 3,
		},
	}

	suite.mockRepo.On("GetCartByUserID", mock.Anything, "userID").
		Return(
			&model.Cart{
				ID:     "cartId1",
				UserID: "userID",
				Lines: []*model.CartLine{
					{
						RacketID: "RacketID1",
						Quantity: 4,
					},
				},
			},
			nil,
		).Times(1)

	suite.mockRepo.On("Update", mock.Anything, &model.Cart{
		ID:     "cartId1",
		UserID: "userID",
		Lines: []*model.CartLine{
			{
				RacketID: "RacketID1",
				Quantity: 4,
			},
			{
				RacketID: "RacketID2",
				Quantity: 3,
			},
		},
	}).Return(errors.New("error")).Times(1)

	cart, err := suite.service.AddRacket(context.Background(), req)
	suite.Nil(cart)
	suite.NotNil(err)
}

func (suite *CartServiceTestSuite) TestAddRacketSuccess() {

	req := &dto.AddRacketReq{
		UserID: "userID",
		Line: &dto.CartLineReq{
			RacketID: "RacketID2",
			Quantity: 3,
		},
	}

	suite.mockRepo.On("GetCartByUserID", mock.Anything, "userID").
		Return(
			&model.Cart{
				ID:     "cartId1",
				UserID: "userID",
				Lines: []*model.CartLine{
					{
						RacketID: "RacketID1",
						Quantity: 4,
					},
				},
			},
			nil,
		).Times(1)

	suite.mockRepo.On("Update", mock.Anything, &model.Cart{
		ID:     "cartId1",
		UserID: "userID",
		Lines: []*model.CartLine{
			{
				RacketID: "RacketID1",
				Quantity: 4,
			},
			{
				RacketID: "RacketID2",
				Quantity: 3,
			},
		},
	}).Return(nil).Times(1)

	cart, err := suite.service.AddRacket(context.Background(), req)

	suite.NotNil(cart)
	suite.Equal("userID", cart.UserID)
	suite.Equal(2, len(cart.Lines))
	suite.Nil(err)
}

// RemoveRacket
func (suite *CartServiceTestSuite) TestRemoveRacketCartNotFound() {

	req := &dto.RemoveRacketReq{
		UserID:   "userID",
		RacketID: "RacketID1",
	}

	suite.mockRepo.On("GetCartByUserID", mock.Anything, "userID").
		Return(nil, errors.New("error")).Times(1)

	suite.mockRepo.On("Create", mock.Anything, &model.Cart{UserID: "userID"}).Return(nil).Times(1)

	cart, err := suite.service.RemoveRacket(context.Background(), req)

	suite.NotNil(cart)
	suite.Equal("userID", cart.UserID)
	suite.Equal(0, len(cart.Lines))
	suite.Nil(err)
}

func (suite *CartServiceTestSuite) TestRemoveRacketCartNotFoundCreateFail() {
	req := &dto.RemoveRacketReq{
		UserID:   "userID",
		RacketID: "RacketID1",
	}

	suite.mockRepo.On("GetCartByUserID", mock.Anything, "userID").
		Return(nil, errors.New("error")).Times(1)

	suite.mockRepo.On("Create", mock.Anything, &model.Cart{UserID: "userID"}).Return(errors.New("error")).Times(1)

	cart, err := suite.service.RemoveRacket(context.Background(), req)
	suite.Nil(cart)
	suite.NotNil(err)
}

func (suite *CartServiceTestSuite) TestRemoveRacketUpdateFail() {

	req := &dto.RemoveRacketReq{
		UserID:   "userID",
		RacketID: "RacketID1",
	}

	suite.mockRepo.On("GetCartByUserID", mock.Anything, "userID").
		Return(
			&model.Cart{
				ID:     "cartId1",
				UserID: "userID",
				Lines: []*model.CartLine{
					{
						RacketID: "RacketID1",
						Quantity: 4,
					},
				},
			},
			nil,
		).Times(1)

	suite.mockRepo.On("Update", mock.Anything, &model.Cart{
		ID:     "cartId1",
		UserID: "userID",
		Lines:  []*model.CartLine{},
	}).Return(errors.New("error")).Times(1)

	cart, err := suite.service.RemoveRacket(context.Background(), req)
	suite.Nil(cart)
	suite.NotNil(err)
}

func (suite *CartServiceTestSuite) TestRemoveRacketSuccess() {

	req := &dto.RemoveRacketReq{
		UserID:   "userID",
		RacketID: "RacketID1",
	}

	suite.mockRepo.On("GetCartByUserID", mock.Anything, "userID").
		Return(
			&model.Cart{
				ID:     "cartId1",
				UserID: "userID",
				Lines: []*model.CartLine{
					{
						RacketID: "RacketID1",
						Quantity: 4,
					},
				},
			},
			nil,
		).Times(1)

	suite.mockRepo.On("Update", mock.Anything, &model.Cart{
		ID:     "cartId1",
		UserID: "userID",
		Lines:  []*model.CartLine{},
	}).Return(nil).Times(1)

	cart, err := suite.service.RemoveRacket(context.Background(), req)

	suite.NotNil(cart)
	suite.Equal("userID", cart.UserID)
	suite.Equal(0, len(cart.Lines))
	suite.Nil(err)
}
