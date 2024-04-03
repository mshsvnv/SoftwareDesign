package repository

import (
	"context"

	"src/internal/cart/model"
)

//go:generate mockery --name=ICartRepository
type ICartRepository interface {
	Create(ctx context.Context, cart *model.Cart) error
	Update(ctx context.Context, cart *model.Cart) error
	GetCartByUserID(ctx context.Context, userID string) (*model.Cart, error)
}
