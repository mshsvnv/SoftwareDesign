package service

import (
	"context"
	"fmt"

	"golang.org/x/crypto/bcrypt"

	"src_new/internal/dto"
	"src_new/internal/model"
	repo "src_new/internal/repository"
	"src_new/pkg/utils"
)

//go:generate mockery --name=IUserService
type IUserService interface {
	Login(ctx context.Context, req *dto.LoginReq) (*model.User, error)
	Register(ctx context.Context, req *dto.RegisterReq) (*model.User, error)
	GetUserByID(ctx context.Context, id int) (*model.User, error)
	// GetAllUsers()
}

type UserService struct {
	repo repo.IUserRepository
}

func NewUserService(repo repo.IUserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) Login(ctx context.Context, req *dto.LoginReq) (*model.User, error) {

	user, err := s.repo.GetUserByEmail(ctx, req.Email)

	if err != nil {
		return nil, fmt.Errorf("Login.GetUserByEmail fail, email: %s, error: %s", req.Email, err)
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(req.Password))

	if err != nil {
		return nil, fmt.Errorf("Login.CompareHashAndPassword fail, wrong password: %s", req.Password)
	}

	return user, nil
}

func (s *UserService) Register(ctx context.Context, req *dto.RegisterReq) (*model.User, error) {

	user, err := s.repo.GetUserByEmail(ctx, req.Email)

	if user != nil {
		return nil, fmt.Errorf("Register.GetUserByID fail, email %s, error %s", req.Email, err)
	}

	user = &model.User{
		ID:       0,
		Name:     req.Name,
		Surname:  req.Surname,
		Email:    req.Email,
		Password: utils.HashAndSalt([]byte(req.Password)),
	}

	err = s.repo.Create(ctx, user)

	if err != nil {
		return nil, fmt.Errorf("Register.Create fail, email %s, error %s", req.Email, err)
	}

	return user, nil
}

func (s *UserService) GetUserByID(ctx context.Context, id int) (*model.User, error) {

	user, err := s.repo.GetUserByID(ctx, id)

	if err != nil {
		return nil, fmt.Errorf("GetUserByID fail, ID %d, error %s", id, err)
	}

	return user, nil
}
