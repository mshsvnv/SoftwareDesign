package model

type CartLine struct {
	RacketID int
	Quantity int
}

type Cart struct {
	UserID     int
	TotalPrice float32
	Quantity   int
	Lines      []*CartLine
}
