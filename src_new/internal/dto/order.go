package dto

import "src_new/internal/model"

// type Order struct {
// 	ID         string
// 	Code       string
// 	Lines      []*OrderLine
// 	TotalPrice float32
// 	Status     string
// }

// type OrderLine struct {
// 	RacketID int
// 	Quantity int
// 	Price    float64
// }

type PlaceOrderReq struct {
	UserID    int
	OrderInfo *model.OrderInfo
	// Status    model.OrderStatus
	// Lines     []*PlaceOrderLineReq
}

// type PlaceOrderLineReq struct {
// 	RacketID int
// 	Quantity int
// }
