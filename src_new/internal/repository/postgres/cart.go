package mypostgres

import (
	"context"
	"src_new/internal/model"
	"src_new/internal/repository"
	"src_new/pkg/storage/postgres"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
)

type CartRepository struct {
	*postgres.Postgres
}

func NewCartRepository(db *postgres.Postgres) repository.ICartRepository {
	return &CartRepository{db}
}

func (r *CartRepository) Create(ctx context.Context, cart *model.Cart) error {

	query := r.Builder.
		Insert(cartTable).
		Columns(
			userIDField,
			totalPriceField,
			quantityField).
		Values(
			cart.UserID,
			cart.TotalPrice,
			cart.Quantity).
		Suffix("returning user_id")

	sql, args, err := query.ToSql()

	if err != nil {
		return err
	}

	row := r.Pool.QueryRow(ctx, sql, args...)

	err = row.Scan(
		&cart.UserID,
	)

	if err != nil {
		return err
	}

	for _, line := range cart.Lines {

		query := r.Builder.
			Insert(cartRacketTable).
			Columns(racketIDField,
				cartIDField,
				quantityField).
			Values(
				line.RacketID,
				cart.UserID,
				line.Quantity).
			Suffix("returning racket_id")

		sql, args, err := query.ToSql()

		if err != nil {
			return err
		}

		row := r.Pool.QueryRow(ctx, sql, args...)

		err = row.Scan(
			&line.RacketID,
		)

		if err != nil {
			return err
		}
	}

	return nil
}

func (r *CartRepository) Update(ctx context.Context, cart *model.Cart) error {

	for _, line := range cart.Lines {

		query := r.Builder.
			Update(cartRacketTable).
			Set(quantityField, line.Quantity).
			Where(squirrel.And{
				squirrel.Eq{cartIDField: cart.UserID},
				squirrel.Eq{racketIDField: line.RacketID},
			})

		sql, args, err := query.ToSql()

		if err != nil {
			return err
		}

		_, err = r.Pool.Exec(ctx, sql, args...)

		if err != nil {
			return err
		}
	}

	query := r.Builder.
		Update(cartTable).
		Set(quantityField, cart.Quantity).
		Set(totalPriceField, cart.TotalPrice).
		Where(squirrel.Eq{userIDField: cart.UserID})

	sql, args, err := query.ToSql()

	if err != nil {
		return err
	}

	_, err = r.Pool.Exec(ctx, sql, args...)

	if err != nil {
		return err
	}

	return nil
}

func (r *CartRepository) GetCartByID(ctx context.Context, userID int) (*model.Cart, error) {

	query := r.Builder.
		Select("*").
		From(cartTable).
		Where(squirrel.Eq{userIDField: userID})

	sql, args, err := query.ToSql()

	if err != nil {
		return nil, err
	}

	row := r.Pool.QueryRow(ctx, sql, args...)

	cart, err := r.rowToModel(row)

	if err != nil {
		return nil, err
	}

	query = r.Builder.
		Select(racketIDField, quantityField).
		From(cartRacketTable).
		Where(squirrel.Eq{cartIDField: userID})

	sql, args, err = query.ToSql()

	if err != nil {
		return nil, err
	}

	rows, err := r.Pool.Query(ctx, sql, args...)

	if err != nil {
		return nil, err
	}

	for rows.Next() {

		line, err := r.rowToModelCartRacket(rows)

		if err != nil {
			return nil, err
		}

		cart.Lines = append(cart.Lines, line)
	}

	return cart, nil
}

func (r *CartRepository) Remove(ctx context.Context, userID int) error {

	query := r.Builder.
		Delete(cartTable).
		Where(squirrel.Eq{userIDField: userID})

	sql, args, err := query.ToSql()

	if err != nil {
		return err
	}

	_, err = r.Pool.Exec(ctx, sql, args...)

	if err != nil {
		return err
	}

	query = r.Builder.
		Delete(cartRacketTable).
		Where(squirrel.Eq{cartIDField: userID})

	sql, args, err = query.ToSql()

	if err != nil {
		return err
	}

	_, err = r.Pool.Exec(ctx, sql, args...)

	if err != nil {
		return err
	}

	return nil
}

func (r *CartRepository) rowToModel(row pgx.Row) (*model.Cart, error) {

	var cart model.Cart

	err := row.Scan(
		&cart.UserID,
		&cart.Quantity,
		&cart.TotalPrice,
	)

	if err != nil {
		return nil, err
	}

	return &cart, nil
}

func (r *CartRepository) rowToModelCartRacket(row pgx.Row) (*model.CartLine, error) {

	var cartLine model.CartLine

	err := row.Scan(
		&cartLine.RacketID,
		&cartLine.Quantity,
	)

	if err != nil {
		return nil, err
	}

	return &cartLine, nil
}
