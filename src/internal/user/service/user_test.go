package service

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	"src/internal/user/dto"
	"src/internal/user/model"
	"src/internal/user/repository/mocks"
	"src/pkg/utils"
)

type UserServiceTestSuite struct {
	suite.Suite

	mockRepo *mocks.IUserRepository
	service  IUserService
}

func (suite *UserServiceTestSuite) SetupTest() {
	suite.mockRepo = mocks.NewIUserRepository(suite.T())
	suite.service = NewUserService(suite.mockRepo)
}

func TestUserServiceTestSuite(t *testing.T) {
	suite.Run(t, new(UserServiceTestSuite))
}

// Login

func (suite *UserServiceTestSuite) TestLoginFail() {

	req := &dto.LoginReq{
		Email:    "test@email.com",
		Password: "test123456",
	}

	suite.mockRepo.On("GetUserByEmail", mock.Anything, req.Email).
		Return(nil, errors.New("error")).Times(1)

	user, err := suite.service.Login(context.Background(), req)

	suite.Nil(user)
	suite.NotNil(err)
}

func (suite *UserServiceTestSuite) TestLoginWrongPassword() {

	req := &dto.LoginReq{
		Email:    "test@test.com",
		Password: "test123456",
	}

	suite.mockRepo.On("GetUserByEmail", mock.Anything, req.Email).
		Return(&model.User{
			Email:    "test@test.com",
			Password: "password",
		}, nil).Times(1)

	user, err := suite.service.Login(context.Background(), req)

	suite.Nil(user)
	suite.NotNil(err)
}

func (suite *UserServiceTestSuite) TestLoginSuccess() {

	req := &dto.LoginReq{
		Email:    "test@test.com",
		Password: "test123456",
	}

	suite.mockRepo.On("GetUserByEmail", mock.Anything, req.Email).
		Return(&model.User{
			Email:    "test@test.com",
			Password: utils.HashAndSalt([]byte("test123456")),
		}, nil).Times(1)

	user, err := suite.service.Login(context.Background(), req)

	suite.Nil(user)
	suite.NotNil(err)
}

// Register
func (suite *UserServiceTestSuite) TestRegisterFail() {

	req := &dto.RegisterReq{
		Email:    "test@test.com",
		Password: "test123456",
	}

	suite.mockRepo.On("GetUserByEmail", mock.Anything, req.Email).
		Return(nil, errors.New("error")).Times(1)

	suite.mockRepo.On("Create", mock.Anything, mock.Anything).
		Return(errors.New("error")).Times(1)

	user, err := suite.service.Register(context.Background(), req)

	suite.Nil(user)
	suite.NotNil(err)
}
func (suite *UserServiceTestSuite) TestRegisterAlreadyRegistered() {

	req := &dto.RegisterReq{
		Email:    "test@test.com",
		Password: "test123456",
	}

	suite.mockRepo.On("GetUserByEmail", mock.Anything, req.Email).
		Return(&model.User{
			Email:    "test@test.com",
			Password: utils.HashAndSalt([]byte("test123456")),
		}, nil).Times(1)

	user, err := suite.service.Register(context.Background(), req)

	suite.Nil(user)
	suite.NotNil(err)
}
func (suite *UserServiceTestSuite) TestRegisterSuccess() {

	req := &dto.RegisterReq{
		Email:    "test@test.com",
		Password: "test123456",
	}

	suite.mockRepo.On("GetUserByEmail", mock.Anything, req.Email).
		Return(nil, errors.New("error")).Times(1)

	suite.mockRepo.On("Create", mock.Anything, mock.Anything).
		Return(nil).Times(1)

	user, err := suite.service.Register(context.Background(), req)

	suite.NotNil(user)
	suite.Nil(err)
}
