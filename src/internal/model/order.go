package model

import "time"

type OrderStatus string

const (
	OrderStatusInProgress = "InProgress"
	OrderStatusDone       = "Done"
	OrderStatusCancelled  = "Cancelled"
)

type OrderLine struct {
	RacketID int
	Quantity int
}

type OrderInfo struct {
	DeliveryDate  time.Time
	Address       string
	RecepientName string
}

type Order struct {
	ID           int
	UserID       int
	CreationDate time.Time
	OrderInfo    *OrderInfo
	Status       OrderStatus
	Lines        []*OrderLine
	TotalPrice   float32
}
