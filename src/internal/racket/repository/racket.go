package repository

import (
	"context"

	"src/internal/racket/dto"
	"src/internal/racket/model"
)

//go:generate mockery --name=IRacketRepository
type IRacketRepository interface {
	Create(ctx context.Context, Racket *model.Racket) error
	Update(ctx context.Context, Racket *model.Racket) error
	ListRackets(ctx context.Context, req *dto.ListRacketReq) ([]*model.Racket, error)
	GetRacketByID(ctx context.Context, id string) (*model.Racket, error)
}
