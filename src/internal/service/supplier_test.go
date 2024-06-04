package service

import (
	"context"
	"errors"
	"testing"

	"src/internal/dto"
	"src/internal/model"
	"src/internal/repository/mocks"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type SupplierServiceTestSuite struct {
	suite.Suite
	mockRepo *mocks.ISupplierRepository
	service  ISupplierService
}

func (suite *SupplierServiceTestSuite) SetupTest() {
	suite.mockRepo = mocks.NewISupplierRepository(suite.T())
	suite.service = NewSupplierService(nil, suite.mockRepo)
}

func TestSupplierServiceTestSuite(t *testing.T) {
	suite.Run(t, new(SupplierServiceTestSuite))
}

// CreateSupplier
func (suite *SupplierServiceTestSuite) TestCreateSupplierFail() {

	req := &dto.CreateSupplierReq{
		Name:  "IP Sidorov",
		Phone: "8-499-676-14-95",
	}

	suite.mockRepo.On("Create", mock.Anything, mock.Anything).
		Return(errors.New("error")).Times(1)

	supplier, err := suite.service.CreateSupplier(context.Background(), req)

	suite.Nil(supplier)
	suite.NotNil(err)
}

func (suite *SupplierServiceTestSuite) TestCreateSupplierSuccess() {

	req := &dto.CreateSupplierReq{
		Name:  "IP Sidorov",
		Phone: "8-499-676-14-95",
	}

	suite.mockRepo.On("Create", mock.Anything,
		&model.Supplier{
			Name:  "IP Sidorov",
			Phone: "8-499-676-14-95",
		}).
		Return(nil).Times(1)

	supplier, err := suite.service.CreateSupplier(context.Background(), req)

	suite.Nil(err)
	suite.NotNil(supplier)
}

// RemoveSupplier
func (suite *SupplierServiceTestSuite) TestRemoveSupplierGetSupplierFail() {

	suite.mockRepo.On("GetSupplierByID", mock.Anything, "@mail.ru").
		Return(nil, errors.New("error")).Times(1)

	err := suite.service.RemoveSupplier(context.Background(), "@mail.ru")

	suite.NotNil(err)
}

func (suite *SupplierServiceTestSuite) TestRemoveSupplierFail() {

	id := 0

	suite.mockRepo.On("GetSupplierByID", mock.Anything, id).
		Return(&model.Supplier{
			Email: "@mail.ru",
		}, nil).Times(1)

	suite.mockRepo.On("Remove", mock.Anything, "@mail.ru").
		Return(errors.New("error")).Times(1)

	err := suite.service.RemoveSupplier(context.Background(), "@mail.ru")

	suite.NotNil(err)
}

func (suite *SupplierServiceTestSuite) TestRemoveSupplierSuccess() {

	id := 0

	suite.mockRepo.On("GetSupplierByID", mock.Anything, id).
		Return(&model.Supplier{
			Email: "@mail.ru",
		}, nil).Times(1)

	suite.mockRepo.On("Remove", mock.Anything, "@mail.ru").
		Return(nil).Times(1)

	err := suite.service.RemoveSupplier(context.Background(), "@mail.ru")

	suite.Nil(err)
}

// GetSupplierByID
func (suite *SupplierServiceTestSuite) TestGetSupplierByIDFail() {

	id := 0

	suite.mockRepo.On("GetSupplierByID", mock.Anything, id).
		Return(nil, errors.New("error")).Times(1)

	user, err := suite.service.GetSupplierByID(context.Background(), id)

	suite.Nil(user)
	suite.NotNil(err)
}

func (suite *SupplierServiceTestSuite) TestGetSupplierByIDSuccess() {

	id := 0

	suite.mockRepo.On("GetSupplierByID", mock.Anything, id).
		Return(&model.Supplier{
			Email: "test@test.com",
		}, nil).Times(1)

	user, err := suite.service.GetSupplierByID(context.Background(), id)

	suite.Nil(err)
	suite.NotNil(user)
}
