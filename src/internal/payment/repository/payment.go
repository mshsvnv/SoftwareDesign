package repository

import (
	"context"

	"src/internal/payment/dto"
	"src/internal/payment/model"
)

//go:generate mockery --name=IPaymentRepository
type IPaymentRepository interface {
	CreatePayment(ctx context.Context, Payment *model.Payment) error
	UpdatePayment(ctx context.Context, Payment *model.Payment) error
	GetPaymentByOrderID(ctx context.Context, orderID string) (*model.Payment, error)
	GetMyPayments(ctx context.Context, req *dto.ListPaymentReq) ([]*model.Payment, error)
}
