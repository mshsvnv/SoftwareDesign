package dto

import (
	"src/internal/model"
)

type User struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	// CreatedAt time.Time `json:"created_at"`
	// UpdatedAt time.Time `json:"updated_at"`
}

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRes struct {
	User             User   `json:"user"`
	AccessToken      string `json:"access_token"`
	RefreshUserToken string `json:"refresh_token"`
}

type RegisterReq struct {
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Email    string `json:"email"`
	Password string `json:"password"`
	// Role     string `json:""`
}

type RegisterRes struct {
	User User `json:"user"`
}

type RefreshTokenReq struct {
	RefreshToken string `json:"refresh_token"`
}

type RefreshTokenRes struct {
	AccessToken string `json:"access_token"`
}

type UpdateRoleReq struct {
	Email string
	Role  model.UserRole
}
