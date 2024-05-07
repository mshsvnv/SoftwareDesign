package repository

import (
	"context"
	"src_new/internal/model"
)

//go:generate mockery --name=IRacketRepository
type IRacketRepository interface {
	Create(ctx context.Context, racket *model.Racket) error
	Update(ctx context.Context, racket *model.Racket) error
	Remove(ctx context.Context, id int) error
	GetRacketByID(ctx context.Context, id int) (*model.Racket, error)
	GetAllRackets(ctx context.Context) ([]*model.Racket, error)
}

// type RacketRepository struct {
// 	db []*model.Racket
// }

// func NewRacketRepository() *RacketRepository {
// 	return &RacketRepository{}
// }

// func (r *RacketRepository) Create(ctx context.Context, racket *model.Racket) error {

// 	r.db = append(r.db, racket)
// 	r.db[len(r.db)-1].ID = len(r.db) - 1

// 	return nil
// }

// func (r *RacketRepository) Update(ctx context.Context, racket *model.Racket) error {

// 	for i := 0; i < len(r.db); i++ {

// 		if r.db[i].ID == racket.ID {
// 			r.db[i] = racket

// 			return nil
// 		}
// 	}

// 	return fmt.Errorf("no racket")
// }

// func (r *RacketRepository) Remove(ctx context.Context, id int) error {

// 	for i := 0; i < len(r.db); i++ {

// 		if r.db[i].ID == id {

// 			r.db = append(r.db[:i], r.db[i+1:]...)
// 			return nil
// 		}
// 	}

// 	return fmt.Errorf("no racket")
// }

// func (r *RacketRepository) GetRacketByID(ctx context.Context, id int) (*model.Racket, error) {

// 	for i := 0; i < len(r.db); i++ {

// 		if r.db[i].ID == id {
// 			return r.db[i], nil
// 		}
// 	}

// 	return nil, fmt.Errorf("no racket")
// }

// func (r *RacketRepository) GetAllRackets(ctx context.Context) ([]*model.Racket, error) {

// 	if len(r.db) == 0 {
// 		return nil, fmt.Errorf("no rackets")
// 	}

// 	return r.db, nil
// }
