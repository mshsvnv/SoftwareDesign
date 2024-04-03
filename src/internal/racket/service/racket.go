package service

import (
	"context"

	"src/internal/racket/dto"
	"src/internal/racket/model"
	"src/internal/racket/repository"
	"src/pkg/utils"
)

type IRacketService interface {
	ListRackets(ctx context.Context, req *dto.ListRacketReq) ([]*model.Racket, error)
	GetRacketByID(ctx context.Context, id string) (*model.Racket, error)
	Create(ctx context.Context, req *dto.CreateRacketReq) (*model.Racket, error)
	Update(ctx context.Context, id string, req *dto.UpdateRacketReq) (*model.Racket, error)
}

type RacketService struct {
	repo repository.IRacketRepository
}

func NewRacketService(
	repo repository.IRacketRepository,
) *RacketService {
	return &RacketService{
		repo: repo,
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
