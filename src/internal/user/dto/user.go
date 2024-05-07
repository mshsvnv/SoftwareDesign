package dto

type User struct {
	ID    string
	Email string
}

type RegisterReq struct {
	Email    string
	Password string
}

type LoginReq struct {
	Email    string
	Password string
}

type ChangePasswordReq struct {
	Password    string
	NewPassword string
}
