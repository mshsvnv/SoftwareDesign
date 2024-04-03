package dto

type Cart struct {
	ID    string
	User  *User
	Lines []*CartLineReq
}

type CartLine struct {
	Racket   *Racket
	Quantity uint
}

type CartLineReq struct {
	RacketID string
	Quantity uint
}

type AddRacketReq struct {
	UserID string
	Line   *CartLineReq
}

type RemoveRacketReq struct {
	UserID   string
	RacketID string
}
