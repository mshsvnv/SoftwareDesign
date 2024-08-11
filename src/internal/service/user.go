package service

import (
	"context"
	"fmt"

	"src/internal/dto"
	"src/internal/model"
	repo "src/internal/repository"
	"src/pkg/logging"
)

//go:generate mockery --name=IUserService
type IUserService interface {
	GetUserByID(ctx context.Context, id int) (*model.User, error)
	GetAllUsers(ctx context.Context) ([]*model.User, error)
	UpdateRole(ctx context.Context, req dto.UpdateRoleReq) (*model.User, error)
}

type UserService struct {
	logger logging.Interface
	repo   repo.IUserRepository
}

func NewUserService(
	logger logging.Interface,
	repo repo.IUserRepository) *UserService {
	return &UserService{
		logger: logger,
		repo:   repo,
	}
}

func (s *UserService) UpdateRole(ctx context.Context, req dto.UpdateRoleReq) (*model.User, error) {

	s.logger.Infof("update role email %s", req.Email)
	user, err := s.repo.GetUserByEmail(ctx, req.Email)

	if err != nil {
		s.logger.Errorf("get user by email fail, error %s", err.Error())
		return nil, fmt.Errorf("get user by email fail, error %s", err)
	}

	user.Role = req.Role

	err = s.repo.UpdateRole(ctx, user)

	if err != nil {
		s.logger.Errorf("update fail, error %s", err.Error())
		return nil, fmt.Errorf("update fail, error %s", err)
	}

	return user, nil
}

func (s *UserService) GetAllUsers(ctx context.Context) ([]*model.User, error) {

	s.logger.Infof("get all users")
	users, err := s.repo.GetAllUsers(ctx)

	if err != nil {
		s.logger.Errorf("get all users fail, error %s", err.Error())
		return nil, fmt.Errorf("get all users fail, error %s", err)
	}

	return users, nil
}

func (s *UserService) GetUserByID(ctx context.Context, id int) (*model.User, error) {

	s.logger.Infof("get user by id %d", id)
	user, err := s.repo.GetUserByID(ctx, id)

	if err != nil {
		s.logger.Errorf("get user by id fail, error %s", err.Error())
		return nil, fmt.Errorf("get user by id fail, error %s", err)
	}

	return user, nil
}
