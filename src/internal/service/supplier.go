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

type ISupplierService interface {
	CreateSupplier(ctx context.Context, req *dto.CreateSupplierReq) (*model.Supplier, error)
	RemoveSupplier(ctx context.Context, email string) error
	GetSupplierByID(ctx context.Context, id int) (*model.Supplier, error)
	GetSupplierByEmail(ctx context.Context, email string) (*model.Supplier, error)
	GetAllSuppliers(ctx context.Context) ([]*model.Supplier, error)
}

type SupplierService struct {
	logger logging.Interface
	repo   repo.ISupplierRepository
}

func NewSupplierService(logger logging.Interface, repo repo.ISupplierRepository) *SupplierService {
	return &SupplierService{
		logger: logger,
		repo:   repo,
	}
}

func (s *SupplierService) CreateSupplier(ctx context.Context, req *dto.CreateSupplierReq) (*model.Supplier, error) {

	s.logger.Infof("create supplier name %s", req.Name)
	_, err := s.repo.GetSupplierByEmail(ctx, req.Email)

	if err == nil {
		s.logger.Errorf("get supplier by email fail")
		return nil, fmt.Errorf("get supplier by email fail, error %s", err)
	}

	var supplier model.Supplier
	utils.Copy(&supplier, req)

	err = s.repo.Create(ctx, &supplier)

	if err != nil {
		s.logger.Errorf("create fail, error %s", err.Error())
		return nil, fmt.Errorf("create fail, error %s", err)
	}

	return &supplier, nil

}

func (s *SupplierService) RemoveSupplier(ctx context.Context, email string) error {

	s.logger.Infof("remove supplier email %s", email)
	supplier, err := s.repo.GetSupplierByEmail(ctx, email)

	if supplier == nil {
		s.logger.Errorf("get supplier by email fail, error %s", err.Error())
		return fmt.Errorf("get supplier by email fail, error %s", err)
	}

	err = s.repo.Remove(ctx, supplier.Email)

	if err != nil {
		s.logger.Errorf("remove fail, error %s", err.Error())
		return fmt.Errorf("remove fail, error %s", err)
	}

	return nil
}

func (s *SupplierService) GetSupplierByID(ctx context.Context, id int) (*model.Supplier, error) {

	s.logger.Infof("get supplier by id")
	supplier, err := s.repo.GetSupplierByID(ctx, id)

	if err != nil {
		s.logger.Errorf("get supplier by id fail, error %s", err.Error())
		return nil, fmt.Errorf("get supplier by id fail, error %s", err)
	}

	return supplier, nil
}

func (s *SupplierService) GetSupplierByEmail(ctx context.Context, email string) (*model.Supplier, error) {

	s.logger.Infof("get supplier by email %s", email)
	supplier, err := s.repo.GetSupplierByEmail(ctx, email)

	if err != nil {
		s.logger.Errorf("get supplier by email fail, error %s", err.Error())
		return nil, fmt.Errorf("get supplier by email fail, error %s", err)
	}

	return supplier, nil
}

func (s *SupplierService) GetAllSuppliers(ctx context.Context) ([]*model.Supplier, error) {

	s.logger.Infof("get all suppliers")
	suppliers, err := s.repo.GetAllSuppliers(ctx)

	if err != nil {
		s.logger.Errorf("get all suppliers fail, error %s", err.Error())
		return nil, fmt.Errorf("get all suppliers fail, error %s", err)
	}

	return suppliers, nil
}
