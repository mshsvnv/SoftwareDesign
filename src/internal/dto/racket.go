package dto

type CreateRacketReq struct {
	Brand         string  `json:"brand"`
	SupplierEmail string  `json:"supplier_email"`
	Weight        float32 `json:"weight"`
	Balance       float32 `json:"balance"`
	HeadSize      float32 `json:"head_size"`
	Avaliable     bool    `json:"avaliable"`
	Quantity      int     `json:"quantity"`
	Price         float32 `json:"price"`
}

type UpdateRacketReq struct {
	ID       int
	Quantity int `json:"quantity"`
}
