package mypostgres

import (
	"context"
	"src/internal/model"
	"src/internal/repository"
	"src/pkg/storage/postgres"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
)

type SupplierRepository struct {
	*postgres.Postgres
}

func NewSupplierRepository(db *postgres.Postgres) repository.ISupplierRepository {
	return &SupplierRepository{db}
}

func (r *SupplierRepository) Create(ctx context.Context, supplier *model.Supplier) error {

	query := r.Builder.
		Insert(supplierTable).
		Columns(nameField,
			phoneField,
			townField,
			emailField).
		Values(supplier.Name,
			supplier.Phone,
			supplier.Town,
			supplier.Email).
		Suffix("returning id")

	sql, ars, err := query.ToSql()

	if err != nil {
		return err
	}

	row := r.Pool.QueryRow(ctx, sql, ars...)

	err = row.Scan(
		&supplier.ID,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *SupplierRepository) Update(ctx context.Context, supplier *model.Supplier) error {

	query := r.Builder.
		Update(supplierTable).
		Set(nameField, supplier.Name).
		Set(townField, supplier.Town).
		Set(phoneField, supplier.Phone).
		Set(emailField, supplier.Email).
		Where(squirrel.Eq{emailField: supplier.Email})

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

func (r *SupplierRepository) Remove(ctx context.Context, email string) error {

	query := r.Builder.
		Delete(supplierTable).
		Where(squirrel.Eq{emailField: email})

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

func (r *SupplierRepository) GetSupplierByID(ctx context.Context, id int) (*model.Supplier, error) {

	query := r.Builder.
		Select("*").
		From(supplierTable).
		Where(squirrel.Eq{idField: id})

	sql, args, err := query.ToSql()

	if err != nil {
		return nil, err
	}

	row := r.Pool.QueryRow(ctx, sql, args...)

	return r.rowToModel(row)
}

func (r *SupplierRepository) GetSupplierByEmail(ctx context.Context, email string) (*model.Supplier, error) {

	query := r.Builder.
		Select("*").
		From(supplierTable).
		Where(squirrel.Eq{emailField: email})

	sql, args, err := query.ToSql()

	if err != nil {
		return nil, err
	}

	row := r.Pool.QueryRow(ctx, sql, args...)

	return r.rowToModel(row)
}

func (r *SupplierRepository) GetAllSuppliers(ctx context.Context) ([]*model.Supplier, error) {

	query := r.Builder.
		Select("*").
		From(supplierTable)

	sql, args, err := query.ToSql()

	if err != nil {
		return nil, err
	}

	rows, err := r.Pool.Query(ctx, sql, args...)

	if err != nil {
		return nil, err
	}

	var suppliers []*model.Supplier

	for rows.Next() {

		supplier, err := r.rowToModel(rows)

		if err != nil {
			return nil, err
		}

		suppliers = append(suppliers, supplier)
	}

	return suppliers, nil
}

func (r *SupplierRepository) rowToModel(row pgx.Row) (*model.Supplier, error) {

	var supplier model.Supplier

	err := row.Scan(
		&supplier.ID,
		&supplier.Email,
		&supplier.Name,
		&supplier.Phone,
		&supplier.Town,
	)

	if err != nil {
		return nil, err
	}

	return &supplier, nil
}
