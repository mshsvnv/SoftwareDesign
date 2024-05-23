package mypostgres

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"src_new/internal/model"
)

func TestOrderRepositoryCreate(t *testing.T) {

	ctx := context.Background()

	repo := NewOrderRepository(testDB)

	dateStr := "01-06-2024"
	layout := "02-01-2006"
	date, _ := time.Parse(layout, dateStr)

	order := &model.Order{
		UserID: ids["userID"],
		Status: model.OrderStatusInProgress,
		OrderInfo: &model.OrderInfo{
			DeliveryDate:  date,
			Address:       "Moscow, Red Square",
			RecepientName: "Petrov Peter",
		},
		Lines: []*model.OrderLine{
			{
				RacketID: ids["racketID"],
			},
		},
		CreationDate: date,
	}

	err := repo.Create(ctx, order)

	require.NoError(t, err)
	require.NotEmpty(t, order.ID)

	err = repo.Remove(ctx, order.ID)
	require.NoError(t, err)
}

func TestOrderRepositoryGetOrderByID(t *testing.T) {

	ctx := context.Background()

	repo := NewOrderRepository(testDB)

	dateStr := "01-06-2024"
	layout := "02-01-2006"
	date, _ := time.Parse(layout, dateStr)

	quantity := 10
	order := &model.Order{
		UserID: ids["userID"],
		Status: model.OrderStatusInProgress,
		OrderInfo: &model.OrderInfo{
			DeliveryDate:  date,
			Address:       "Moscow, Red Square",
			RecepientName: "Petrov Peter",
		},
		Lines: []*model.OrderLine{
			{
				RacketID: ids["racketID"],
				Quantity: quantity,
			},
		},
		CreationDate: date,
	}

	err := repo.Create(ctx, order)

	require.NoError(t, err)
	require.NotEmpty(t, order.ID)

	res, err := repo.GetOrderByID(ctx, order.ID)
	require.NoError(t, err)
	require.Equal(t, res, order)

	err = repo.Remove(ctx, order.ID)
	require.NoError(t, err)
}

func TestOrderRepositoryGetMyOrders(t *testing.T) {

	ctx := context.Background()

	repo := NewOrderRepository(testDB)

	dateStr := "01-06-2024"
	layout := "02-01-2006"
	date, _ := time.Parse(layout, dateStr)

	quantity := 10
	order := &model.Order{
		UserID: ids["userID"],
		Status: model.OrderStatusInProgress,
		OrderInfo: &model.OrderInfo{
			DeliveryDate:  date,
			Address:       "Moscow, Red Square",
			RecepientName: "Petrov Peter",
		},
		Lines: []*model.OrderLine{
			{
				RacketID: ids["racketID"],
				Quantity: quantity,
			},
		},
		CreationDate: date,
	}

	err := repo.Create(ctx, order)

	require.NoError(t, err)
	require.NotEmpty(t, order.ID)

	res, err := repo.GetMyOrders(ctx, ids["userID"])
	require.NoError(t, err)
	require.Equal(t, 1, len(res))

	err = repo.Remove(ctx, order.ID)
	require.NoError(t, err)
}
