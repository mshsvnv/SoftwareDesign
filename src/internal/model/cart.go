package model

type CartLine struct {
	RacketID int
	Quantity int
	Price    float32
}

type Cart struct {
	UserID     int         `json:"user_id"`
	TotalPrice float32     `json:"total_price"`
	Quantity   int         `json:"quantity"`
	Lines      []*CartLine `json:"lines"`
}
