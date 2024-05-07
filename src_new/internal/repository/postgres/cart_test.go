package mypostgres

import (
	"context"
	"src_new/internal/model"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCartRepositoryCreate(t *testing.T) {

	ctx := context.Background()

	repo := NewCartRepository(testDB)
	repoUser := NewUserRepository(testDB)
	repoSupplier := NewSupplierRepository(testDB)
	repoRacket := NewRacketRepository(testDB)

	user := &model.User{
		Name:     "Ivan",
		Surname:  "Ivanov",
		Email:    "ivanov@mail.ru",
		Password: "123",
		Role:     "Customer",
	}

	supplier := &model.Supplier{
		Name:  "IP Ivanov",
		Email: "ivanov@mail.ru",
		Town:  "Armavir",
		Phone: "8-800-555-35-35",
	}

	err := repoUser.Create(ctx, user)
	require.NoError(t, err)
	require.NotEmpty(t, user.ID)

	err = repoSupplier.Create(ctx, supplier)
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

	err = repoRacket.Create(ctx, racket)
	require.NoError(t, err)
	require.NotEmpty(t, racket.ID)

	quantity := 10
	cart := &model.Cart{
		UserID:     user.ID,
		Quantity:   quantity,
		TotalPrice: float32(quantity) * float32(racket.Price),
		Lines: []*model.CartLine{{
			RacketID: racket.ID,
			Quantity: quantity,
		}},
	}

	err = repo.Create(ctx, cart)
	require.NoError(t, err)

	err = repo.Remove(ctx, user.ID)
	require.NoError(t, err)

	err = repoUser.Remove(ctx, user.Email)
	require.NoError(t, err)

	err = repoRacket.Remove(ctx, racket.ID)
	require.NoError(t, err)

	err = repoSupplier.Remove(ctx, supplier.Email)
	require.NoError(t, err)
}

func TestCartRepositoryUpdate(t *testing.T) {

	ctx := context.Background()

	repo := NewCartRepository(testDB)
	repoUser := NewUserRepository(testDB)
	repoSupplier := NewSupplierRepository(testDB)
	repoRacket := NewRacketRepository(testDB)

	user := &model.User{
		Name:     "Ivan",
		Surname:  "Ivanov",
		Email:    "ivanov@mail.ru",
		Password: "123",
		Role:     "Customer",
	}

	supplier := &model.Supplier{
		Name:  "IP Ivanov",
		Email: "ivanov@mail.ru",
		Town:  "Armavir",
		Phone: "8-800-555-35-35",
	}

	err := repoUser.Create(ctx, user)
	require.NoError(t, err)
	require.NotEmpty(t, user.ID)

	err = repoSupplier.Create(ctx, supplier)
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

	err = repoRacket.Create(ctx, racket)
	require.NoError(t, err)
	require.NotEmpty(t, racket.ID)

	quantity := 10
	cart := &model.Cart{
		UserID:     user.ID,
		Quantity:   quantity,
		TotalPrice: float32(quantity) * float32(racket.Price),
		Lines: []*model.CartLine{{
			RacketID: racket.ID,
			Quantity: quantity,
		}},
	}

	err = repo.Create(ctx, cart)
	require.NoError(t, err)

	cart.Quantity -= 1
	cart.TotalPrice -= 100

	err = repo.Update(ctx, cart)
	require.NoError(t, err)

	err = repo.Remove(ctx, user.ID)
	require.NoError(t, err)

	err = repoUser.Remove(ctx, user.Email)
	require.NoError(t, err)

	err = repoRacket.Remove(ctx, racket.ID)
	require.NoError(t, err)

	err = repoSupplier.Remove(ctx, supplier.Email)
	require.NoError(t, err)
}

func TestCartRepositoryGetCartByID(t *testing.T) {

	ctx := context.Background()

	repo := NewCartRepository(testDB)
	repoUser := NewUserRepository(testDB)
	repoSupplier := NewSupplierRepository(testDB)
	repoRacket := NewRacketRepository(testDB)

	user := &model.User{
		Name:     "Ivan",
		Surname:  "Ivanov",
		Email:    "ivanov@mail.ru",
		Password: "123",
		Role:     "Customer",
	}

	supplier := &model.Supplier{
		Name:  "IP Ivanov",
		Email: "ivanov@mail.ru",
		Town:  "Armavir",
		Phone: "8-800-555-35-35",
	}

	err := repoUser.Create(ctx, user)
	require.NoError(t, err)
	require.NotEmpty(t, user.ID)

	err = repoSupplier.Create(ctx, supplier)
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

	err = repoRacket.Create(ctx, racket)
	require.NoError(t, err)
	require.NotEmpty(t, racket.ID)

	quantity := 10
	cart := &model.Cart{
		UserID:     user.ID,
		Quantity:   quantity,
		TotalPrice: float32(quantity) * float32(racket.Price),
		Lines: []*model.CartLine{{
			RacketID: racket.ID,
			Quantity: quantity,
		}},
	}

	err = repo.Create(ctx, cart)
	require.NoError(t, err)

	res, err := repo.GetCartByID(ctx, user.ID)
	require.NoError(t, err)
	require.Equal(t, res, cart)

	err = repo.Remove(ctx, user.ID)
	require.NoError(t, err)

	err = repoUser.Remove(ctx, user.Email)
	require.NoError(t, err)

	err = repoRacket.Remove(ctx, racket.ID)
	require.NoError(t, err)

	err = repoSupplier.Remove(ctx, supplier.Email)
	require.NoError(t, err)
}
