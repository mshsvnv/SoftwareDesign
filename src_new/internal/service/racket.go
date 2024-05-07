package service

import (
	"context"
	"fmt"

	"src_new/internal/dto"
	"src_new/internal/model"
	repo "src_new/internal/repository"
	"src_new/pkg/utils"
)

type IRacketService interface {
	CreateRacket(ctx context.Context, req *dto.CreateRacketReq) (*model.Racket, error)
	RemoveRacket(ctx context.Context, id int) error
	UpdateRacket(ctx context.Context, req *dto.UpdateRacketReq) error
	GetRacketByID(ctx context.Context, id int) (*model.Racket, error)
	GetAllRackets(ctx context.Context) ([]*model.Racket, error)
}

type RacketService struct {
	repo repo.IRacketRepository
}

func NewRacketService(repo repo.IRacketRepository) *RacketService {
	return &RacketService{
		repo: repo,
	}
}

func (s *RacketService) CreateRacket(ctx context.Context, req *dto.CreateRacketReq) (*model.Racket, error) {

	var racket model.Racket

	utils.Copy(&racket, req)

	err := s.repo.Create(ctx, &racket)

	if err != nil {
		return nil, fmt.Errorf("CreateRacket.GetRacketByID fail, %s", err)
	}

	return &racket, nil

}

func (s *RacketService) RemoveRacket(ctx context.Context, id int) error {

	racket, err := s.repo.GetRacketByID(ctx, id)

	if racket == nil {
		return fmt.Errorf("RemoveRacket.GetRacketByID fail, %d %s", id, err)
	}

	err = s.repo.Remove(ctx, id)

	if err != nil {
		return fmt.Errorf("RemoveRacket.Remove fail, %d %s", id, err)
	}

	return nil
}

func (s *RacketService) UpdateRacket(ctx context.Context, req *dto.UpdateRacketReq) error {

	racket, err := s.repo.GetRacketByID(ctx, req.ID)

	if racket == nil {
		return fmt.Errorf("UpdateRacket.GetRacketByID fail, %d %s", req.ID, err)
	}

	utils.Copy(&racket, req)

	err = s.repo.Update(ctx, racket)

	if err != nil {
		return fmt.Errorf("UpdateRacket.Update fail, %s", err)
	}

	return nil
}

func (s *RacketService) GetRacketByID(ctx context.Context, id int) (*model.Racket, error) {

	racket, err := s.repo.GetRacketByID(ctx, id)

	if err != nil {
		return nil, fmt.Errorf("GetRacketByID.GetRacketByID fail, %s", err)
	}

	return racket, nil
}

func (s *RacketService) GetAllRackets(ctx context.Context) ([]*model.Racket, error) {

	rackets, err := s.repo.GetAllRackets(ctx)

	if err != nil {
		return nil, fmt.Errorf("GetAllRackets.GetAllRackets fail, %s", err)
	}

	return rackets, nil
}
