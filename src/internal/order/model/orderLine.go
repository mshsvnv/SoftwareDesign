package model

type OrderRacket struct {
	OrderID  string
	RacketID string
	Racket   *Racket
	Quantity uint
	Price    float64
}
