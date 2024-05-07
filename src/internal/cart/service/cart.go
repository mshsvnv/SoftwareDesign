package service

import (
	"context"

	"src/internal/cart/dto"
	"src/internal/cart/model"
	"src/internal/cart/repository"
)

type ICartService interface {
	AddRacket(ctx context.Context, req *dto.AddRacketReq) (*model.Cart, error)
	RemoveRacket(ctx context.Context, req *dto.RemoveRacketReq) (*model.Cart, error)
	GetCartByUserID(ctx context.Context, userID string) (*model.Cart, error)
}

type CartService struct {
	repo repository.ICartRepository
}

func NewCartService(
	repo repository.ICartRepository,
) *CartService {
	return &CartService{
		repo: repo,
	}
}

func (p *CartService) AddRacket(ctx context.Context, req *dto.AddRacketReq) (*model.Cart, error) {

	cart, err := p.repo.GetCartByUserID(ctx, req.UserID)

	if err != nil {

		cart = &model.Cart{
			UserID: req.UserID,
			Rackets: []*model.CartRacket{{
				RacketID: req.Racket.RacketID,
				Quantity: req.Racket.Quantity,
			}},
		}

		err = p.repo.Create(ctx, cart)

		if err != nil {
			return nil, err
		}

		return cart, nil
	}

	for _, Racket := range cart.Rackets {

		if Racket.RacketID == req.Racket.RacketID {
			return cart, nil
		}
	}

	cart.Rackets = append(cart.Rackets, &model.CartRacket{
		RacketID: req.Racket.RacketID,
		Quantity: req.Racket.Quantity,
	})

	err = p.repo.Update(ctx, cart)

	if err != nil {
		return nil, err
	}

	return cart, nil
}

func (p *CartService) GetCartByUserID(ctx context.Context, userID string) (*model.Cart, error) {

	cart, err := p.repo.GetCartByUserID(ctx, userID)

	if err != nil {
		cart = &model.Cart{
			UserID: userID,
		}

		err = p.repo.Create(ctx, cart)

		if err != nil {
			return nil, err
		}

		return cart, nil
	}

	return cart, nil
}

func (p *CartService) RemoveRacket(ctx context.Context, req *dto.RemoveRacketReq) (*model.Cart, error) {

	cart, err := p.repo.GetCartByUserID(ctx, req.UserID)

	if err != nil {

		cart = &model.Cart{
			UserID: req.UserID,
		}

		err = p.repo.Create(ctx, cart)

		if err != nil {
			return nil, err
		}

		return cart, nil
	}

	for i, Racket := range cart.Rackets {

		if Racket.RacketID == req.RacketID {
			cart.Rackets = append(cart.Rackets[:i], cart.Rackets[i+1:]...)
			break
		}
	}

	err = p.repo.Update(ctx, cart)

	if err != nil {
		return nil, err
	}

	return cart, nil
}
