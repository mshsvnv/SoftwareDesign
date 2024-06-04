package service

import (
	"context"
	"fmt"
	"src/internal/dto"
	"src/internal/model"
	repo "src/internal/repository"
	"src/pkg/logging"
	"src/pkg/utils"
	"time"
)

type IOrderService interface {
	CreateOrder(ctx context.Context, req *dto.PlaceOrderReq) (*model.Order, error)
	GetMyOrders(ctx context.Context, userID int) ([]*model.Order, error)
	GetAllInProgressOrders(ctx context.Context) ([]*model.Order, error)
	GetOrderByID(ctx context.Context, orderID int) (*model.Order, error)
	UpdateOrder(ctx context.Context, orderID int, userID int) (*model.Order, error)
}

type OrderService struct {
	logger     logging.Interface
	repo       repo.IOrderRepository
	repoCart   repo.ICartRepository
	repoRacket repo.IRacketRepository
}

func NewOrderService(logger logging.Interface, repo repo.IOrderRepository, repoCart repo.ICartRepository, repoRacket repo.IRacketRepository) *OrderService {
	return &OrderService{
		logger:     logger,
		repo:       repo,
		repoCart:   repoCart,
		repoRacket: repoRacket,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, req *dto.PlaceOrderReq) error {

	s.logger.Infof("create order user %d", req.UserID)
	order := &model.Order{}
	utils.Copy(&order, &req)

	order.CreationDate = time.Now()

	cart, err := s.repoCart.GetCartByID(ctx, order.UserID)

	if err != nil {
		s.logger.Errorf("get cart fail, error %s", err.Error())
		return fmt.Errorf("get cart fail, error %s", err)
	}

	for _, line := range cart.Lines {

		racket, err := s.repoRacket.GetRacketByID(ctx, line.RacketID)

		if err != nil {
			s.logger.Errorf("get racket by id fail, error %s", err.Error())
			return fmt.Errorf("get racket by id fail, error %s", err)
		}

		if line.Quantity <= racket.Quantity {
			racket.Quantity -= line.Quantity
		} else {
			s.logger.Errorf("not avaliable amount or rackets, error")
			return fmt.Errorf("not avaliable amount or rackets, error")
		}

		err = s.repoRacket.Update(ctx, racket)

		if err != nil {
			s.logger.Errorf("update racket after order creation fail, error %s", err.Error())
			return fmt.Errorf("update racket after order creation fail, error %s", err)
		}
	}

	order.Status = model.OrderStatusInProgress
	order.TotalPrice = cart.TotalPrice

	for _, line := range cart.Lines {
		order.Lines = append(order.Lines,
			&model.OrderLine{
				RacketID: line.RacketID,
				Quantity: line.Quantity,
			})
	}

	err = s.repo.Create(ctx, order)

	if err != nil {
		s.logger.Errorf("create order fail, error %s", err.Error())
		return fmt.Errorf("create order fail, error %s", err)
	}

	err = s.repoCart.Remove(ctx, req.UserID)

	if err != nil {
		s.logger.Errorf("remove racket fail, error %s", err.Error())
		return fmt.Errorf("remove racket fail, error %s", err)
	}

	return nil
}

func (s *OrderService) GetMyOrders(ctx context.Context, userID int) ([]*model.Order, error) {

	s.logger.Infof("get user's orders %d", userID)
	orders, err := s.repo.GetMyOrders(ctx, userID)

	if err != nil {
		s.logger.Errorf("get my orders fail, error %s", err.Error())
		return nil, fmt.Errorf("get my orders fail, error %s", err)
	}

	return orders, nil
}

func (s *OrderService) GetAllInProgressOrders(ctx context.Context) ([]*model.Order, error) {

	s.logger.Infof("get all progressive orders")
	orders, err := s.repo.GetAllInProgressOrders(ctx)

	if err != nil {
		s.logger.Errorf("get all in progress fail, error %s", err.Error())
		return nil, fmt.Errorf("get all in progress fail, error %s", err)
	}

	return orders, nil
}

func (s *OrderService) GetAllOrders(ctx context.Context) ([]*model.Order, error) {

	s.logger.Infof("get all progressive orders")
	orders, err := s.repo.GetAllOrders(ctx)

	if err != nil {
		s.logger.Errorf("get all in progress fail, error %s", err.Error())
		return nil, fmt.Errorf("get all in progress fail, error %s", err)
	}

	return orders, nil
}

func (s *OrderService) GetOrderByID(ctx context.Context, orderID int) (*model.Order, error) {

	s.logger.Infof("get order by id %d", orderID)
	order, err := s.repo.GetOrderByID(ctx, orderID)

	if err != nil {
		s.logger.Errorf("get order by id fail, error %s", err.Error())
		return nil, fmt.Errorf("get order by id fail, error %s", err)
	}

	return order, nil
}

func (s *OrderService) UpdateOrder(ctx context.Context, req *dto.UpdateOrder) (*model.Order, error) {

	s.logger.Infof("update order order %d", req.OrderID)
	order, err := s.repo.GetOrderByID(ctx, req.OrderID)

	if err != nil {
		s.logger.Errorf("get by id fail, error %s", err.Error())
		return nil, fmt.Errorf("get by id fail, error %s", err)
	}

	order.Status = req.Status

	err = s.repo.Update(ctx, order)

	if err != nil {
		s.logger.Errorf("update order fail, error %s", err.Error())
		return nil, fmt.Errorf("update order fail, error %s", err)
	}

	return order, nil
}
