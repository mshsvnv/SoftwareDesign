package dto

import "src/internal/model"

type UpdateOrder struct {
	OrderID int
	Status  model.OrderStatus
}

type PlaceOrderReq struct {
	UserID    int              `json:"user_id"`
	OrderInfo *model.OrderInfo `json:"order_info"`
}
