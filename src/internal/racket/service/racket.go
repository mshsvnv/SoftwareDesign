package service

import (
	"context"
	"errors"

	"src/internal/racket/dto"
	"src/internal/racket/model"
	"src/internal/racket/repository"

	cart "src/internal/cart/repository"
	order "src/internal/order/repository"

	"src/pkg/utils"
)

type IRacketService interface {
	ListRackets(ctx context.Context, req *dto.ListRacketReq) ([]*model.Racket, error)
	GetRacketByID(ctx context.Context, id string) (*model.Racket, error)
	Create(ctx context.Context, req *dto.CreateRacketReq) (*model.Racket, error)
	Update(ctx context.Context, id string, req *dto.UpdateRacketReq) (*model.Racket, error)
	Delete(ctx context.Context, id string) error
}

type RacketService struct {
	repo      repository.IRacketRepository
	repoCart  cart.ICartRepository
	repoOrder order.IOrderRepository
}

func NewRacketService(
	repo repository.IRacketRepository,
	repoCart cart.ICartRepository,
	repoOrder order.IOrderRepository,
) *RacketService {
	return &RacketService{
		repo:      repo,
		repoCart:  repoCart,
		repoOrder: repoOrder,
	}
}

func (p *RacketService) GetRacketByID(ctx context.Context, id string) (*model.Racket, error) {

	racket, err := p.repo.GetRacketByID(ctx, id)

	if err != nil {
		return nil, err
	}

	return racket, nil
}

func (p *RacketService) ListRackets(ctx context.Context, req *dto.ListRacketReq) ([]*model.Racket, error) {

	rackets, err := p.repo.ListRackets(ctx, req)

	if err != nil {
		return nil, err
	}

	return rackets, nil
}

func (p *RacketService) Create(ctx context.Context, req *dto.CreateRacketReq) (*model.Racket, error) {

	var racket model.Racket
	utils.Copy(req, &racket)

	err := p.repo.Create(ctx, &racket)

	if err != nil {
		return nil, err
	}

	return &racket, nil
}

func (p *RacketService) Update(ctx context.Context, id string, req *dto.UpdateRacketReq) (*model.Racket, error) {

	racket, err := p.repo.GetRacketByID(ctx, id)

	if err != nil {
		return nil, err
	}

	utils.Copy(req, racket)
	err = p.repo.Update(ctx, racket)

	if err != nil {
		return nil, err
	}

	return racket, nil
}

func (p *RacketService) Delete(ctx context.Context, id string) error {

	_, err := p.repo.GetRacketByID(ctx, id)

	if err != nil {
		return errors.New("racket not found")
	}

	err = p.repoCart.DeleteByRacketID(ctx, id)

	if err != nil {
		return errors.New("failed to delete associated carts")
	}

	err = p.repoOrder.DeleteOrdersByRacketID(ctx, id)

	if err != nil {
		return errors.New("failed to delete associated orders")
	}

	err = p.repo.Delete(ctx, id)

	if err != nil {
		return errors.New("failed to delete rocket")
	}

	return nil
}
