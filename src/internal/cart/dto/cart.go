package dto

type Cart struct {
	ID      string
	User    *User
	Rackets []*CartRacketReq
}

type CartRacket struct {
	Racket   *Racket
	Quantity uint
}

type CartRacketReq struct {
	RacketID string
	Quantity uint
}

type AddRacketReq struct {
	UserID string
	Racket *CartRacketReq
}

type RemoveRacketReq struct {
	UserID   string
	RacketID string
}
