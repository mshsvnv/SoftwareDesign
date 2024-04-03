package dto

type Order struct {
	ID         string
	Lines      []*OrderLine
	Status     string
	TotalPrice float64
}

type OrderLine struct {
	Racket   Racket
	Quantity uint
	Price    float64
}

type PlaceOrderReq struct {
	UserID string
	Lines  []PlaceOrderLineReq
}

type PlaceOrderLineReq struct {
	RacketID string
	Quantity uint
}

type ListOrderReq struct {
	UserID string
	Status string
}
