package repository

import (
	"context"
	"src_new/internal/model"
)

//go:generate mockery --name=ICartRepository
type ICartRepository interface {
	Create(ctx context.Context, cart *model.Cart) error
	Update(ctx context.Context, cart *model.Cart) error
	Remove(ctx context.Context, userID int) error
	GetCartByID(ctx context.Context, userID int) (*model.Cart, error)
}
