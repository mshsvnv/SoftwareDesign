package mypostgres

import (
	"context"
	"src_new/internal/dto"
	"src_new/internal/model"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestFeedbackRepositoryCreate(t *testing.T) {

	ctx := context.Background()

	repo := NewFeedbackRepository(testDB)
	repoUser := NewUserRepository(testDB)
	repoSupplier := NewSupplierRepository(testDB)
	repoRacket := NewRacketRepository(testDB)

	user := &model.User{
		Name:     "Ivan",
		Surname:  "Ivanov",
		Email:    "ivanov@mail.ru",
		Password: "123",
		Role:     "Customer",
	}

	supplier := &model.Supplier{
		Name:  "IP Ivanov",
		Email: "ivanov@mail.ru",
		Town:  "Armavir",
		Phone: "8-800-555-35-35",
	}

	err := repoUser.Create(ctx, user)
	require.NoError(t, err)
	require.NotEmpty(t, user.ID)

	err = repoSupplier.Create(ctx, supplier)
	require.NoError(t, err)
	require.NotEmpty(t, supplier.ID)

	racket := &model.Racket{
		SupplierID: supplier.ID,
		Brand:      "Babolat",
		Weight:     1000,
		Balance:    3.5,
		HeadSize:   20.2,
		Quantity:   100,
		Price:      100,
	}

	err = repoRacket.Create(ctx, racket)
	require.NoError(t, err)
	require.NotEmpty(t, racket.ID)

	feedback := &model.Feedback{
		UserID:   user.ID,
		RacketID: racket.ID,
		Feedback: "The best one!",
		Rating:   10,
		Date:     time.Now(),
	}

	err = repo.Create(ctx, feedback)
	require.NoError(t, err)

	err = repo.Remove(ctx, &dto.RemoveFeedbackReq{
		RacketID: racket.ID,
		UserID:   user.ID,
	})

	require.NoError(t, err)

	err = repoUser.Remove(ctx, user.Email)
	require.NoError(t, err)

	err = repoRacket.Remove(ctx, racket.ID)
	require.NoError(t, err)

	err = repoSupplier.Remove(ctx, supplier.Email)
	require.NoError(t, err)
}

func TestFeedbackRepositoryUpdate(t *testing.T) {

	ctx := context.Background()

	repo := NewFeedbackRepository(testDB)
	repoUser := NewUserRepository(testDB)
	repoSupplier := NewSupplierRepository(testDB)
	repoRacket := NewRacketRepository(testDB)

	user := &model.User{
		Name:     "Ivan",
		Surname:  "Ivanov",
		Email:    "ivanov@mail.ru",
		Password: "123",
		Role:     "Customer",
	}

	supplier := &model.Supplier{
		Name:  "IP Ivanov",
		Email: "ivanov@mail.ru",
		Town:  "Armavir",
		Phone: "8-800-555-35-35",
	}

	err := repoUser.Create(ctx, user)
	require.NoError(t, err)
	require.NotEmpty(t, user.ID)

	err = repoSupplier.Create(ctx, supplier)
	require.NoError(t, err)
	require.NotEmpty(t, supplier.ID)

	racket := &model.Racket{
		SupplierID: supplier.ID,
		Brand:      "Babolat",
		Weight:     1000,
		Balance:    3.5,
		HeadSize:   20.2,
		Quantity:   100,
		Price:      100,
	}

	err = repoRacket.Create(ctx, racket)
	require.NoError(t, err)
	require.NotEmpty(t, racket.ID)

	feedback := &model.Feedback{
		UserID:   user.ID,
		RacketID: racket.ID,
		Feedback: "The best one!",
		Rating:   10,
		Date:     time.Now(),
	}

	err = repo.Create(ctx, feedback)
	require.NoError(t, err)

	feedback.Rating = 5
	err = repo.Update(ctx, feedback)
	require.NoError(t, err)

	err = repo.Remove(ctx, &dto.RemoveFeedbackReq{
		RacketID: racket.ID,
		UserID:   user.ID,
	})

	require.NoError(t, err)

	err = repoUser.Remove(ctx, user.Email)
	require.NoError(t, err)

	err = repoRacket.Remove(ctx, racket.ID)
	require.NoError(t, err)

	err = repoSupplier.Remove(ctx, supplier.Email)
	require.NoError(t, err)
}

func TestFeedbackRepositoryGetFeedback(t *testing.T) {

	ctx := context.Background()

	repo := NewFeedbackRepository(testDB)
	repoUser := NewUserRepository(testDB)
	repoSupplier := NewSupplierRepository(testDB)
	repoRacket := NewRacketRepository(testDB)

	user := &model.User{
		Name:     "Ivan",
		Surname:  "Ivanov",
		Email:    "ivanov@mail.ru",
		Password: "123",
		Role:     "Customer",
	}

	supplier := &model.Supplier{
		Name:  "IP Ivanov",
		Email: "ivanov@mail.ru",
		Town:  "Armavir",
		Phone: "8-800-555-35-35",
	}

	err := repoUser.Create(ctx, user)
	require.NoError(t, err)
	require.NotEmpty(t, user.ID)

	err = repoSupplier.Create(ctx, supplier)
	require.NoError(t, err)
	require.NotEmpty(t, supplier.ID)

	racket := &model.Racket{
		SupplierID: supplier.ID,
		Brand:      "Babolat",
		Weight:     1000,
		Balance:    3.5,
		HeadSize:   20.2,
		Quantity:   100,
		Price:      100,
	}

	err = repoRacket.Create(ctx, racket)
	require.NoError(t, err)
	require.NotEmpty(t, racket.ID)

	dateStr := "01-06-2024"
	layout := "02-01-2006"
	date, _ := time.Parse(layout, dateStr)

	feedback := &model.Feedback{
		UserID:   user.ID,
		RacketID: racket.ID,
		Feedback: "The best one!",
		Rating:   10,
		Date:     date,
	}

	err = repo.Create(ctx, feedback)
	require.NoError(t, err)

	res, err := repo.GetFeedback(ctx, &dto.GetFeedbackReq{
		RacketID: racket.ID,
		UserID:   user.ID,
	})

	require.NoError(t, err)
	require.Equal(t, res, feedback)

	err = repo.Remove(ctx, &dto.RemoveFeedbackReq{
		RacketID: racket.ID,
		UserID:   user.ID,
	})

	require.NoError(t, err)

	err = repoUser.Remove(ctx, user.Email)
	require.NoError(t, err)

	err = repoRacket.Remove(ctx, racket.ID)
	require.NoError(t, err)

	err = repoSupplier.Remove(ctx, supplier.Email)
	require.NoError(t, err)
}

func TestFeedbackRepositoryGetFeedbackByUserID(t *testing.T) {

	ctx := context.Background()

	repo := NewFeedbackRepository(testDB)
	repoUser := NewUserRepository(testDB)
	repoSupplier := NewSupplierRepository(testDB)
	repoRacket := NewRacketRepository(testDB)

	user := &model.User{
		Name:     "Ivan",
		Surname:  "Ivanov",
		Email:    "ivanov@mail.ru",
		Password: "123",
		Role:     "Customer",
	}

	supplier := &model.Supplier{
		Name:  "IP Ivanov",
		Email: "ivanov@mail.ru",
		Town:  "Armavir",
		Phone: "8-800-555-35-35",
	}

	err := repoUser.Create(ctx, user)
	require.NoError(t, err)
	require.NotEmpty(t, user.ID)

	err = repoSupplier.Create(ctx, supplier)
	require.NoError(t, err)
	require.NotEmpty(t, supplier.ID)

	racket := &model.Racket{
		SupplierID: supplier.ID,
		Brand:      "Babolat",
		Weight:     1000,
		Balance:    3.5,
		HeadSize:   20.2,
		Quantity:   100,
		Price:      100,
	}

	err = repoRacket.Create(ctx, racket)
	require.NoError(t, err)
	require.NotEmpty(t, racket.ID)

	dateStr := "01-06-2024"
	layout := "02-01-2006"
	date, _ := time.Parse(layout, dateStr)

	feedback := &model.Feedback{
		UserID:   user.ID,
		RacketID: racket.ID,
		Feedback: "The best one!",
		Rating:   10,
		Date:     date,
	}

	err = repo.Create(ctx, feedback)
	require.NoError(t, err)

	res, err := repo.GetFeedbacksByUserID(ctx, user.ID)
	require.NoError(t, err)
	require.Equal(t, res[0], feedback)

	err = repo.Remove(ctx, &dto.RemoveFeedbackReq{
		RacketID: racket.ID,
		UserID:   user.ID,
	})

	require.NoError(t, err)

	err = repoUser.Remove(ctx, user.Email)
	require.NoError(t, err)

	err = repoRacket.Remove(ctx, racket.ID)
	require.NoError(t, err)

	err = repoSupplier.Remove(ctx, supplier.Email)
	require.NoError(t, err)
}

func TestFeedbackRepositoryGetFeedbackByRacketID(t *testing.T) {

	ctx := context.Background()

	repo := NewFeedbackRepository(testDB)
	repoUser := NewUserRepository(testDB)
	repoSupplier := NewSupplierRepository(testDB)
	repoRacket := NewRacketRepository(testDB)

	user := &model.User{
		Name:     "Ivan",
		Surname:  "Ivanov",
		Email:    "ivanov@mail.ru",
		Password: "123",
		Role:     "Customer",
	}

	supplier := &model.Supplier{
		Name:  "IP Ivanov",
		Email: "ivanov@mail.ru",
		Town:  "Armavir",
		Phone: "8-800-555-35-35",
	}

	err := repoUser.Create(ctx, user)
	require.NoError(t, err)
	require.NotEmpty(t, user.ID)

	err = repoSupplier.Create(ctx, supplier)
	require.NoError(t, err)
	require.NotEmpty(t, supplier.ID)

	racket := &model.Racket{
		SupplierID: supplier.ID,
		Brand:      "Babolat",
		Weight:     1000,
		Balance:    3.5,
		HeadSize:   20.2,
		Quantity:   100,
		Price:      100,
	}

	err = repoRacket.Create(ctx, racket)
	require.NoError(t, err)
	require.NotEmpty(t, racket.ID)

	dateStr := "01-06-2024"
	layout := "02-01-2006"
	date, _ := time.Parse(layout, dateStr)

	feedback := &model.Feedback{
		UserID:   user.ID,
		RacketID: racket.ID,
		Feedback: "The best one!",
		Rating:   10,
		Date:     date,
	}

	err = repo.Create(ctx, feedback)
	require.NoError(t, err)

	res, err := repo.GetFeedbacksByRacketID(ctx, racket.ID)
	require.NoError(t, err)
	require.Equal(t, res[0], feedback)

	err = repo.Remove(ctx, &dto.RemoveFeedbackReq{
		RacketID: racket.ID,
		UserID:   user.ID,
	})

	require.NoError(t, err)

	err = repoUser.Remove(ctx, user.Email)
	require.NoError(t, err)

	err = repoRacket.Remove(ctx, racket.ID)
	require.NoError(t, err)

	err = repoSupplier.Remove(ctx, supplier.Email)
	require.NoError(t, err)
}
