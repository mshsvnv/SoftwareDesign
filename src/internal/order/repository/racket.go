package repository

import (
	"context"
	"src/internal/order/model"
)

//go:generate mockery --name=IOrderRacketRepository
type IOrderRacketRepository interface {
	GetRacketByID(ctx context.Context, id string) (*model.Racket, error)
}
