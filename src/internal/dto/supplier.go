package dto

type CreateSupplierReq struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Town  string `json:"town"`
	Email string `json:"email"`
}
