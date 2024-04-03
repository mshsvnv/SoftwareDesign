package service

import (
	"context"
	"errors"

	"src/internal/payment/dto"
	"src/internal/payment/model"
	"src/internal/payment/repository"
)

type IPaymentService interface {
	GetPaymentByOrderID(ctx context.Context, orderID string) (*model.Payment, error)
	GetMyPayments(ctx context.Context, req *dto.ListPaymentReq) ([]*model.Payment, error)
	CancelPayment(ctx context.Context, paymentID, userID string) (*model.Payment, error)
}

type PaymentService struct {
	repo repository.IPaymentRepository
}

func NewPaymentService(
	repo repository.IPaymentRepository,
) *PaymentService {
	return &PaymentService{
		repo: repo}
}

func (p *PaymentService) GetPaymentByOrderID(ctx context.Context, OrderID string) (*model.Payment, error) {

	payment, err := p.repo.GetPaymentByOrderID(ctx, OrderID)

	if err != nil {
		payment = &model.Payment{
			OrderID: OrderID,
		}

		err = p.repo.CreatePayment(ctx, payment)

		if err != nil {
			return nil, err
		}

		return payment, nil
	}

	return payment, nil
}

func (p *PaymentService) GetMyPayments(ctx context.Context, req *dto.ListPaymentReq) ([]*model.Payment, error) {

	payments, err := p.repo.GetMyPayments(ctx, req)

	if err != nil {
		return nil, err
	}

	return payments, err
}

func (s *PaymentService) CancelPayment(ctx context.Context, paymentID, orderID string) (*model.Payment, error) {

	payment, err := s.repo.GetPaymentByOrderID(ctx, orderID)

	if err != nil {
		return nil, err
	}

	if orderID != payment.OrderID {
		return nil, errors.New("permission denied")
	}

	if payment.Status == model.PaymentStatusDone || payment.Status == model.PaymentStatusCancelled {
		return nil, errors.New("invalid Payment status")
	}

	payment.Status = model.PaymentStatusCancelled

	err = s.repo.UpdatePayment(ctx, payment)

	if err != nil {
		return nil, err
	}

	return payment, nil
}
