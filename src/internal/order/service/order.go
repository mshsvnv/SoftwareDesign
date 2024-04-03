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
	repo       repository.IOrderRepository
	racketRepo repository.IRacketRepository
}

func NewOrderService(
	repo repository.IOrderRepository,
	racketRepo repository.IRacketRepository,
) *OrderService {
	return &OrderService{
		repo:       repo,
		racketRepo: racketRepo,
	}
}

func (s *OrderService) PlaceOrder(ctx context.Context, req *dto.PlaceOrderReq) (*model.Order, error) {

	var lines []*model.OrderLine
	utils.Copy(&req.Lines, &lines)

	racketMap := make(map[string]*model.Racket)

	for _, line := range lines {

		racket, err := s.racketRepo.GetRacketByID(ctx, line.RacketID)

		if err != nil {
			return nil, err
		}

		line.Price = racket.Price * float64(line.Quantity)
		racketMap[line.RacketID] = racket
	}

	order, err := s.repo.CreateOrder(ctx, req.UserID, lines)

	if err != nil {
		return nil, err
	}

	for _, line := range order.Lines {
		line.Racket = racketMap[line.RacketID]
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

	err = s.repo.UpdateOrder(ctx, order)

	if err != nil {
		return nil, err
	}

	return order, nil
}
