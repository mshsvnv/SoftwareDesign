package model

type Racket struct {
	ID            int     `json:"id"`
	SupplierEmail string  `json:"supplier_email"`
	Brand         string  `json:"brand"`
	Weight        float32 `json:"weight"`
	Balance       float32 `json:"balance"`
	HeadSize      float32 `json:"headsize"`
	Avaliable     bool    `json:"avaliable"`
	Quantity      int     `json:"quantity"`
	Price         int     `json:"price"`
}
