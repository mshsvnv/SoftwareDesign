package model

type OrderLine struct {
	OrderID  string
	RacketID string
	Racket   *Racket
	Quantity uint
	Price    float64
}
