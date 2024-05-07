package dto

type Order struct {
	ID         string
	Rackets    []*OrderRacket
	Status     string
	TotalPrice float64
}

type OrderRacket struct {
	Racket   Racket
	Quantity uint
	Price    float64
}

type PlaceOrderReq struct {
	UserID  string
	Rackets []PlaceOrderRacketReq
}

type PlaceOrderRacketReq struct {
	RacketID string
	Quantity uint
}

type ListOrderReq struct {
	UserID string
	Status string
}
