package service

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	"src_new/internal/dto"
	"src_new/internal/model"
	"src_new/internal/repository/mocks"
)

type RacketServiceTestSuite struct {
	suite.Suite
	mockRepo         *mocks.IRacketRepository
	mockRepoSupplier *mocks.ISupplierRepository
	service          IRacketService
}

func (suite *RacketServiceTestSuite) SetupTest() {
	suite.mockRepo = mocks.NewIRacketRepository(suite.T())
	suite.service = NewRacketService(suite.mockRepo, suite.mockRepoSupplier)
}

func TestRacketServiceTestSuite(t *testing.T) {
	suite.Run(t, new(RacketServiceTestSuite))
}

// CreateRacket
func (suite *RacketServiceTestSuite) TestCreateRacketFail() {

	req := &dto.CreateRacketReq{
		Brand:    "Wilson",
		Quantity: 10,
		Price:    100,
	}

	suite.mockRepo.On("Create", mock.Anything, mock.Anything).
		Return(errors.New("error")).Times(1)

	racket, err := suite.service.CreateRacket(context.Background(), req)

	suite.Nil(racket)
	suite.NotNil(err)
}

func (suite *RacketServiceTestSuite) TestCreateRacketSuccess() {

	req := &dto.CreateRacketReq{
		Brand:    "Wilson",
		Quantity: 10,
		Price:    100,
	}

	suite.mockRepo.On("Create", mock.Anything, mock.Anything).
		Return(nil).Times(1)

	racket, err := suite.service.CreateRacket(context.Background(), req)

	suite.Nil(err)
	suite.NotNil(racket)
}

// RemoveRacket
func (suite *RacketServiceTestSuite) TestRemoveRacketGetRacketFail() {

	id := 0

	suite.mockRepo.On("GetRacketByID", mock.Anything, id).
		Return(nil, errors.New("error")).Times(1)

	err := suite.service.RemoveRacket(context.Background(), id)

	suite.NotNil(err)
}

func (suite *RacketServiceTestSuite) TestRemoveRacketFail() {

	id := 0

	suite.mockRepo.On("GetRacketByID", mock.Anything, id).
		Return(&model.Racket{
			ID: 0,
		}, nil).Times(1)

	suite.mockRepo.On("Remove", mock.Anything, id).
		Return(errors.New("error")).Times(1)

	err := suite.service.RemoveRacket(context.Background(), id)

	suite.NotNil(err)
}

func (suite *RacketServiceTestSuite) TestRemoveRacketSuccess() {

	id := 0

	suite.mockRepo.On("GetRacketByID", mock.Anything, id).
		Return(&model.Racket{
			ID: 0,
		}, nil).Times(1)

	suite.mockRepo.On("Remove", mock.Anything, id).
		Return(nil).Times(1)

	err := suite.service.RemoveRacket(context.Background(), id)

	suite.Nil(err)
}

// UpdateRacket
func (suite *RacketServiceTestSuite) TestUpdateRacketGetRacketFail() {

	suite.mockRepo.On("GetRacketByID", mock.Anything, mock.Anything).
		Return(nil, errors.New("error")).Times(1)

	err := suite.service.UpdateRacket(context.Background(),
		&dto.UpdateRacketReq{})

	suite.NotNil(err)
}
func (suite *RacketServiceTestSuite) TestUpdateRacketGetFail() {

	req := &dto.UpdateRacketReq{
		ID:    0,
		Price: 200,
	}

	suite.mockRepo.On("GetRacketByID", mock.Anything, req.ID).
		Return(&model.Racket{
			ID:    0,
			Price: 100,
		}, nil).Times(1)

	suite.mockRepo.On("Update", mock.Anything,
		&model.Racket{
			ID:    0,
			Price: 200,
		}).Return(errors.New("error")).Times(1)

	err := suite.service.UpdateRacket(context.Background(), req)

	suite.NotNil(err)
}

func (suite *RacketServiceTestSuite) TestUpdateRacketSuccess() {

	req := &dto.UpdateRacketReq{
		ID:    0,
		Price: 200,
	}

	suite.mockRepo.On("GetRacketByID", mock.Anything, req.ID).
		Return(&model.Racket{
			ID:    0,
			Price: 100,
		}, nil).Times(1)

	suite.mockRepo.On("Update", mock.Anything,
		&model.Racket{
			ID:    0,
			Price: 200,
		}).Return(nil).Times(1)

	err := suite.service.UpdateRacket(context.Background(), req)

	suite.Nil(err)
}

// GetRacketByID
func (suite *RacketServiceTestSuite) TestGetRacketFail() {

	id := 0

	suite.mockRepo.On("GetRacketByID", mock.Anything, id).
		Return(nil, errors.New("error")).Times(1)

	racket, err := suite.service.GetRacketByID(context.Background(), id)

	suite.NotNil(err)
	suite.Nil(racket)
}

func (suite *RacketServiceTestSuite) TestGetRacketSuccess() {

	id := 0

	suite.mockRepo.On("GetRacketByID", mock.Anything, id).
		Return(&model.Racket{
			ID: 0,
		}, nil).Times(1)

	racket, err := suite.service.GetRacketByID(context.Background(), id)

	suite.Nil(err)
	suite.NotNil(racket)
}

// GetAllRackets
func (suite *RacketServiceTestSuite) TestGetAllRacketsFail() {

	suite.mockRepo.On("GetAllRackets", mock.Anything).
		Return(nil, errors.New("error")).Times(1)

	rackets, err := suite.service.GetAllRackets(context.Background())

	suite.NotNil(err)
	suite.Nil(rackets)
}

func (suite *RacketServiceTestSuite) TestGetAllRacketsSuccess() {

	suite.mockRepo.On("GetAllRackets", mock.Anything).
		Return([]*model.Racket{
			{
				ID:       0,
				Quantity: 10,
			},
			{
				ID:       1,
				Quantity: 20,
			},
		}, nil).Times(1)

	rackets, err := suite.service.GetAllRackets(context.Background())

	suite.NotNil(rackets)
	suite.Nil(err)
}
