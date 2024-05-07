package repository

import (
	"context"
	"src_new/internal/model"
)

//go:generate mockery --name=IOrderRepository
type IOrderRepository interface {
	Create(ctx context.Context, order *model.Order) error
	Update(ctx context.Context, order *model.Order) error
	Remove(ctx context.Context, orderID int) error
	GetMyOrders(ctx context.Context, userID int) ([]*model.Order, error)
	GetOrderByID(ctx context.Context, orderID int) (*model.Order, error)
}
