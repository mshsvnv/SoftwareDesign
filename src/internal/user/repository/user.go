package repository

import "src/internal/user/model"

type IUserService interface {
	Create() error
	Update() error
	GetUserByID(id int) *model.User
}

type UserRepo struct {
	db dbs.IDatabase
}
