package repository

import (
	"context"
	"src/internal/order/dto"
	"src/internal/order/model"
)

//go:generate mockery --name=IOrderRepository
type IOrderRepository interface {
	CreateOrder(ctx context.Context, userID string, lines []*model.OrderLine) (*model.Order, error)
	UpdateOrder(ctx context.Context, order *model.Order) error
	GetOrderByID(ctx context.Context, id string, preload bool) (*model.Order, error)
	GetMyOrders(ctx context.Context, req *dto.ListOrderReq) ([]*model.Order, error)
}
