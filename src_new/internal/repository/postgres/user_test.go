package mypostgres

import (
	"context"
	"src_new/internal/model"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUserRepositoryCreate(t *testing.T) {

	ctx := context.Background()

	repo := NewUserRepository(testDB)

	user := &model.User{
		Name:     "Ivan",
		Surname:  "Ivanov",
		Email:    "ivanov@mail.ru",
		Password: "123",
		Role:     "Customer",
	}

	err := repo.Create(ctx, user)
	require.NoError(t, err)
	require.NotEmpty(t, user.ID)

	err = repo.Remove(ctx, user.Email)
	require.NoError(t, err)
}

func TestUserRepositoryGetUserByID(t *testing.T) {

	ctx := context.Background()

	repo := NewUserRepository(testDB)

	user := &model.User{
		Name:     "Ivan",
		Surname:  "Ivanov",
		Email:    "ivanov@mail.ru",
		Password: "123",
		Role:     "Customer",
	}

	err := repo.Create(ctx, user)
	require.NoError(t, err)
	require.NotEmpty(t, user.ID)

	user, err = repo.GetUserByID(ctx, user.ID)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	err = repo.Remove(ctx, user.Email)
	require.NoError(t, err)
}

func TestUserRepositoryGetUserByEmail(t *testing.T) {

	ctx := context.Background()

	repo := NewUserRepository(testDB)

	user := &model.User{
		Name:     "Ivan",
		Surname:  "Ivanov",
		Email:    "ivanov@mail.ru",
		Password: "123",
		Role:     "Customer",
	}

	err := repo.Create(ctx, user)
	require.NoError(t, err)
	require.NotEmpty(t, user.ID)

	user, err = repo.GetUserByEmail(ctx, user.Email)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	err = repo.Remove(ctx, user.Email)
	require.NoError(t, err)
}

func TestUserRepositoryRemove(t *testing.T) {

	ctx := context.Background()

	repo := NewUserRepository(testDB)

	user := &model.User{
		Name:     "Ivan",
		Surname:  "Ivanov",
		Email:    "ivanov@mail.ru",
		Password: "123",
		Role:     "Customer",
	}

	err := repo.Create(ctx, user)
	require.NoError(t, err)

	err = repo.Remove(ctx, user.Email)
	require.NoError(t, err)
}
