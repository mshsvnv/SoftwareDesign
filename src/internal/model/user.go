package model

type UserRole string

const (
	UserRoleCustomer = "Customer"
	UserRoleAdmin    = "Admin"
	UserRoleSeller   = "Seller"
)

type User struct {
	ID       int      `json:"id"`
	Name     string   `json:"name"`
	Surname  string   `json:"surname"`
	Email    string   `json:"email"`
	Password string   `json:"password"`
	Role     UserRole `json:"role"`
}

