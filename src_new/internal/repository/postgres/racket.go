package mypostgres

import (
	"context"
	"src_new/internal/model"
	"src_new/internal/repository"
	"src_new/pkg/storage/postgres"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
)

type RacketRepository struct {
	*postgres.Postgres
}

func NewRacketRepository(db *postgres.Postgres) repository.IRacketRepository {
	return &RacketRepository{db}
}

func (r *RacketRepository) Create(ctx context.Context, racket *model.Racket) error {

	query := r.Builder.
		Insert(racketTable).
		Columns(
			supplierIDField,
			brandField,
			weightField,
			balanceField,
			headSizeField,
			quantityField,
			priceField,
			avaliableField).
		Values(
			racket.SupplierID,
			racket.Brand,
			racket.Weight,
			racket.Balance,
			racket.HeadSize,
			racket.Quantity,
			racket.Price,
			racket.Avaliable).
		Suffix("returning id")

	sql, ars, err := query.ToSql()

	if err != nil {
		return err
	}

	row := r.Pool.QueryRow(ctx, sql, ars...)

	err = row.Scan(
		&racket.ID,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *RacketRepository) Update(ctx context.Context, racket *model.Racket) error {

	query := r.Builder.
		Update(racketTable).
		Set(supplierIDField, racket.SupplierID).
		Set(brandField, racket.Brand).
		Set(weightField, racket.Weight).
		Set(balanceField, racket.Balance).
		Set(headSizeField, racket.HeadSize).
		Set(quantityField, racket.Quantity).
		Set(priceField, racket.Price).
		Set(avaliableField, racket.Avaliable).
		Where(squirrel.Eq{idField: racket.ID})

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

func (r *RacketRepository) Remove(ctx context.Context, id int) error {

	query := r.Builder.
		Delete(racketTable).
		Where(squirrel.Eq{idField: id})

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

func (r *RacketRepository) GetRacketByID(ctx context.Context, id int) (*model.Racket, error) {

	query := r.Builder.
		Select("*").
		From(racketTable).
		Where(squirrel.Eq{idField: id})

	sql, args, err := query.ToSql()

	if err != nil {
		return nil, err
	}

	row := r.Pool.QueryRow(ctx, sql, args...)

	return r.rowToModel(row)
}

func (r *RacketRepository) GetAllRackets(ctx context.Context) ([]*model.Racket, error) {

	query := r.Builder.
		Select("*").
		From(racketTable)

	sql, args, err := query.ToSql()

	if err != nil {
		return nil, err
	}

	rows, err := r.Pool.Query(ctx, sql, args...)

	if err != nil {
		return nil, err
	}

	var rackets []*model.Racket

	for rows.Next() {

		racket, err := r.rowToModel(rows)

		if err != nil {
			return nil, err
		}

		rackets = append(rackets, racket)
	}

	return rackets, nil
}

func (r *RacketRepository) rowToModel(row pgx.Row) (*model.Racket, error) {

	var racket model.Racket

	err := row.Scan(
		&racket.ID,
		&racket.SupplierID,
		&racket.Brand,
		&racket.Weight,
		&racket.Balance,
		&racket.HeadSize,
		&racket.Avaliable,
		&racket.Quantity,
		&racket.Price,
	)

	if err != nil {
		return nil, err
	}

	return &racket, nil
}
