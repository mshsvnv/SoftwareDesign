package dto

type LoginReq struct {
	Id       int
	Email    string
	Password string
}

type RegisterReq struct {
	Name     string
	Surname  string
	Email    string
	Password string
	Role     string
}
