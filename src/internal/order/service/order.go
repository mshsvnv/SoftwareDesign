package service

import (
	"context"
	"errors"

	"src/internal/order/dto"
	"src/internal/order/model"
	"src/internal/order/repository"
	"src/pkg/utils"
)

type IOrderService interface {
	PlaceOrder(ctx context.Context, req *dto.PlaceOrderReq) (*model.Order, error)
	GetOrderByID(ctx context.Context, id string) (*model.Order, error)
	GetMyOrders(ctx context.Context, req *dto.ListOrderReq) ([]*model.Order, error)
	CancelOrder(ctx context.Context, orderID, userID string) (*model.Order, error)
}

type OrderService struct {
	repo        repository.IOrderRepository
	racketRepo  repository.IOrderRacketRepository
	paymentRepo repository.IPaymentRepository
}

func NewOrderService(
	repo repository.IOrderRepository,
	racketRepo repository.IOrderRacketRepository,
	paymentRepo repository.IPaymentRepository,
) *OrderService {
	return &OrderService{
		repo:        repo,
		racketRepo:  racketRepo,
		paymentRepo: paymentRepo,
	}
}

func (s *OrderService) PlaceOrder(ctx context.Context, req *dto.PlaceOrderReq) (*model.Order, error) {

	var Rackets []*model.OrderRacket
	utils.Copy(&req.Rackets, &Rackets)

	racketMap := make(map[string]*model.Racket)

	totalPrice := 0.

	for _, Racket := range Rackets {

		racket, err := s.racketRepo.GetRacketByID(ctx, Racket.RacketID)

		if err != nil {
			return nil, err
		}

		Racket.Price = racket.Price * float64(Racket.Quantity)
		totalPrice += Racket.Price
		racketMap[Racket.RacketID] = racket
	}

	order, err := s.repo.Create(ctx, req.UserID, Rackets)

	if err != nil {
		return nil, err
	}

	payment := &model.Payment{
		OrderID:    order.ID,
		Status:     model.PaymentStatusInProgress,
		TotalPrice: totalPrice,
	}

	err = s.paymentRepo.Create(ctx, payment)

	if err != nil {
		return nil, err
	}

	for _, Racket := range order.Rackets {
		Racket.Racket = racketMap[Racket.RacketID]
	}

	return order, nil
}

func (s *OrderService) GetOrderByID(ctx context.Context, id string) (*model.Order, error) {

	order, err := s.repo.GetOrderByID(ctx, id, true)

	if err != nil {
		return nil, err
	}

	return order, nil
}

func (s *OrderService) GetMyOrders(ctx context.Context, req *dto.ListOrderReq) ([]*model.Order, error) {

	orders, err := s.repo.GetMyOrders(ctx, req)

	if err != nil {
		return nil, err
	}

	return orders, err
}

func (s *OrderService) CancelOrder(ctx context.Context, orderID, userID string) (*model.Order, error) {

	order, err := s.repo.GetOrderByID(ctx, orderID, false)

	if err != nil {
		return nil, err
	}

	if userID != order.UserID {
		return nil, errors.New("permission denied")
	}

	if order.Status == model.OrderStatusDone || order.Status == model.OrderStatusCancelled {
		return nil, errors.New("invalid order status")
	}

	order.Status = model.OrderStatusCancelled

	err = s.repo.Update(ctx, order)

	if err != nil {
		return nil, err
	}

	payment, err := s.paymentRepo.GetPaymentByID(ctx, orderID)

	if err != nil {
		return nil, err
	}

	payment.Status = model.PaymentStatusCancelled

	err = s.paymentRepo.Update(ctx, payment)

	if err != nil {
		return nil, err
	}

	return order, nil
}
