package model

import "time"

type OrderStatus string

const (
	OrderStatusInProgress = "InProgress"
	OrderStatusDone       = "Done"
	OrderStatusCancelled  = "Cancelled"
)

type OrderLine struct {
	RacketID int `json:"racket_id"`
	Quantity int `json:"quantity"`
}

type OrderInfo struct {
	DeliveryDate  time.Time `json:"delivery_date"`
	Address       string    `json:"address"`
	RecepientName string    `json:"recepient_name"`
}

type Order struct {
	ID           int          `json:"id"`
	UserID       int          `json:"user_id"`
	CreationDate time.Time    `json:"creation_date"`
	OrderInfo    *OrderInfo   `json:"order_info"`
	Status       OrderStatus  `json:"status"`
	Lines        []*OrderLine `json:"lines"`
	TotalPrice   float32      `json:"total_price"`
}
