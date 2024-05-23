package mypostgres

import (
	"context"
	"src_new/internal/model"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSupplierRepositoryCreate(t *testing.T) {

	ctx := context.Background()

	repo := NewSupplierRepository(testDB)

	supplier := &model.Supplier{
		Name:  "IP petrov",
		Email: "petrov@mail.ru",
		Town:  "Armavir",
		Phone: "8-800-555-35-35",
	}

	err := repo.Create(ctx, supplier)
	require.NoError(t, err)
	require.NotEmpty(t, supplier.ID)

	err = repo.Remove(ctx, supplier.Email)
	require.NoError(t, err)
}

func TestSupplierRepositoryUpdate(t *testing.T) {

	ctx := context.Background()

	repo := NewSupplierRepository(testDB)

	supplier := &model.Supplier{
		Name:  "IP petrov",
		Email: "petrov@mail.ru",
		Town:  "Armavir",
		Phone: "8-800-555-35-35",
	}

	err := repo.Create(ctx, supplier)
	require.NoError(t, err)
	require.NotEmpty(t, supplier.ID)

	supplier.Town = "Lipetsck"
	err = repo.Update(ctx, supplier)
	require.NoError(t, err)

	err = repo.Remove(ctx, supplier.Email)
	require.NoError(t, err)
}

func TestSupplierRepositoryGetSupplierByID(t *testing.T) {

	ctx := context.Background()

	repo := NewSupplierRepository(testDB)

	supplier := &model.Supplier{
		Name:  "IP petrov",
		Email: "petrov@mail.ru",
		Town:  "Armavir",
		Phone: "8-800-555-35-35",
	}

	err := repo.Create(ctx, supplier)
	require.NoError(t, err)
	require.NotEmpty(t, supplier.ID)

	res, err := repo.GetSupplierByID(ctx, supplier.ID)
	require.NoError(t, err)
	require.Equal(t, res, supplier)

	err = repo.Remove(ctx, supplier.Email)
	require.NoError(t, err)
}

func TestSupplierRepositoryGetSupplierByEmail(t *testing.T) {

	ctx := context.Background()

	repo := NewSupplierRepository(testDB)

	supplier := &model.Supplier{
		Name:  "IP petrov",
		Email: "petrov@mail.ru",
		Town:  "Armavir",
		Phone: "8-800-555-35-35",
	}

	err := repo.Create(ctx, supplier)
	require.NoError(t, err)
	require.NotEmpty(t, supplier.ID)

	res, err := repo.GetSupplierByEmail(ctx, supplier.Email)
	require.NoError(t, err)
	require.Equal(t, res, supplier)

	err = repo.Remove(ctx, supplier.Email)
	require.NoError(t, err)
}

func TestSupplierRepositoryRemove(t *testing.T) {

	ctx := context.Background()

	repo := NewSupplierRepository(testDB)

	supplier := &model.Supplier{
		Name:  "IP petrov",
		Email: "petrov@mail.ru",
		Town:  "Armavir",
		Phone: "8-800-555-35-35",
	}

	err := repo.Create(ctx, supplier)
	require.NoError(t, err)
	require.NotEmpty(t, supplier.ID)

	err = repo.Remove(ctx, supplier.Email)
	require.NoError(t, err)
}
