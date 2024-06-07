package service

import (
	"context"
	"fmt"

	"src/internal/dto"
	"src/internal/model"
	repo "src/internal/repository"
	"src/pkg/logging"
	"src/pkg/utils"
)

type IRacketService interface {
	CreateRacket(ctx context.Context, req *dto.CreateRacketReq) (*model.Racket, error)
	UpdateRacket(ctx context.Context, req *dto.UpdateRacketReq) error
	GetRacketByID(ctx context.Context, id int) (*model.Racket, error)
	GetAllAvaliableRackets(ctx context.Context) ([]*model.Racket, error)
	GetAllRackets(ctx context.Context) ([]*model.Racket, error)
}

type RacketService struct {
	logger       logging.Interface
	repo         repo.IRacketRepository
	repoSupplier repo.ISupplierRepository
}

func NewRacketService(logger logging.Interface, repo repo.IRacketRepository, repoSupplier repo.ISupplierRepository) *RacketService {
	return &RacketService{
		logger:       logger,
		repo:         repo,
		repoSupplier: repoSupplier,
	}
}

func (s *RacketService) CreateRacket(ctx context.Context, req *dto.CreateRacketReq) (*model.Racket, error) {

	s.logger.Infof("create racket")
	var racket model.Racket

	utils.Copy(&racket, req)

	if racket.Quantity > 0 {
		racket.Avaliable = true
	} else if racket.Quantity == 0 {
		racket.Avaliable = false
	} else {
		s.logger.Errorf("unavaliable amount of rackets, error")
		return nil, fmt.Errorf("unavaliable amount of rackets, error")
	}

	supplier, err := s.repoSupplier.GetSupplierByEmail(ctx, req.SupplierEmail)

	if err != nil {
		s.logger.Errorf("get racket by email fail, error %s", err.Error())
		return nil, fmt.Errorf("get racket by email fail, error %s", err)
	}

	racket.SupplierEmail = supplier.Email

	err = s.repo.Create(ctx, &racket)

	if err != nil {
		s.logger.Errorf("create fail, error %s", err.Error())
		return nil, fmt.Errorf("create fail, error %s", err)
	}

	return &racket, nil

}

func (s *RacketService) UpdateRacket(ctx context.Context, req *dto.UpdateRacketReq) error {

	s.logger.Infof("update racket by id %d", req.ID)
	racket, err := s.repo.GetRacketByID(ctx, req.ID)

	if racket == nil {
		s.logger.Errorf("get racket by id, error %s", err.Error())
		return fmt.Errorf("get racket by id, error %s", err)
	}

	if req.Quantity <= 0 {
		s.logger.Errorf("unavalibale amount of rackets")
		return fmt.Errorf("unavalibale amount of rackets")
	}

	racket.Quantity = req.Quantity

	err = s.repo.Update(ctx, racket)

	if err != nil {
		s.logger.Errorf("update racket fail, error %s", err.Error())
		return fmt.Errorf("update racket fail, error %s", err)
	}

	return nil
}

func (s *RacketService) GetRacketByID(ctx context.Context, id int) (*model.Racket, error) {

	s.logger.Infof("get racket by id %d", id)
	racket, err := s.repo.GetRacketByID(ctx, id)

	if err != nil {
		s.logger.Errorf("get racket by id fail, error %s", err.Error())
		return nil, fmt.Errorf("get racket by id fail, error %s", err)
	}

	return racket, nil
}

func (s *RacketService) GetAllAvaliableRackets(ctx context.Context) ([]*model.Racket, error) {

	s.logger.Infof("get all avaliable rackets")
	rackets, err := s.repo.GetAllAvaliableRackets(ctx)

	if err != nil {
		s.logger.Errorf("get all avaliable fail, error %s", err.Error())
		return nil, fmt.Errorf("get all avaliable fail, error %s", err)
	}

	return rackets, nil
}

func (s *RacketService) GetAllRackets(ctx context.Context) ([]*model.Racket, error) {

	s.logger.Infof("get all rackets")
	rackets, err := s.repo.GetAllRackets(ctx)

	if err != nil {
		s.logger.Errorf("get all fail, error %s", err.Error())
		return nil, fmt.Errorf("get all fail, error %s", err)
	}

	return rackets, nil
}
