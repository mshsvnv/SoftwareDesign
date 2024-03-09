package model

type UserRole string

const (
	UserRoleAdmin    UserRole = "admin"
	UserRoleCustomer UserRole = "customer"
)

type User struct {
	ID       string
	Email    string
	Password string
	Role     UserRole
}

// func (user *User) BeforeCreate(tx *gorm.DB) error {

// 	user.ID =
// }
