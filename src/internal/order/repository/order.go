package repository

import (
	"context"
	"src/internal/order/dto"
	"src/internal/order/model"
)

//go:generate mockery --name=IOrderRepository
type IOrderRepository interface {
	Create(ctx context.Context, userID string, Rackets []*model.OrderRacket) (*model.Order, error)
	Update(ctx context.Context, order *model.Order) error
	Delete(ctx context.Context, order *model.Order) error
	DeleteOrdersByRacketID(ctx context.Context, id string) error
	GetOrderByID(ctx context.Context, id string, preload bool) (*model.Order, error)
	GetMyOrders(ctx context.Context, req *dto.ListOrderReq) ([]*model.Order, error)
}
