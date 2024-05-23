package mypostgres

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"src_new/internal/dto"
	"src_new/internal/model"
)

func TestFeedbackRepositoryCreate(t *testing.T) {

	ctx := context.Background()

	repo := NewFeedbackRepository(testDB)

	feedback := &model.Feedback{
		UserID:   ids["userID"],
		RacketID: ids["racketID"],
		Feedback: "The best one!",
		Rating:   10,
		Date:     time.Now(),
	}

	err := repo.Create(ctx, feedback)
	require.NoError(t, err)

	err = repo.Remove(ctx, &dto.RemoveFeedbackReq{
		RacketID: ids["racketID"],
		UserID:   ids["userID"],
	})

	require.NoError(t, err)
}

func TestFeedbackRepositoryUpdate(t *testing.T) {

	ctx := context.Background()

	repo := NewFeedbackRepository(testDB)

	feedback := &model.Feedback{
		UserID:   ids["userID"],
		RacketID: ids["racketID"],
		Feedback: "The best one!",
		Rating:   10,
		Date:     time.Now(),
	}

	err := repo.Create(ctx, feedback)
	require.NoError(t, err)

	feedback.Rating = 5
	err = repo.Update(ctx, feedback)
	require.NoError(t, err)

	err = repo.Remove(ctx, &dto.RemoveFeedbackReq{
		RacketID: ids["racketID"],
		UserID:   ids["userID"],
	})

	require.NoError(t, err)
}

func TestFeedbackRepositoryGetFeedback(t *testing.T) {

	ctx := context.Background()

	repo := NewFeedbackRepository(testDB)

	dateStr := "01-06-2024"
	layout := "02-01-2006"
	date, _ := time.Parse(layout, dateStr)

	feedback := &model.Feedback{
		UserID:   ids["userID"],
		RacketID: ids["racketID"],
		Feedback: "The best one!",
		Rating:   10,
		Date:     date,
	}

	err := repo.Create(ctx, feedback)
	require.NoError(t, err)

	res, err := repo.GetFeedback(ctx, &dto.GetFeedbackReq{
		RacketID: ids["racketID"],
		UserID:   ids["userID"],
	})

	require.NoError(t, err)
	require.Equal(t, res, feedback)

	err = repo.Remove(ctx, &dto.RemoveFeedbackReq{
		RacketID: ids["racketID"],
		UserID:   ids["userID"],
	})

	require.NoError(t, err)
}

func TestFeedbackRepositoryGetFeedbackByUserID(t *testing.T) {

	ctx := context.Background()

	repo := NewFeedbackRepository(testDB)

	dateStr := "01-06-2024"
	layout := "02-01-2006"
	date, _ := time.Parse(layout, dateStr)

	feedback := &model.Feedback{
		UserID:   ids["userID"],
		RacketID: ids["racketID"],
		Feedback: "The best one!",
		Rating:   10,
		Date:     date,
	}

	err := repo.Create(ctx, feedback)
	require.NoError(t, err)

	res, err := repo.GetFeedbacksByUserID(ctx, ids["userID"])
	require.NoError(t, err)
	require.Equal(t, res[0], feedback)

	err = repo.Remove(ctx, &dto.RemoveFeedbackReq{
		RacketID: ids["racketID"],
		UserID:   ids["userID"],
	})

	require.NoError(t, err)
}

func TestFeedbackRepositoryGetFeedbackByRacketID(t *testing.T) {

	ctx := context.Background()

	repo := NewFeedbackRepository(testDB)

	dateStr := "01-06-2024"
	layout := "02-01-2006"
	date, _ := time.Parse(layout, dateStr)

	feedback := &model.Feedback{
		UserID:   ids["userID"],
		RacketID: ids["racketID"],
		Feedback: "The best one!",
		Rating:   10,
		Date:     date,
	}

	err := repo.Create(ctx, feedback)
	require.NoError(t, err)

	res, err := repo.GetFeedbacksByRacketID(ctx, ids["racketID"])
	require.NoError(t, err)
	require.Equal(t, 1, len(res))

	err = repo.Remove(ctx, &dto.RemoveFeedbackReq{
		RacketID: ids["racketID"],
		UserID:   ids["userID"],
	})

	require.NoError(t, err)
}
