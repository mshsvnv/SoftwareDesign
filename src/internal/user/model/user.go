package model

type UserRole string

const (
	UserRoleAdmin    UserRole = "admin"
	UserRoleCustomer UserRole = "customer"
	UserRoleSeller   UserRole = "seller"
)

type User struct {
	ID       string
	Email    string
	Password string
	Role     UserRole
}
