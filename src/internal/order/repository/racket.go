package repository

import (
	"context"
	"src/internal/order/model"
)

//go:generate mockery --name=IRacketRepository
type IRacketRepository interface {
	GetRacketByID(ctx context.Context, id string) (*model.Racket, error)
}
