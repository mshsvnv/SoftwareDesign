package dto

type CreateRacketReq struct {
	Brand         string
	SupplierEmail string
	Weight        float32
	Balance       float32
	HeadSize      float32
	Quantity      int
	Price         float32
}

type UpdateRacketReq struct {
	ID            int
	Brand         string
	SupplierEmail string
	Weight        float32
	Balance       float32
	HeadSize      float32
	Quantity      int
	Price         float32
}
