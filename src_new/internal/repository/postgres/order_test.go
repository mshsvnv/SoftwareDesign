package mypostgres

import (
	"context"
	"src_new/internal/model"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestOrderRepositoryCreate(t *testing.T) {

	ctx := context.Background()

	repo := NewOrderRepository(testDB)
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

	dateStr := "01-06-2024"
	layout := "02-01-2006"
	date, _ := time.Parse(layout, dateStr)

	quantity := 10
	order := &model.Order{
		UserID: user.ID,
		Status: model.OrderStatusInProgress,
		OrderInfo: &model.OrderInfo{
			DeliveryDate:  date,
			Address:       "Moscow, Red Square",
			RecepientName: "Petrov Peter",
		},
		Lines: []*model.OrderLine{
			{
				RacketID: racket.ID,
				Quantity: quantity,
			},
		},
		TotalPrice: float32(racket.Price) * float32(quantity),
	}

	err = repo.Create(ctx, order)

	require.NoError(t, err)
	require.NotEmpty(t, order.ID)

	err = repo.Remove(ctx, order.ID)
	require.NoError(t, err)

	err = repoUser.Remove(ctx, user.Email)
	require.NoError(t, err)

	err = repoRacket.Remove(ctx, racket.ID)
	require.NoError(t, err)

	err = repoSupplier.Remove(ctx, supplier.Email)
	require.NoError(t, err)
}

func TestOrderRepositoryGetOrderByID(t *testing.T) {

	ctx := context.Background()

	repo := NewOrderRepository(testDB)
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

	dateStr := "01-06-2024"
	layout := "02-01-2006"
	date, _ := time.Parse(layout, dateStr)

	quantity := 10
	order := &model.Order{
		UserID: user.ID,
		Status: model.OrderStatusInProgress,
		OrderInfo: &model.OrderInfo{
			DeliveryDate:  date,
			Address:       "Moscow, Red Square",
			RecepientName: "Petrov Peter",
		},
		Lines: []*model.OrderLine{
			{
				RacketID: racket.ID,
				Quantity: quantity,
			},
		},
		TotalPrice: float32(racket.Price) * float32(quantity),
	}

	err = repo.Create(ctx, order)

	require.NoError(t, err)
	require.NotEmpty(t, order.ID)

	res, err := repo.GetOrderByID(ctx, order.ID)
	res.Lines = order.Lines
	require.NoError(t, err)
	require.Equal(t, res, order)

	err = repo.Remove(ctx, order.ID)
	require.NoError(t, err)

	err = repoUser.Remove(ctx, user.Email)
	require.NoError(t, err)

	err = repoRacket.Remove(ctx, racket.ID)
	require.NoError(t, err)

	err = repoSupplier.Remove(ctx, supplier.Email)
	require.NoError(t, err)
}

func TestOrderRepositoryGetMyOrders(t *testing.T) {

	ctx := context.Background()

	repo := NewOrderRepository(testDB)
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

	dateStr := "01-06-2024"
	layout := "02-01-2006"
	date, _ := time.Parse(layout, dateStr)

	quantity := 10
	order := &model.Order{
		UserID: user.ID,
		Status: model.OrderStatusInProgress,
		OrderInfo: &model.OrderInfo{
			DeliveryDate:  date,
			Address:       "Moscow, Red Square",
			RecepientName: "Petrov Peter",
		},
		Lines: []*model.OrderLine{
			{
				RacketID: racket.ID,
				Quantity: quantity,
			},
		},
		TotalPrice: float32(racket.Price) * float32(quantity),
	}

	err = repo.Create(ctx, order)

	require.NoError(t, err)
	require.NotEmpty(t, order.ID)

	res, err := repo.GetMyOrders(ctx, user.ID)
	res[0].Lines = order.Lines
	require.NoError(t, err)
	require.Equal(t, res[0], order)

	err = repo.Remove(ctx, order.ID)
	require.NoError(t, err)

	err = repoUser.Remove(ctx, user.Email)
	require.NoError(t, err)

	err = repoRacket.Remove(ctx, racket.ID)
	require.NoError(t, err)

	err = repoSupplier.Remove(ctx, supplier.Email)
	require.NoError(t, err)
}
