package model

import "time"

type OrderStatus string

const (
	OrderStatusNew        OrderStatus = "new"
	OrderStatusInProgress OrderStatus = "in_progress"
	OrderStatusDone       OrderStatus = "done"
	OrderStatusCancelled  OrderStatus = "cancelled"
)

type Order struct {
	ID            string
	UserID        string
	CreationDate  time.Time
	DeliveryDate  time.Time
	Address       string
	RecepientName string
	Lines         []*OrderLine
	Status        OrderStatus
	TotalPrice    float64
}
