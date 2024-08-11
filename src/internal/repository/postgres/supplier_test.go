package mypostgres

import (
	"context"
	"src/internal/model"
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
		Phone: "88005553535",
	}

	err := repo.Create(ctx, supplier)
	require.NoError(t, err)
	require.NotEmpty(t, supplier.Email)

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
		Phone: "88005553535",
	}

	err := repo.Create(ctx, supplier)
	require.NoError(t, err)
	require.NotEmpty(t, supplier.Email)

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
		Phone: "88005553535",
	}

	err := repo.Create(ctx, supplier)
	require.NoError(t, err)
	require.NotEmpty(t, supplier.Email)

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
		Phone: "88005553535",
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
		Phone: "88005553535",
	}

	err := repo.Create(ctx, supplier)
	require.NoError(t, err)
	require.NotEmpty(t, supplier.ID)

	err = repo.Remove(ctx, supplier.Email)
	require.NoError(t, err)
}
