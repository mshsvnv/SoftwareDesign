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

	quantity := 10
	cart := &model.Cart{
		UserID:   ids["userID"],
		Quantity: quantity,
		Lines: []*model.CartLine{{
			RacketID: ids["racketID"],
			Quantity: quantity,
		}},
	}

	err := repo.Create(ctx, cart)
	require.NoError(t, err)

	err = repo.Remove(ctx, ids["userID"])
	require.NoError(t, err)
}

func TestCartRepositoryUpdate(t *testing.T) {

	ctx := context.Background()

	repo := NewCartRepository(testDB)

	quantity := 10
	cart := &model.Cart{
		UserID:   ids["userID"],
		Quantity: quantity,
	
		Lines: []*model.CartLine{{
			RacketID: ids["racketID"],
			Quantity: quantity,
		}},
	}

	err := repo.Create(ctx, cart)
	require.NoError(t, err)

	cart.Quantity -= 1

	err = repo.Update(ctx, cart)
	require.NoError(t, err)

	err = repo.Remove(ctx, ids["userID"])
	require.NoError(t, err)
}

func TestCartRepositoryGetCartByID(t *testing.T) {

	ctx := context.Background()

	repo := NewCartRepository(testDB)

	quantity := 10
	cart := &model.Cart{
		UserID:   ids["userID"],
		Quantity: quantity,
		Lines: []*model.CartLine{{
			RacketID: ids["racketID"],
			Quantity: quantity,
		}},
	}

	err := repo.Create(ctx, cart)
	require.NoError(t, err)

	res, err := repo.GetCartByID(ctx, ids["userID"])
	require.NoError(t, err)
	require.Equal(t, res, cart)
}
