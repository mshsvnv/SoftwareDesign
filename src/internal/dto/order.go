package dto

import "src/internal/model"

type UpdateOrder struct {
	OrderID int
	Status  model.OrderStatus
}

type PlaceOrderReq struct {
	UserID    int
	OrderInfo *model.OrderInfo `json:"orderinfo"`
}
