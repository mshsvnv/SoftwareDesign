package model

type CartLine struct {
	RacketID int     `json:"id"`
	Quantity int     `json:"quantity"`
	Price    float32 `json:"price"`
}

type Cart struct {
	UserID     int         `json:"user_id"`
	TotalPrice float32     `json:"total_price"`
	Quantity   int         `json:"quantity"`
	Lines      []*CartLine `json:"lines"`
}
