package service

import (
	"fmt"

	"src/internal/user/model"
)

type IuserService interface {
	Login()
	Register()
	GetUserByID(id int) *model.User
	ChangePassword()
}

type UserService struct {
	ID int
}

func NewUserService() *UserService {
	return &UserService{ID: 1}
}

func (s *UserService) Login() {
	fmt.Print(s.ID)
}

func (s *UserService) Register() {
	fmt.Print(s.ID)
}

// func (s *UserService) GetUserByID(id int) (*model.User) {

// }
