package service

import (
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"

	"src/internal/user/dto"
	"src/internal/user/model"
	"src/internal/user/repository"
	"src/pkg/utils"
)

type IUserService interface {
	Login(ctx context.Context, req *dto.LoginReq) (*model.User, error)
	Register(ctx context.Context, req *dto.RegisterReq) (*model.User, error)
}

type UserService struct {
	repo repository.IUserRepository
}

func NewUserService(repo repository.IUserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) Login(ctx context.Context, req *dto.LoginReq) (*model.User, error) {

	user, err := s.repo.GetUserByEmail(ctx, req.Email)

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("user not found")
	}

	if err = bcrypt.CompareHashAndPassword([]byte(req.Password), []byte(user.Password)); err != nil {
		return nil, errors.New("wrong password")
	}

	return user, nil
}

func (s *UserService) Register(ctx context.Context, req *dto.RegisterReq) (*model.User, error) {

	user, _ := s.repo.GetUserByEmail(ctx, req.Email)

	if user != nil {
		return nil, errors.New("user already registered")
	}

	var newUser model.User
	utils.Copy(&req, &newUser)

	err := s.repo.Create(ctx, &newUser)

	if err != nil {
		return nil, err
	}

	return &newUser, nil
}
