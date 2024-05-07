package repository

import (
	"context"

	"src/internal/order/model"
)

//go:generate mockery --name=IPaymentRepository
type IPaymentRepository interface {
	Create(ctx context.Context, Payment *model.Payment) error
	Update(ctx context.Context, Payment *model.Payment) error
	GetPaymentByID(ctx context.Context, orderID string) (*model.Payment, error)
}
