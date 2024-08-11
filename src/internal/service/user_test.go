package service

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	"src/internal/model"
	"src/internal/repository/mocks"
)

type UserServiceTestSuite struct {
	suite.Suite
	mockRepo *mocks.IUserRepository
	service  IUserService
}

func (suite *UserServiceTestSuite) SetupTest() {
	suite.mockRepo = mocks.NewIUserRepository(suite.T())
	suite.service = NewUserService(nil, suite.mockRepo)
}

func TestUserServiceTestSuite(t *testing.T) {
	suite.Run(t, new(UserServiceTestSuite))
}

// GetUserByID
func (suite *UserServiceTestSuite) TestGetUserByIDFail() {

	id := 0

	suite.mockRepo.On("GetUserByID", mock.Anything, id).
		Return(nil, errors.New("error")).Times(1)

	user, err := suite.service.GetUserByID(context.Background(), id)

	suite.Nil(user)
	suite.NotNil(err)
}

func (suite *UserServiceTestSuite) TestGetUserByIDSuccess() {

	id := 0

	suite.mockRepo.On("GetUserByID", mock.Anything, id).
		Return(&model.User{
			ID:    id,
			Email: "test@test.com",
		}, nil).Times(1)

	user, err := suite.service.GetUserByID(context.Background(), id)

	suite.Nil(err)
	suite.NotNil(user)
}
