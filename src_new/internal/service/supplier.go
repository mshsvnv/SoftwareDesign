package service

import (
	"context"
	"fmt"

	"src_new/internal/dto"
	"src_new/internal/model"
	repo "src_new/internal/repository"
	"src_new/pkg/utils"
)

//go:generate mockery --name=ISupplierService
type ISupplierService interface {
	CreateSupplier(ctx context.Context, req *dto.CreateSupplierReq) (*model.Supplier, error)
	RemoveSupplier(ctx context.Context, id int) error
	UpdateSupplier(ctx context.Context, req *dto.UpdateSupplierReq) error
	GetSupplierByID(ctx context.Context, id int) (*model.Supplier, error)
}

type SupplierService struct {
	repo repo.ISupplierRepository
}

func NewSupplierService(repo repo.ISupplierRepository) *SupplierService {
	return &SupplierService{
		repo: repo,
	}
}

func (s *SupplierService) CreateSupplier(ctx context.Context, req *dto.CreateSupplierReq) (*model.Supplier, error) {

	var supplier model.Supplier

	utils.Copy(&supplier, req)

	err := s.repo.Create(ctx, &supplier)

	if err != nil {
		return nil, fmt.Errorf("CreateSupplier.GetSupplierByID fail, %s", err)
	}

	return &supplier, nil

}

func (s *SupplierService) RemoveSupplier(ctx context.Context, id int) error {

	supplier, err := s.repo.GetSupplierByID(ctx, id)

	if supplier == nil {
		return fmt.Errorf("RemoveSupplier.GetSupplierByID fail, %d %s", id, err)
	}

	err = s.repo.Remove(ctx, supplier.Email)

	if err != nil {
		return fmt.Errorf("RemoveSupplier.Remove fail, %d %s", id, err)
	}

	return nil
}

func (s *SupplierService) UpdateSupplier(ctx context.Context, req *dto.UpdateSupplierReq) error {

	supplier, err := s.repo.GetSupplierByID(ctx, req.ID)

	if supplier == nil {
		return fmt.Errorf("UpdateSupplier.GetSupplierByID fail, %d %s", req.ID, err)
	}

	utils.Copy(&supplier, req)

	err = s.repo.Update(ctx, supplier)

	if err != nil {
		return fmt.Errorf("UpdateSupplier.Update fail, %s", err)
	}

	return nil
}

func (s *SupplierService) GetSupplierByID(ctx context.Context, id int) (*model.Supplier, error) {

	supplier, err := s.repo.GetSupplierByID(ctx, id)

	if err != nil {
		return nil, fmt.Errorf("GetSupplierByID.GetSupplierByID fail, %s", err)
	}

	return supplier, nil
}
