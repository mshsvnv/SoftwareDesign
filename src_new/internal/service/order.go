package service

import (
	"context"
	"fmt"
	"src_new/internal/dto"
	"src_new/internal/model"
	repo "src_new/internal/repository"
	"src_new/pkg/utils"
)

type IOrderService interface {
	CreateOrder(ctx context.Context, req *dto.PlaceOrderReq) (*model.Order, error)
	GetMyOrders(ctx context.Context, userID int) ([]*model.Order, error)
	GetOrderByID(ctx context.Context, orderID int) (*model.Order, error)
	CancelOrder(ctx context.Context, orderID int, userID int) (*model.Order, error)
}

type OrderService struct {
	repo       repo.IOrderRepository
	repoRacket repo.IRacketRepository
}

func NewOrderService(repo repo.IOrderRepository, repoRacket repo.IRacketRepository) *OrderService {
	return &OrderService{
		repo:       repo,
		repoRacket: repoRacket,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, req *dto.PlaceOrderReq) error {

	var order *model.Order
	utils.Copy(&order, &req)

	for _, line := range order.Lines {

		racket, err := s.repoRacket.GetRacketByID(ctx, line.RacketID)

		if err != nil {
			return fmt.Errorf("CreateOrder.GetRacketByID fail, %s", err)
		}

		order.TotalPrice += float32(racket.Price) * float32(line.Quantity)
	}

	err := s.repo.Create(ctx, order)

	if err != nil {
		return fmt.Errorf("CreateOrder.Create fail, %s", err)
	}

	return nil
}

func (s *OrderService) GetMyOrders(ctx context.Context, userID int) ([]*model.Order, error) {

	orders, err := s.repo.GetMyOrders(ctx, userID)

	if err != nil {
		return nil, fmt.Errorf("GetMyOrders.GetMyOrders fail, %s", err)
	}

	return orders, nil
}

func (s *OrderService) GetOrderByID(ctx context.Context, orderID int) (*model.Order, error) {

	order, err := s.repo.GetOrderByID(ctx, orderID)

	if err != nil {
		return nil, fmt.Errorf("GetOrderByID.GetOrderByID fail, %s", err)
	}

	return order, nil
}

func (s *OrderService) CancelOrder(ctx context.Context, orderID int, userID int) (*model.Order, error) {

	order, err := s.repo.GetOrderByID(ctx, orderID)

	if err != nil {
		return nil, fmt.Errorf("CancelOrder.GetOrderByID fail, %s", err)
	}

	if order.Status == model.OrderStatusInProgress {
		order.Status = model.OrderStatusCanceled
	}

	err = s.repo.Update(ctx, order)

	if err != nil {
		return nil, fmt.Errorf("CancelOrder.Update fail, %s", err)
	}

	return order, nil
}
