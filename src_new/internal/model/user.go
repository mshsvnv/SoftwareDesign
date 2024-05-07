package model

type UserRole string

const (
	UserRoleCustomer = "Customer"
	UserRoleSeller   = "Seller"
)

type User struct {
	ID       int
	Name     string
	Surname  string
	Email    string
	Password string
	Role     UserRole
}
