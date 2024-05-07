package service

import (
	"context"
	"fmt"
	"src_new/internal/dto"
	"src_new/internal/model"
	repo "src_new/internal/repository"
)

type ICartService interface {
	AddRacket(ctx context.Context, req *dto.AddRacketCartReq) (*model.Cart, error)
	RemoveRacket(ctx context.Context, req *dto.RemoveRacketCartReq) (*model.Cart, error)
	UpdateRacket(ctx context.Context, req *dto.UpdateRacketCartReq) (*model.Cart, error)
	GetCartByID(ctx context.Context, userID int) (*model.Cart, error)
	// CleanCart(ctx context.Context, userID int) (*model.Cart, error)
}

type CartService struct {
	repo       repo.ICartRepository
	repoRacket repo.IRacketRepository
}

func NewCartService(repo repo.ICartRepository, repoRacket repo.IRacketRepository) *CartService {
	return &CartService{
		repo:       repo,
		repoRacket: repoRacket,
	}
}

func (s *CartService) AddRacket(ctx context.Context, req *dto.AddRacketCartReq) (*model.Cart, error) {

	cart, err := s.repo.GetCartByID(ctx, req.UserID)

	if err != nil {

		racket, err := s.repoRacket.GetRacketByID(ctx, req.RacketID)

		if err != nil {
			return nil, fmt.Errorf("AddRacket.GetRacketByID fail, error %s", err)
		}

		cart = &model.Cart{
			UserID:     req.UserID,
			TotalPrice: float32(racket.Price) * float32(req.Quantity),
			Quantity:   req.Quantity,
			Lines: []*model.CartLine{{
				RacketID: req.RacketID,
				Quantity: req.Quantity,
			}},
		}

		racket.Quantity -= req.Quantity
		err = s.repoRacket.Update(ctx, racket)

		if err != nil {
			return nil, fmt.Errorf("AddRacket.UpdateRacket fail, error %s", err)
		}

		err = s.repo.Create(ctx, cart)

		if err != nil {
			return nil, fmt.Errorf("AddRacket.Create fail, userID: %d, error %s", req.UserID, err)
		}

		return cart, nil
	}

	for _, line := range cart.Lines {

		if line.RacketID == req.RacketID {
			return cart, nil
		}
	}

	cart.Lines = append(cart.Lines,
		&model.CartLine{
			RacketID: req.RacketID,
			Quantity: req.Quantity,
		})

	racket, err := s.repoRacket.GetRacketByID(ctx, req.RacketID)

	if err != nil {
		return nil, fmt.Errorf("AddRacket.GetRacketByID fail, error %s", err)
	}

	racket.Quantity -= req.Quantity
	err = s.repoRacket.Update(ctx, racket)

	if err != nil {
		return nil, fmt.Errorf("AddRacket.UpdateRacket fail, error %s", err)
	}

	cart.Quantity += req.Quantity
	cart.TotalPrice += float32(racket.Price) * float32(req.Quantity)
	err = s.repo.Update(ctx, cart)

	if err != nil {
		return nil, fmt.Errorf("AddRacket.Update fail, userID: %d, error %s", req.UserID, err)
	}

	return cart, nil
}

func (s *CartService) RemoveRacket(ctx context.Context, req *dto.RemoveRacketCartReq) (*model.Cart, error) {

	cart, err := s.repo.GetCartByID(ctx, req.UserID)

	if err != nil {

		cart = &model.Cart{
			UserID: req.UserID,
		}

		err = s.repo.Create(ctx, cart)

		if err != nil {
			return nil, fmt.Errorf("UpdateRacket.Create fail, userID: %d, error %s", req.UserID, err)
		}

		return cart, nil
	}

	for i := 0; i < len(cart.Lines); i++ {

		if cart.Lines[i].RacketID == req.RacketID {

			racket, err := s.repoRacket.GetRacketByID(ctx, req.RacketID)

			if err != nil {
				return nil, fmt.Errorf("AddRacket.GetRacketByID fail, error %s", err)
			}

			cart.Quantity -= cart.Lines[i].Quantity
			cart.TotalPrice -= float32(racket.Price) * float32(cart.Lines[i].Quantity)

			racket.Quantity += cart.Lines[i].Quantity

			err = s.repoRacket.Update(ctx, racket)

			if err != nil {
				return nil, fmt.Errorf("AddRacket.Update fail, error %s", err)
			}

			cart.Lines = append(cart.Lines[:i], cart.Lines[i+1:]...)
			break
		}
	}

	err = s.repo.Update(ctx, cart)

	if err != nil {
		return nil, fmt.Errorf("RemoveRacket.Update fail, userID: %d, error %s", req.UserID, err)
	}

	return cart, nil
}

func (s *CartService) UpdateRacket(ctx context.Context, req *dto.UpdateRacketCartReq) (*model.Cart, error) {

	cart, err := s.repo.GetCartByID(ctx, req.UserID)

	if err != nil {

		cart := &model.Cart{
			UserID: req.UserID,
		}

		err = s.repo.Create(ctx, cart)

		if err != nil {
			return nil, fmt.Errorf("UpdateRacket.Create fail, userID: %d, error %s", req.UserID, err)
		}

		return cart, nil
	}

	for _, line := range cart.Lines {

		if line.RacketID == req.RacketID {
			line.Quantity = req.Quantity

			racket, err := s.repoRacket.GetRacketByID(ctx, req.RacketID)

			if err != nil {
				return nil, fmt.Errorf("UpdateRacket.GetRacketByID fail, error %s", err)
			}

			cart.TotalPrice -= float32(req.Quantity) * float32(racket.Price)
			cart.Quantity -= req.Quantity

			racket.Quantity += req.Quantity

			err = s.repoRacket.Update(ctx, racket)

			if err != nil {
				return nil, fmt.Errorf("UpdateRacket.Update fail, error %s", err)
			}
		}
	}

	err = s.repo.Update(ctx, cart)

	if err != nil {
		return nil, fmt.Errorf("UpdateRacket.Update fail, userID: %d, error %s", req.UserID, err)
	}

	return cart, nil
}

func (s *CartService) GetCartByID(ctx context.Context, userID int) (*model.Cart, error) {

	cart, err := s.repo.GetCartByID(ctx, userID)

	if err != nil {

		cart := &model.Cart{
			UserID: userID,
		}

		err = s.repo.Create(ctx, cart)

		if err != nil {
			return nil, fmt.Errorf("GetCartByID.Create fail, userID: %d, error %s", userID, err)
		}

		return cart, nil
	}

	return cart, nil
}
