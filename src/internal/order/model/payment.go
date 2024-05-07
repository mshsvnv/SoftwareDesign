package model

import "time"

type PaymentStatus string

const (
	PaymentStatusNew        PaymentStatus = "new"
	PaymentStatusInProgress PaymentStatus = "in_progress"
	PaymentStatusDone       PaymentStatus = "done"
	PaymentStatusCancelled  PaymentStatus = "cancelled"
)

type Payment struct {
	OrderID    string
	Status     PaymentStatus
	Date       time.Time
	TotalPrice float64
}
