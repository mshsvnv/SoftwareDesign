package model

type Cart struct {
	ID     string
	UserID string
	User   *User
	Lines  []*CartLine
}

type CartLine struct {
	RacketID string
	Racket   *Racket
	Quantity uint
}
