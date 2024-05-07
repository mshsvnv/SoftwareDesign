package repository

import (
	"context"
	"src_new/internal/model"
)

//go:generate mockery --name=ISupplierRepository
type ISupplierRepository interface {
	Create(ctx context.Context, supplier *model.Supplier) error
	Update(ctx context.Context, supplier *model.Supplier) error
	Remove(ctx context.Context, email string) error
	GetSupplierByID(ctx context.Context, id int) (*model.Supplier, error)
}

// type RacketSupplier struct {
// 	RacketID   int
// 	SupplierID int
// 	Quantitiy  int
// }

// type SupplierRepository struct {
// 	db []*model.Supplier
// }

// func NewSupplierRepository() *SupplierRepository {
// 	return &SupplierRepository{}
// }

// func (r *SupplierRepository) Create(ctx context.Context, supplier *model.Supplier) error {

// 	r.db = append(r.db, supplier)
// 	r.db[len(r.db)-1].ID = len(r.db) - 1

// 	return nil
// }

// func (r *SupplierRepository) Update(ctx context.Context, supplier *model.Supplier) error {

// 	for i := 0; i < len(r.db); i++ {

// 		if r.db[i].ID == supplier.ID {
// 			r.db[i] = supplier

// 			return nil
// 		}
// 	}

// 	return fmt.Errorf("no supplier")
// }

// func (r *SupplierRepository) Remove(ctx context.Context, id int) error {

// 	for i := 0; i < len(r.db); i++ {

// 		if r.db[i].ID == id {

// 			r.db = append(r.db[:i], r.db[i+1:]...)
// 			return nil
// 		}
// 	}

// 	return fmt.Errorf("no supplier")
// }

// func (r *SupplierRepository) GetSupplierByID(ctx context.Context, id int) (*model.Supplier, error) {

// 	for i := 0; i < len(r.db); i++ {

// 		if r.db[i].ID == id {
// 			return r.db[i], nil
// 		}
// 	}

// 	return nil, fmt.Errorf("no supplier")
// }

// func (r *SupplierRepository) GetAllSupplier(ctx context.Context) ([]*model.Supplier, error) {

// 	if len(r.db) == 0 {
// 		return nil, fmt.Errorf("no supplier")
// 	}

// 	return r.db, nil
// }
