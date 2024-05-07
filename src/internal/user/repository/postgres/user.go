package postgres

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"

	"src/internal/user/model"
	"src/internal/user/repository"
	"src/pkg/dbs/postgres"
)

type userRepository struct {
	*postgres.Postgres
}

func NewUserRepository(db *postgres.Postgres) repository.IUserRepository {
	return &userRepository{db}
}

func (p *userRepository) Create(ctx context.Context, user *model.User) error {

	query := p.Builder.
		Insert(userTable).
		Columns(
			emailField,
			passwordField,
			roleField,
		).
		Values(
			user.Email,
			user.Password,
			user.Role)

	sql, args, err := query.ToSql()

	if err != nil {
		return err
	}

	p.Pool.QueryRow(ctx, sql, args...)

	return nil
}

func (p *userRepository) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {

	query := p.Builder.
		Select(
			idField,
			emailField,
			passwordField,
			roleField,
		).
		From(userTable).
		Where(squirrel.Eq{emailField: email})

	sql, args, err := query.ToSql()

	if err != nil {
		return nil, err
	}

	row := p.Pool.QueryRow(ctx, sql, args...)

	return p.rowToModel(row)
}

func (p *userRepository) GetUserByID(ctx context.Context, id string) (*model.User, error) {

	query := p.Builder.
		Select(
			idField,
			emailField,
			passwordField,
			roleField,
		).
		From(userTable).
		Where(squirrel.Eq{idField: id})

	sql, args, err := query.ToSql()

	if err != nil {
		return nil, err
	}

	row := p.Pool.QueryRow(ctx, sql, args...)

	return p.rowToModel(row)
}

func (p *userRepository) rowToModel(row pgx.Row) (*model.User, error) {

	var user model.User

	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.Role,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
