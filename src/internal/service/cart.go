package service

import (
	"context"
	"fmt"

	"src/internal/dto"
	"src/internal/model"

	repo "src/internal/repository"
	"src/pkg/logging"
)

type ICartService interface {
	AddRacket(ctx context.Context, req *dto.AddRacketCartReq) (*model.Cart, error)
	RemoveRacket(ctx context.Context, req *dto.RemoveRacketCartReq) (*model.Cart, error)
	UpdateRacket(ctx context.Context, req *dto.UpdateRacketCartReq) (*model.Cart, error)
	GetCartByID(ctx context.Context, userID int) (*model.Cart, error)
}

type CartService struct {
	logger     logging.Interface
	repo       repo.ICartRepository
	repoRacket repo.IRacketRepository
}

func NewCartService(logger logging.Interface, repo repo.ICartRepository, repoRacket repo.IRacketRepository) *CartService {
	return &CartService{
		logger:     logger,
		repo:       repo,
		repoRacket: repoRacket,
	}
}

func (s *CartService) UpdateRacket(ctx context.Context, req *dto.UpdateRacketCartReq) (*model.Cart, error) {

	// TODO
	cart, _ := s.repo.GetCartByID(ctx, req.UserID)

	for _, lines := range cart.Lines {

		if lines.RacketID == req.RacketID {

			racket, _ := s.repoRacket.GetRacketByID(ctx, req.RacketID)

			curQuantity := lines.Quantity + req.Quantity
			if curQuantity == -1 {
				return cart, nil
			} else if curQuantity >= racket.Quantity {
				curQuantity = racket.Quantity
			}

			lines.Quantity = curQuantity
			cart.Quantity += req.Quantity
			cart.TotalPrice += float32(req.Quantity) * lines.Price

			err := s.repo.Update(ctx, cart)
			if err != nil {
				s.logger.Errorf("UpdateRacket fail, error %s", err.Error())
				return nil, fmt.Errorf("UpdateRacket fail, error %s", err)
			}

			return cart, nil
		}
	}

	return cart, nil
}

func (s *CartService) AddRacket(ctx context.Context, req *dto.AddRacketCartReq) (*model.Cart, error) {

	s.logger.Infof("add racket to user %d, racket %d", req.UserID, req.RacketID)
	cart, err := s.repo.GetCartByID(ctx, req.UserID)

	if err != nil {

		racket, err := s.repoRacket.GetRacketByID(ctx, req.RacketID)

		if err != nil {
			s.logger.Errorf("getRacketByID fail, error %s", err.Error())
			return nil, fmt.Errorf("getRacketByID fail, error %s", err)
		}

		if req.Quantity <= 0 {
			s.logger.Errorf("addRacketByID fail, error request quantity <= 0")
			return nil, fmt.Errorf("getRacketByID fail, error %s", err)
		} else if req.Quantity >= racket.Quantity {
			req.Quantity = racket.Quantity
			racket.Quantity = 0
		} else {
			racket.Quantity -= req.Quantity
		}

		// racket.Quantity -= req.Quantity

		cart = &model.Cart{
			UserID:     req.UserID,
			TotalPrice: float32(racket.Price) * float32(req.Quantity),
			Quantity:   req.Quantity,
			Lines: []*model.CartLine{{
				RacketID: req.RacketID,
				Quantity: req.Quantity,
				Price:    float32(racket.Price),
			}},
		}

		// err = s.repoRacket.Update(ctx, racket)

		// if err != nil {
		// 	s.logger.Errorf("update racket fail, error %s", err.Error())
		// 	return nil, fmt.Errorf("update racket fail, error %s", err)
		// }

		err = s.repo.Create(ctx, cart)

		if err != nil {
			s.logger.Errorf("create cart fail, error %s", err.Error())
			return nil, fmt.Errorf("create cart fail, error %s", err)
		}

		return cart, nil
	}

	for _, line := range cart.Lines {

		if line.RacketID == req.RacketID {
			return cart, nil
		}
	}

	racket, err := s.repoRacket.GetRacketByID(ctx, req.RacketID)

	if err != nil {
		s.logger.Errorf("get racket fail, error %s", err.Error())
		return nil, fmt.Errorf("get racket fail, error %s", err)
	}

	if req.Quantity <= 0 {
		s.logger.Errorf("addRacketByID fail, error request quantity <= 0")
		return nil, fmt.Errorf("getRacketByID fail, error %s", err)
	} else if req.Quantity >= racket.Quantity {
		req.Quantity = racket.Quantity
		racket.Quantity = 0
	} else {
		racket.Quantity -= req.Quantity
	}

	err = s.repo.AddRacket(ctx, req)

	if err != nil {
		s.logger.Errorf("add racket fail, error %s", err.Error())
		return nil, fmt.Errorf("add racket fail, error %s", err)
	}

	// err = s.repoRacket.Update(ctx, racket)

	// if err != nil {
	// 	s.logger.Errorf("update racket fail, error %s", err.Error())
	// 	return nil, fmt.Errorf("update racket fail, error %s", err)
	// }

	cart.Lines = append(cart.Lines,
		&model.CartLine{
			RacketID: req.RacketID,
			Quantity: req.Quantity,
			Price:    float32(racket.Price),
		})

	cart.Quantity += req.Quantity
	cart.TotalPrice += float32(racket.Price) * float32(req.Quantity)

	err = s.repo.Update(ctx, cart)

	if err != nil {
		s.logger.Errorf("update cart fail, error %s", err.Error())
		return nil, fmt.Errorf("update cart fail, error %s", err)
	}

	return cart, nil
}

func (s *CartService) RemoveRacket(ctx context.Context, req *dto.RemoveRacketCartReq) (*model.Cart, error) {

	s.logger.Infof("add racket to user %d, racket %d", req.UserID, req.RacketID)
	cart, err := s.repo.GetCartByID(ctx, req.UserID)

	if err != nil {

		cart = &model.Cart{
			UserID: req.UserID,
		}

		err = s.repo.Create(ctx, cart)

		if err != nil {
			s.logger.Errorf("create cart fail, error %s", err.Error())
			return nil, fmt.Errorf("create cart fail, error %s", err)
		}

		return cart, nil
	}

	for i := 0; i < len(cart.Lines); i++ {

		if cart.Lines[i].RacketID == req.RacketID {

			racket, err := s.repoRacket.GetRacketByID(ctx, req.RacketID)

			if err != nil {
				s.logger.Errorf("add racket fail, error %s", err.Error())
				return nil, fmt.Errorf("add racket fail, error %s", err)
			}

			cart.Quantity -= cart.Lines[i].Quantity
			cart.TotalPrice -= float32(racket.Price) * float32(cart.Lines[i].Quantity)

			cart.Lines = append(cart.Lines[:i], cart.Lines[i+1:]...)
			break
		}
	}

	err = s.repo.RemoveRacket(ctx, req)

	if err != nil {
		s.logger.Errorf("remove racket fail, error %s", err.Error())
		return nil, fmt.Errorf("remove racket fail, error %s", err)
	}

	err = s.repo.Update(ctx, cart)

	if err != nil {
		s.logger.Errorf("update racket fail, error %s", err.Error())
		return nil, fmt.Errorf("update racket fail, error %s", err)
	}

	return cart, nil
}

func (s *CartService) GetCartByID(ctx context.Context, userID int) (*model.Cart, error) {

	s.logger.Infof("get cart, user %d", userID)
	cart, err := s.repo.GetCartByID(ctx, userID)

	if err != nil {

		cart := &model.Cart{
			UserID: userID,
		}

		err = s.repo.Create(ctx, cart)

		if err != nil {
			s.logger.Errorf("create cart fail, error %s", err.Error())
			return nil, fmt.Errorf("create cart fail, error %s", err)
		}

		return cart, nil
	} else {
		err = s.repo.Update(ctx, cart)

		if err != nil {
			s.logger.Errorf("update cart fail, error %s", err.Error())
			return nil, fmt.Errorf("update cart fail, error %s", err)
		}
	}

	return cart, nil
}
