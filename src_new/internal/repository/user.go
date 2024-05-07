package repository

import (
	"context"
	"src_new/internal/model"
)

//go:generate mockery --name=IUserRepository
type IUserRepository interface {
	Create(ctx context.Context, user *model.User) error
	GetUserByID(ctx context.Context, id int) (*model.User, error)
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
	Remove(ctx context.Context, email string) error
}

// type UserRepository struct {
// 	db []*model.User
// }

// func NewUserRepository() *UserRepository {
// 	return &UserRepository{}
// }

// func (r *UserRepository) Create(ctx context.Context, user *model.User) error {

// 	max := 0

// 	for i := 0; i < len(r.db); i++ {
// 		if r.db[i].ID > max {
// 			max = r.db[i].ID
// 		}
// 	}

// 	user.ID = max
// 	r.db = append(r.db, user)

// 	return nil
// }

// func (r *UserRepository) GetUserByID(ctx context.Context, id int) (*model.User, error) {

// 	for _, user := range r.db {

// 		if user.ID == id {
// 			return user, nil
// 		}
// 	}

// 	return nil, fmt.Errorf("Error")
// }

// func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {

// 	for _, user := range r.db {

// 		if user.Email == email {
// 			return user, nil
// 		}
// 	}

// 	return nil, fmt.Errorf("Error")
// }

// func (r *UserRepository) GetAllUsers() {

// 	for _, user := range r.db {
// 		fmt.Printf("ID: %d Email: %s\n", user.ID, user.Email)
// 	}
// }
