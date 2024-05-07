package dto

type CreateSupplierReq struct {
	Name  string
	Phone string
	Town  string
	Email string
}

type UpdateSupplierReq struct {
	ID       int
	Name  string
	Phone string
	Town  string
	Email string
}
