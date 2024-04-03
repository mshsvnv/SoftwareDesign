package dto

type User struct {
	ID    string
	Email string
}

type RegisterReq struct {
	Email    string
	Password string
}

type RegisterRes struct {
	User User
}

type LoginReq struct {
	Email    string
	Password string
}

type LoginRes struct {
	User User
}

type ChangePasswordReq struct {
	Password    string
	NewPassword string
}
