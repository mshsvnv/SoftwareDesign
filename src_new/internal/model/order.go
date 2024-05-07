package model

import "time"

type OrderStatus string

const (
	// OrderStatusNew        = "New"
	OrderStatusInProgress = "In progress"
	OrderStatusCanceled   = "Canceled"
	OrderStatusDone       = "Done"
)

type OrderLine struct {
	OrderID  int
	RacketID int
	Quantity int
}

type OrderInfo struct {
	DeliveryDate  time.Time
	Address       string
	RecepientName string
}

type Order struct {
	ID         int
	UserID     int
	OrderInfo  *OrderInfo
	Status     OrderStatus
	Lines      []*OrderLine
	TotalPrice float32
}
