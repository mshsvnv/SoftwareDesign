package model

type Cart struct {
	ID      string
	UserID  string
	User    *User
	Rackets []*CartRacket
}

type CartRacket struct {
	RacketID string
	Racket   *Racket
	Quantity uint
}
