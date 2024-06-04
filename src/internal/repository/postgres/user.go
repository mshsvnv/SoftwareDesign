package mypostgres

import (
	"context"
	"src/internal/model"
	"src/internal/repository"
	"src/pkg/storage/postgres"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
)

type UserRepository struct {
	*postgres.Postgres
}

func NewUserRepository(db *postgres.Postgres) repository.IUserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) Create(ctx context.Context, user *model.User) error {

	query := r.Builder.
		Insert(userTable).
		Columns(
			nameField,
			surnameField,
			emailField,
			passwordField,
			roleField).
		Values(user.Name,
			user.Surname,
			user.Email,
			user.Password,
			user.Role).
		Suffix("returning id")

	sql, ars, err := query.ToSql()

	if err != nil {
		return err
	}

	row := r.Pool.QueryRow(ctx, sql, ars...)

	err = row.Scan(
		&user.ID,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) UpdateRole(ctx context.Context, user *model.User) error {

	query := r.Builder.
		Update(userTable).
		Set(roleField, user.Role).
		Where(squirrel.Eq{emailField: user.Email})

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

func (r *UserRepository) GetAllUsers(ctx context.Context) ([]*model.User, error) {

	query := r.Builder.
		Select("*").
		From(userTable)

	sql, args, err := query.ToSql()

	if err != nil {
		return nil, err
	}

	rows, err := r.Pool.Query(ctx, sql, args...)

	if err != nil {
		return nil, err
	}

	var users []*model.User

	for rows.Next() {

		user, err := r.rowToModel(rows)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (r *UserRepository) GetUserByID(ctx context.Context, id int) (*model.User, error) {

	query := r.Builder.
		Select("*").
		From(userTable).
		Where(squirrel.Eq{idField: id})

	sql, args, err := query.ToSql()

	if err != nil {
		return nil, err
	}

	row := r.Pool.QueryRow(ctx, sql, args...)

	return r.rowToModel(row)
}

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {

	query := r.Builder.
		Select("*").
		From(userTable).
		Where(squirrel.Eq{emailField: email})

	sql, args, err := query.ToSql()

	if err != nil {
		return nil, err
	}

	row := r.Pool.QueryRow(ctx, sql, args...)

	return r.rowToModel(row)
}

func (r *UserRepository) Remove(ctx context.Context, email string) error {

	query := r.Builder.
		Delete(userTable).
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

func (r *UserRepository) rowToModel(row pgx.Row) (*model.User, error) {

	var user model.User

	err := row.Scan(
		&user.ID,
		&user.Name,
		&user.Surname,
		&user.Email,
		&user.Password,
		&user.Role,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
