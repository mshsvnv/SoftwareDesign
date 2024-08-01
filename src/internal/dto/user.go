package dto

import (
	"src/internal/model"
)

// type User struct {
// 	ID    string `json:"id"`
// 	Email string `json:"email"`
// 	// CreatedAt time.Time `json:"created_at"`
// 	// UpdatedAt time.Time `json:"updated_at"`
// }

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterReq struct {
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateRoleReq struct {
	Email string
	Role  model.UserRole
}
