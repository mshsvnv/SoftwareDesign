package service

import (
	"context"
	"fmt"

	"src_new/internal/dto"
	"src_new/internal/model"
	repo "src_new/internal/repository"
	"src_new/pkg/utils"
)

type ISupplierService interface {
	CreateSupplier(ctx context.Context, req *dto.CreateSupplierReq) (*model.Supplier, error)
	RemoveSupplier(ctx context.Context, email string) error
	UpdateSupplier(ctx context.Context, req *dto.UpdateSupplierReq) error
	GetSupplierByID(ctx context.Context, id int) (*model.Supplier, error)
	GetSupplierByEmail(ctx context.Context, email string) (*model.Supplier, error)
	GetAllSuppliers(ctx context.Context) ([]*model.Supplier, error)
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

	_, err := s.repo.GetSupplierByEmail(ctx, req.Email)

	if err == nil {
		return nil, fmt.Errorf("CreateSupplier.GetSupplierByEmail fail, %s", err)
	}

	var supplier model.Supplier
	utils.Copy(&supplier, req)

	err = s.repo.Create(ctx, &supplier)

	if err != nil {
		return nil, fmt.Errorf("CreateSupplier.GetSupplierByID fail, %s", err)
	}

	return &supplier, nil

}

func (s *SupplierService) RemoveSupplier(ctx context.Context, email string) error {

	supplier, err := s.repo.GetSupplierByEmail(ctx, email)

	if supplier == nil {
		return fmt.Errorf("RemoveSupplier.GetSupplierByID fail, %s", err)
	}

	err = s.repo.Remove(ctx, supplier.Email)

	if err != nil {
		return fmt.Errorf("RemoveSupplier.Remove fail, %s", err)
	}

	return nil
}

func (s *SupplierService) UpdateSupplier(ctx context.Context, req *dto.UpdateSupplierReq) error {

	supplier, err := s.repo.GetSupplierByEmail(ctx, req.Email)

	if supplier == nil {
		return fmt.Errorf("UpdateSupplier.GetSupplierByEmail fail, %s %s", req.Email, err)
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

func (s *SupplierService) GetSupplierByEmail(ctx context.Context, email string) (*model.Supplier, error) {

	supplier, err := s.repo.GetSupplierByEmail(ctx, email)

	if err != nil {
		return nil, fmt.Errorf("GetSupplierByEmail.GetSupplierByID fail, %s", err)
	}

	return supplier, nil
}

func (s *SupplierService) GetAllSuppliers(ctx context.Context) ([]*model.Supplier, error) {

	suppliers, err := s.repo.GetAllSuppliers(ctx)

	if err != nil {
		return nil, fmt.Errorf("GetAllSuppliera.GetAllSuppliera fail, %s", err)
	}

	return suppliers, nil
}
