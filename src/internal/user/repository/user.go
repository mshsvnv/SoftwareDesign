package repository

import (
	"context"
	"src/internal/user/model"
)

//go:generate mockery --name=IUserRepository
type IUserRepository interface {
	Create(ctx context.Context, user *model.User) error
	Update(ctx context.Context, user *model.User) error
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
	GetUserByID(ctx context.Context, id string) (*model.User, error)
}

type UserRepository struct {
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

// func (r *UserRepository) GetUserByEmail(ctx context.Context, id string) (*model.User, error) {
// 	return &model.User{
// 		ID:       "1",
// 		Email:    "test@email.com",
// 		Password: "test123456",
// 		Role:     model.UserRoleCustomer,
// 	}, nil
// }

// func (r *UserRepository) Create(ctx context.Context, user *model.User) error {
// 	return nil
// }

// func (r *UserRepository) Update(ctx context.Context, user *model.User) error {
// 	return nil
// }
