package model

type UserRole string

const (
	UserRoleCustomer = "Customer"
	UserRoleAdmin    = "Admin"
)

type User struct {
	ID           int
	Name         string
	Surname      string
	Email        string
	Password     string
	Role         UserRole
	Subscription bool
}
