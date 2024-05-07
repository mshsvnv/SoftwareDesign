package mypostgres

import (
	"context"
	"src_new/internal/model"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRacketRepositoryCreate(t *testing.T) {

	ctx := context.Background()

	repo := NewRacketRepository(testDB)
	repoSupplier := NewSupplierRepository(testDB)

	supplier := &model.Supplier{
		Name:  "IP Ivanov",
		Email: "ivanov@mail.ru",
		Town:  "Armavir",
		Phone: "8-800-555-35-35",
	}

	err := repoSupplier.Create(ctx, supplier)
	require.NoError(t, err)
	require.NotEmpty(t, supplier.ID)

	racket := &model.Racket{
		SupplierID: supplier.ID,
		Brand:      "Babolat",
		Weight:     1000,
		Balance:    3.5,
		HeadSize:   20.2,
		Quantity:   100,
		Price:      100,
	}

	err = repo.Create(ctx, racket)
	require.NoError(t, err)
	require.NotEmpty(t, racket.ID)

	err = repo.Remove(ctx, racket.ID)
	require.NoError(t, err)

	err = repoSupplier.Remove(ctx, supplier.Email)
	require.NoError(t, err)
}

func TestRacketRepositoryUpdate(t *testing.T) {

	ctx := context.Background()

	repo := NewRacketRepository(testDB)
	repoSupplier := NewSupplierRepository(testDB)

	supplier := &model.Supplier{
		Name:  "IP Ivanov",
		Email: "ivanov@mail.ru",
		Town:  "Armavir",
		Phone: "8-800-555-35-35",
	}

	err := repoSupplier.Create(ctx, supplier)
	require.NoError(t, err)
	require.NotEmpty(t, supplier.ID)

	racket := &model.Racket{
		SupplierID: supplier.ID,
		Brand:      "Babolat",
		Weight:     1000,
		Balance:    3.5,
		HeadSize:   20.2,
		Quantity:   100,
		Price:      100,
	}

	err = repo.Create(ctx, racket)
	require.NoError(t, err)
	require.NotEmpty(t, racket.ID)

	racket.Brand = "Head"
	err = repo.Update(ctx, racket)
	require.NoError(t, err)

	err = repo.Remove(ctx, racket.ID)
	require.NoError(t, err)

	err = repoSupplier.Remove(ctx, supplier.Email)
	require.NoError(t, err)
}

func TestRacketRepositoryGetRacketByID(t *testing.T) {

	ctx := context.Background()

	repo := NewRacketRepository(testDB)
	repoSupplier := NewSupplierRepository(testDB)

	supplier := &model.Supplier{
		Name:  "IP Ivanov",
		Email: "ivanov@mail.ru",
		Town:  "Armavir",
		Phone: "8-800-555-35-35",
	}

	err := repoSupplier.Create(ctx, supplier)
	require.NoError(t, err)
	require.NotEmpty(t, supplier.ID)

	racket := &model.Racket{
		SupplierID: supplier.ID,
		Brand:      "Babolat",
		Weight:     1000,
		Balance:    3.5,
		HeadSize:   20.2,
		Quantity:   100,
		Price:      100,
	}

	err = repo.Create(ctx, racket)
	require.NoError(t, err)
	require.NotEmpty(t, racket.ID)

	res, err := repo.GetRacketByID(ctx, racket.ID)
	require.NoError(t, err)
	require.Equal(t, res, racket)

	err = repo.Remove(ctx, racket.ID)
	require.NoError(t, err)

	err = repoSupplier.Remove(ctx, supplier.Email)
	require.NoError(t, err)
}

func TestRacketRepositoryGetAllRackets(t *testing.T) {

	ctx := context.Background()

	repo := NewRacketRepository(testDB)
	repoSupplier := NewSupplierRepository(testDB)

	supplier := &model.Supplier{
		Name:  "IP Ivanov",
		Email: "ivanov@mail.ru",
		Town:  "Armavir",
		Phone: "8-800-555-35-35",
	}

	err := repoSupplier.Create(ctx, supplier)
	require.NoError(t, err)
	require.NotEmpty(t, supplier.ID)

	racket := &model.Racket{
		SupplierID: supplier.ID,
		Brand:      "Babolat",
		Weight:     1000,
		Balance:    3.5,
		HeadSize:   20.2,
		Quantity:   100,
		Price:      100,
	}

	err = repo.Create(ctx, racket)
	require.NoError(t, err)
	require.NotEmpty(t, racket.ID)

	res, err := repo.GetAllRackets(ctx)
	require.NoError(t, err)
	require.Equal(t, res[0], racket)

	err = repo.Remove(ctx, racket.ID)
	require.NoError(t, err)

	err = repoSupplier.Remove(ctx, supplier.Email)
	require.NoError(t, err)
}